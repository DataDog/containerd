//go:build linux

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/containerd/containerd/v2/core/containers"
	"github.com/containerd/containerd/v2/core/content"
	"github.com/containerd/containerd/v2/core/images"
	"github.com/containerd/containerd/v2/core/mount"
	"github.com/containerd/containerd/v2/core/snapshots"
	"github.com/containerd/containerd/v2/internal/userns"
	"github.com/containerd/containerd/v2/pkg/testutil"
	"github.com/containerd/containerd/v2/plugins/content/local"
	"github.com/containerd/containerd/v2/plugins/snapshots/native"
	"github.com/containerd/errdefs"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
)

func TestRemappedSnapshotFailureCleanup(t *testing.T) {
	testutil.RequiresRoot(t)
	ctx := context.Background()
	store, err := local.NewStore(filepath.Join(t.TempDir(), "content"))
	require.NoError(t, err)
	sn, err := native.NewSnapshotter(filepath.Join(t.TempDir(), "snapshots"))
	require.NoError(t, err)
	defer sn.Close()
	c, err := New("", WithServices(WithContentStore(store), WithSnapshotters(map[string]snapshots.Snapshotter{"native": sn})))
	require.NoError(t, err)
	defer c.Close()
	writeBlob := func(value any, mediaType string) ocispec.Descriptor {
		data, err := json.Marshal(value)
		require.NoError(t, err)
		desc := ocispec.Descriptor{MediaType: mediaType, Digest: digest.FromBytes(data), Size: int64(len(data))}
		require.NoError(t, content.WriteBlob(ctx, store, desc.Digest.String(), bytes.NewReader(data), desc))
		return desc
	}
	diffID := digest.FromString("cleanup-test-layer")
	config := writeBlob(ocispec.Image{RootFS: ocispec.RootFS{Type: "layers", DiffIDs: []digest.Digest{diffID}}}, ocispec.MediaTypeImageConfig)
	manifest := ocispec.Manifest{Config: config}
	manifest.SchemaVersion = 2
	img := NewImage(c, images.Image{Name: "cleanup-test", Target: writeBlob(manifest, ocispec.MediaTypeImageManifest)})

	mounts, err := sn.Prepare(ctx, "base", "")
	require.NoError(t, err)
	require.NoError(t, mount.WithTempMount(ctx, mounts, func(root string) error {
		path := filepath.Join(root, "unmapped-owner")
		if err := os.WriteFile(path, nil, 0644); err != nil {
			return err
		}
		return os.Lchown(path, 65536, 0)
	}))
	require.NoError(t, sn.Commit(ctx, diffID.String(), "base"))
	idMap := userns.IDMap{
		UidMap: []specs.LinuxIDMapping{{ContainerID: 0, HostID: 100000, Size: 65536}},
		GidMap: []specs.LinuxIDMapping{{ContainerID: 0, HostID: 200000, Size: 65536}},
	}
	rsn := remappedSnapshot{Parent: diffID.String(), IDMap: idMap}
	remappedID, err := rsn.ID()
	require.NoError(t, err)
	opt := WithUserNSRemappedSnapshot("container", img, idMap.UidMap, idMap.GidMap)
	for range 2 {
		err := opt(ctx, c, &containers.Container{Snapshotter: "native"})
		require.ErrorContains(t, err, "container ID 65536 cannot be mapped")
		_, err = sn.Stat(ctx, remappedID+"-remap")
		require.True(t, errdefs.IsNotFound(err), "temporary snapshot should be removed: %v", err)
		_, err = sn.Stat(ctx, diffID.String())
		require.NoError(t, err)
	}
}

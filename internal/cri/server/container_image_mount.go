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

package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/core/leases"
	"github.com/containerd/containerd/v2/core/mount"
	"github.com/containerd/errdefs"
	"github.com/containerd/log"
	"github.com/containerd/platforms"
	"github.com/opencontainers/image-spec/identity"
	imagespec "github.com/opencontainers/image-spec/specs-go/v1"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (c *criService) mutateMounts(
	ctx context.Context,
	extraMounts []*runtime.Mount,
	snapshotter string,
	sandboxID string,
	platform imagespec.Platform,
) error {
	if err := c.ensureLeaseExist(ctx, sandboxID); err != nil {
		return fmt.Errorf("failed to ensure lease %v for sandbox: %w", sandboxID, err)
	}

	ctx = leases.WithLease(ctx, sandboxID)
	for _, m := range extraMounts {
		err := c.mutateImageMount(ctx, m, snapshotter, sandboxID, platform)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *criService) ensureLeaseExist(ctx context.Context, sandboxID string) error {
	leaseSvc := c.client.LeasesService()
	_, err := leaseSvc.Create(ctx, leases.WithID(sandboxID))
	if err != nil {
		if errdefs.IsAlreadyExists(err) {
			err = nil
		}
	}
	return err
}

func (c *criService) mutateImageMount(
	ctx context.Context,
	extraMount *runtime.Mount,
	snapshotter string,
	sandboxID string,
	platform imagespec.Platform,
) (retErr error) {
	imageSpec := extraMount.GetImage()
	if imageSpec == nil {
		return nil
	}
	if extraMount.GetHostPath() != "" {
		return fmt.Errorf("hostpath must be empty while mount image: %+v", extraMount)
	}
	// POC: Force all image volumes to be writable via overlay filesystem
	// TODO: Remove when Kubernetes API supports writable image volumes
	// Original check:
	// if !extraMount.GetReadonly() {
	// 	return fmt.Errorf("readonly must be true while mount image: %+v", extraMount)
	// }

	ref := imageSpec.GetImage()
	if ref == "" {
		return fmt.Errorf("image not specified in: %+v", imageSpec)
	}
	image, err := c.LocalResolve(ref)
	if err != nil {
		return fmt.Errorf("failed to resolve image %q: %w", ref, err)
	}
	containerdImage, err := c.toContainerdImage(ctx, image)
	if err != nil {
		return fmt.Errorf("failed to get image from containerd %q: %w", image.ID, err)
	}

	// This is a digest of the manifest
	imageID := containerdImage.Target().Digest.Encoded()

	// POC: Use overlay filesystem to make image volumes writable
	// Paths for overlay components
	target := c.getImageVolumeHostPath(sandboxID, imageID+"-overlay")
	lowerDir := c.getImageVolumeHostPath(sandboxID, imageID+"-lower")
	// Use /dev/shm for upper/work directories for in-memory performance
	upperDir := filepath.Join("/dev/shm/containerd-image-volumes", sandboxID, imageID+"-upper")
	workDir := filepath.Join("/dev/shm/containerd-image-volumes", sandboxID, imageID+"-work")

	// Already mounted in another container on the same pod
	mounted, err := ensureImageVolumeMounted(target)
	if err != nil {
		return fmt.Errorf("failed to ensure %s is mounted: %w", target, err)
	}
	if mounted {
		extraMount.HostPath = target
		// POC: Mark mount as writable
		extraMount.Readonly = false
		return nil
	}

	img, err := c.client.ImageService().Get(ctx, ref)
	if err != nil {
		return fmt.Errorf("failed to get image volume ref %q: %w", ref, err)
	}

	i := containerd.NewImageWithPlatform(c.client, img, platforms.Only(platform))
	if err := i.Unpack(ctx, snapshotter); err != nil {
		return fmt.Errorf("failed to unpack image volume: %w", err)
	}

	diffIDs, err := i.RootFS(ctx)
	if err != nil {
		return fmt.Errorf("failed to get diff IDs for image volume %q: %w", ref, err)
	}
	chainID := identity.ChainID(diffIDs).String()

	s := c.client.SnapshotService(snapshotter)
	
	// Prepare snapshot for lower directory with lowerDir as the key
	mounts, err := s.Prepare(ctx, lowerDir, chainID)
	if err != nil {
		if errdefs.IsAlreadyExists(err) {
			mounts, err = s.Mounts(ctx, lowerDir)
		}
	}
	if err != nil {
		return fmt.Errorf("failed to prepare for image volume %q: %w", ref, err)
	}
	defer func() {
		if retErr != nil {
			_ = s.Remove(ctx, lowerDir)
		}
	}()
	
	// Mount the snapshot to the lower layer (this puts the image content there)
	if err := os.MkdirAll(lowerDir, 0755); err != nil {
		return fmt.Errorf("failed to create lower dir %q: %w", lowerDir, err)
	}
	mounts = addVolatileOptionOnImageVolumeMount(mounts)
	log.G(ctx).Infof("POC DEBUG: About to mount snapshot to lower dir %s with %d mounts", lowerDir, len(mounts))
	if err := mount.All(mounts, lowerDir); err != nil {
		return fmt.Errorf("failed to mount lower layer %q: %w", lowerDir, err)
	}
	log.G(ctx).Infof("POC DEBUG: Successfully mounted lower layer %s", lowerDir)
	defer func() {
		if retErr != nil {
			_ = mount.UnmountAll(lowerDir, 0)
		}
	}()
	
	// Create upper and work directories in /dev/shm for in-memory performance
	log.G(ctx).Infof("POC DEBUG: Creating /dev/shm directories - upper: %s, work: %s", upperDir, workDir)
	if err := os.MkdirAll(upperDir, 0755); err != nil {
		return fmt.Errorf("failed to create upper dir %q: %w", upperDir, err)
	}
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return fmt.Errorf("failed to create work dir %q: %w", workDir, err)
	}
	log.G(ctx).Infof("POC DEBUG: Created /dev/shm directories successfully")
	defer func() {
		if retErr != nil {
			_ = os.RemoveAll(upperDir)
			_ = os.RemoveAll(workDir)
		}
	}()
	if err := os.MkdirAll(target, 0755); err != nil {
		return fmt.Errorf("failed to create target dir %q: %w", target, err)
	}
	
	// Mount overlay filesystem using /dev/shm directories
	overlayOpts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", lowerDir, upperDir, workDir)
	log.G(ctx).Infof("POC DEBUG: Mounting overlay with opts: %s", overlayOpts)
	overlayMount := mount.Mount{
		Type:    "overlay",
		Source:  "overlay",
		Options: []string{overlayOpts},
	}
	
	if err := overlayMount.Mount(target); err != nil {
		log.G(ctx).Errorf("POC DEBUG: Overlay mount failed - lower exists: %v, upper exists: %v, work exists: %v", 
			dirExists(lowerDir), dirExists(upperDir), dirExists(workDir))
		return fmt.Errorf("failed to mount writable overlay at %q: %w", target, err)
	}
	
	log.G(ctx).Infof("POC DEBUG: Successfully mounted overlay at %s", target)
	
	extraMount.HostPath = target
	// POC: Mark mount as writable
	log.G(ctx).Infof("POC DEBUG: Setting extraMount.Readonly = false (was %v)", extraMount.GetReadonly())
	extraMount.Readonly = false
	log.G(ctx).Infof("POC DEBUG: Final mount - HostPath: %s, Readonly: %v", extraMount.HostPath, extraMount.Readonly)
	return nil
}

func (c *criService) cleanupImageMounts(
	ctx context.Context,
	sandboxID string,
) (retErr error) {
	// Some checks to avoid affecting old pods.
	ociRuntime, err := c.getPodSandboxRuntime(sandboxID)
	if err != nil {
		log.G(ctx).WithError(err).Errorf("failed to get sandbox runtime handler %q", sandboxID)
		return nil
	}
	snapshotter := c.RuntimeSnapshotter(ctx, ociRuntime)
	s := c.client.SnapshotService(snapshotter)
	if s == nil {
		return nil
	}
	targetBase := c.getImageVolumeBaseDir(sandboxID)
	entries, err := os.ReadDir(targetBase)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		target := filepath.Join(targetBase, entry.Name())
		entryName := entry.Name()

		// Unmount the target (overlay)
		err = mount.UnmountAll(target, 0)
		if err != nil {
			log.G(ctx).WithError(err).Warnf("failed to unmount image volume component %q", target)
		}
		
		// Also unmount tmpfs upper and work directories
		if strings.HasSuffix(entryName, "-upper") || strings.HasSuffix(entryName, "-work") {
			err = mount.UnmountAll(target, 0)
			if err != nil {
				log.G(ctx).WithError(err).Debugf("failed to unmount tmpfs at %q", target)
			}
		}
		
		// POC: Handle snapshot cleanup for overlay setup
		// For lower directories, use the snapshot key format
		if strings.HasSuffix(entryName, "-lower") {
			imageID := strings.TrimSuffix(entryName, "-lower")
			snapshotKey := fmt.Sprintf("%s-lower-%s", sandboxID, imageID)
			err = s.Remove(ctx, snapshotKey)
			if err != nil && !errdefs.IsNotFound(err) {
				log.G(ctx).WithError(err).Debugf("failed to remove snapshot %q", snapshotKey)
			}
		}
		
		// Remove the directory
		err = os.RemoveAll(target)
		if err != nil && !os.IsNotExist(err) {
			log.G(ctx).WithError(err).Warnf("failed to remove directory %q", target)
		}
	}

	err = os.Remove(targetBase)
	if err != nil && !errdefs.IsNotFound(err) {
		return fmt.Errorf("failed to remove directory to cleanup image volume mounts: %w", err)
	}
	return nil
}

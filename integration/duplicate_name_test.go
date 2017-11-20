/*
Copyright 2017 The Kubernetes Authors.

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

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDuplicateName(t *testing.T) {
	t.Logf("Create a sandbox")
	sbConfig := PodSandboxConfig("sandbox", "duplicate-name")
	sb, err := runtimeService.RunPodSandbox(sbConfig)
	require.NoError(t, err)
	defer func() {
		assert.NoError(t, runtimeService.StopPodSandbox(sb))
		assert.NoError(t, runtimeService.RemovePodSandbox(sb))
	}()

	t.Logf("Create the sandbox again should fail")
	_, err = runtimeService.RunPodSandbox(sbConfig)
	require.Error(t, err)

	t.Logf("Create a container")
	cnConfig := ContainerConfig(
		"container",
		pauseImage,
	)
	_, err = runtimeService.CreateContainer(sb, cnConfig, sbConfig)
	require.NoError(t, err)

	t.Logf("Create the container again should fail")
	_, err = runtimeService.CreateContainer(sb, cnConfig, sbConfig)
	require.Error(t, err)
}

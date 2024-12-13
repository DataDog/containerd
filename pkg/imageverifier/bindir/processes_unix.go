//go:build !windows

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

package bindir

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"

	"golang.org/x/sys/unix"
)

type process struct {
	cmd  *exec.Cmd
	pgid int
}

// Configure the verifier command so that killing it kills all child
// processes of the verifier process.
func startProcess(ctx context.Context, cmd *exec.Cmd, s *string) (*process, error) {
	// Assign the verifier a new process group so that killing its process group
	// in Cancel() doesn't kill the parent process (containerd).
	cmd.SysProcAttr = &unix.SysProcAttr{Setpgid: true}

	cmd.Cancel = func() error {
		// Passing a negative PID causes kill(2) to kill all processes in the
		// process group whose ID is cmd.Process.Pid.
		*s += "+++KILL=" + strconv.Itoa(cmd.Process.Pid) + ":" + strconv.Itoa(cmd.SysProcAttr.Pgid) + "+++"
		return unix.Kill(-cmd.Process.Pid, unix.SIGKILL)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("starting process: %w", err)
	}

	pgid := -1
	if cmd.SysProcAttr != nil {
		pgid = cmd.SysProcAttr.Pgid
	}

	return &process{
		cmd:  cmd,
		pgid: pgid,
	}, nil
}

func (p *process) cleanup(ctx context.Context) {}

// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/runtime/task/v2/shim.proto
package task

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TaskService interface {
	State(context.Context, *StateRequest) (*StateResponse, error)
	Create(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Pids(context.Context, *PidsRequest) (*PidsResponse, error)
	Pause(context.Context, *PauseRequest) (*emptypb.Empty, error)
	Resume(context.Context, *ResumeRequest) (*emptypb.Empty, error)
	Checkpoint(context.Context, *CheckpointTaskRequest) (*emptypb.Empty, error)
	Kill(context.Context, *KillRequest) (*emptypb.Empty, error)
	Exec(context.Context, *ExecProcessRequest) (*emptypb.Empty, error)
	ResizePty(context.Context, *ResizePtyRequest) (*emptypb.Empty, error)
	CloseIO(context.Context, *CloseIORequest) (*emptypb.Empty, error)
	Update(context.Context, *UpdateTaskRequest) (*emptypb.Empty, error)
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*emptypb.Empty, error)
}

func RegisterTaskService(srv *ttrpc.Server, svc TaskService) {
	srv.RegisterService("containerd.task.v2.Task", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"State": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.State(ctx, &req)
			},
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CreateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Create(ctx, &req)
			},
			"Start": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StartRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Start(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req DeleteRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Delete(ctx, &req)
			},
			"Pids": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PidsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pids(ctx, &req)
			},
			"Pause": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PauseRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pause(ctx, &req)
			},
			"Resume": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResumeRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Resume(ctx, &req)
			},
			"Checkpoint": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CheckpointTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Checkpoint(ctx, &req)
			},
			"Kill": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req KillRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Kill(ctx, &req)
			},
			"Exec": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ExecProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Exec(ctx, &req)
			},
			"ResizePty": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResizePtyRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ResizePty(ctx, &req)
			},
			"CloseIO": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CloseIORequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.CloseIO(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req UpdateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
			"Wait": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req WaitRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Wait(ctx, &req)
			},
			"Stats": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StatsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Stats(ctx, &req)
			},
			"Connect": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ConnectRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Connect(ctx, &req)
			},
			"Shutdown": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ShutdownRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Shutdown(ctx, &req)
			},
		},
	})
}

type taskClient struct {
	client *ttrpc.Client
}

func NewTaskClient(client *ttrpc.Client) TaskService {
	return &taskClient{
		client: client,
	}
}

func (c *taskClient) State(ctx context.Context, req *StateRequest) (*StateResponse, error) {
	var resp StateResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "State", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Create(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var resp CreateTaskResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Pids(ctx context.Context, req *PidsRequest) (*PidsResponse, error) {
	var resp PidsResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Pids", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Pause(ctx context.Context, req *PauseRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Pause", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Resume(ctx context.Context, req *ResumeRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Resume", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Checkpoint(ctx context.Context, req *CheckpointTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Checkpoint", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Kill(ctx context.Context, req *KillRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Kill", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Exec(ctx context.Context, req *ExecProcessRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Exec", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) ResizePty(ctx context.Context, req *ResizePtyRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "ResizePty", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) CloseIO(ctx context.Context, req *CloseIORequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "CloseIO", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Update(ctx context.Context, req *UpdateTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Wait(ctx context.Context, req *WaitRequest) (*WaitResponse, error) {
	var resp WaitResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Wait", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Stats(ctx context.Context, req *StatsRequest) (*StatsResponse, error) {
	var resp StatsResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Stats", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Connect(ctx context.Context, req *ConnectRequest) (*ConnectResponse, error) {
	var resp ConnectResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Connect", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Shutdown(ctx context.Context, req *ShutdownRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Shutdown", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/runtime/task/v2/shim.proto
package task

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TaskService interface {
	State(context.Context, *StateRequest) (*StateResponse, error)
	Create(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Pids(context.Context, *PidsRequest) (*PidsResponse, error)
	Pause(context.Context, *PauseRequest) (*emptypb.Empty, error)
	Resume(context.Context, *ResumeRequest) (*emptypb.Empty, error)
	Checkpoint(context.Context, *CheckpointTaskRequest) (*emptypb.Empty, error)
	Kill(context.Context, *KillRequest) (*emptypb.Empty, error)
	Exec(context.Context, *ExecProcessRequest) (*emptypb.Empty, error)
	ResizePty(context.Context, *ResizePtyRequest) (*emptypb.Empty, error)
	CloseIO(context.Context, *CloseIORequest) (*emptypb.Empty, error)
	Update(context.Context, *UpdateTaskRequest) (*emptypb.Empty, error)
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*emptypb.Empty, error)
}

func RegisterTaskService(srv *ttrpc.Server, svc TaskService) {
	srv.RegisterService("containerd.task.v2.Task", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"State": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.State(ctx, &req)
			},
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CreateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Create(ctx, &req)
			},
			"Start": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StartRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Start(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req DeleteRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Delete(ctx, &req)
			},
			"Pids": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PidsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pids(ctx, &req)
			},
			"Pause": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PauseRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pause(ctx, &req)
			},
			"Resume": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResumeRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Resume(ctx, &req)
			},
			"Checkpoint": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CheckpointTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Checkpoint(ctx, &req)
			},
			"Kill": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req KillRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Kill(ctx, &req)
			},
			"Exec": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ExecProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Exec(ctx, &req)
			},
			"ResizePty": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResizePtyRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ResizePty(ctx, &req)
			},
			"CloseIO": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CloseIORequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.CloseIO(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req UpdateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
			"Wait": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req WaitRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Wait(ctx, &req)
			},
			"Stats": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StatsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Stats(ctx, &req)
			},
			"Connect": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ConnectRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Connect(ctx, &req)
			},
			"Shutdown": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ShutdownRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Shutdown(ctx, &req)
			},
		},
	})
}

type taskClient struct {
	client *ttrpc.Client
}

func NewTaskClient(client *ttrpc.Client) TaskService {
	return &taskClient{
		client: client,
	}
}

func (c *taskClient) State(ctx context.Context, req *StateRequest) (*StateResponse, error) {
	var resp StateResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "State", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Create(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var resp CreateTaskResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Pids(ctx context.Context, req *PidsRequest) (*PidsResponse, error) {
	var resp PidsResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Pids", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Pause(ctx context.Context, req *PauseRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Pause", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Resume(ctx context.Context, req *ResumeRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Resume", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Checkpoint(ctx context.Context, req *CheckpointTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Checkpoint", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Kill(ctx context.Context, req *KillRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Kill", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Exec(ctx context.Context, req *ExecProcessRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Exec", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) ResizePty(ctx context.Context, req *ResizePtyRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "ResizePty", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) CloseIO(ctx context.Context, req *CloseIORequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "CloseIO", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Update(ctx context.Context, req *UpdateTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Wait(ctx context.Context, req *WaitRequest) (*WaitResponse, error) {
	var resp WaitResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Wait", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Stats(ctx context.Context, req *StatsRequest) (*StatsResponse, error) {
	var resp StatsResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Stats", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Connect(ctx context.Context, req *ConnectRequest) (*ConnectResponse, error) {
	var resp ConnectResponse
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Connect", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *taskClient) Shutdown(ctx context.Context, req *ShutdownRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v2.Task", "Shutdown", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

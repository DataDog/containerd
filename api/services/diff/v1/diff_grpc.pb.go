// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/diff/v1/diff.proto

package diff

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DiffClient is the client API for Diff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiffClient interface {
	// Apply applies the content associated with the provided digests onto
	// the provided mounts. Archive content will be extracted and
	// decompressed if necessary.
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
	// Diff creates a diff between the given mounts and uploads the result
	// to the content store.
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
}

type diffClient struct {
	cc grpc.ClientConnInterface
}

func NewDiffClient(cc grpc.ClientConnInterface) DiffClient {
	return &diffClient{cc}
}

func (c *diffClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.diff.v1.Diff/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.diff.v1.Diff/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServer is the server API for Diff service.
// All implementations must embed UnimplementedDiffServer
// for forward compatibility
type DiffServer interface {
	// Apply applies the content associated with the provided digests onto
	// the provided mounts. Archive content will be extracted and
	// decompressed if necessary.
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
	// Diff creates a diff between the given mounts and uploads the result
	// to the content store.
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
	mustEmbedUnimplementedDiffServer()
}

// UnimplementedDiffServer must be embedded to have forward compatible implementations.
type UnimplementedDiffServer struct {
}

func (UnimplementedDiffServer) Apply(context.Context, *ApplyRequest) (*ApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedDiffServer) Diff(context.Context, *DiffRequest) (*DiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}
func (UnimplementedDiffServer) mustEmbedUnimplementedDiffServer() {}

// UnsafeDiffServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiffServer will
// result in compilation errors.
type UnsafeDiffServer interface {
	mustEmbedUnimplementedDiffServer()
}

func RegisterDiffServer(s grpc.ServiceRegistrar, srv DiffServer) {
	s.RegisterService(&Diff_ServiceDesc, srv)
}

func _Diff_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.diff.v1.Diff/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diff_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.diff.v1.Diff/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Diff_ServiceDesc is the grpc.ServiceDesc for Diff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.diff.v1.Diff",
	HandlerType: (*DiffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _Diff_Apply_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _Diff_Diff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/diff/v1/diff.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/diff/v1/diff.proto

package diff

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DiffClient is the client API for Diff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiffClient interface {
	// Apply applies the content associated with the provided digests onto
	// the provided mounts. Archive content will be extracted and
	// decompressed if necessary.
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
	// Diff creates a diff between the given mounts and uploads the result
	// to the content store.
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
}

type diffClient struct {
	cc grpc.ClientConnInterface
}

func NewDiffClient(cc grpc.ClientConnInterface) DiffClient {
	return &diffClient{cc}
}

func (c *diffClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.diff.v1.Diff/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.diff.v1.Diff/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServer is the server API for Diff service.
// All implementations must embed UnimplementedDiffServer
// for forward compatibility
type DiffServer interface {
	// Apply applies the content associated with the provided digests onto
	// the provided mounts. Archive content will be extracted and
	// decompressed if necessary.
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
	// Diff creates a diff between the given mounts and uploads the result
	// to the content store.
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
	mustEmbedUnimplementedDiffServer()
}

// UnimplementedDiffServer must be embedded to have forward compatible implementations.
type UnimplementedDiffServer struct {
}

func (UnimplementedDiffServer) Apply(context.Context, *ApplyRequest) (*ApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedDiffServer) Diff(context.Context, *DiffRequest) (*DiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}
func (UnimplementedDiffServer) mustEmbedUnimplementedDiffServer() {}

// UnsafeDiffServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiffServer will
// result in compilation errors.
type UnsafeDiffServer interface {
	mustEmbedUnimplementedDiffServer()
}

func RegisterDiffServer(s grpc.ServiceRegistrar, srv DiffServer) {
	s.RegisterService(&Diff_ServiceDesc, srv)
}

func _Diff_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.diff.v1.Diff/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diff_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.diff.v1.Diff/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Diff_ServiceDesc is the grpc.ServiceDesc for Diff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.diff.v1.Diff",
	HandlerType: (*DiffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _Diff_Apply_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _Diff_Diff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/diff/v1/diff.proto",
}

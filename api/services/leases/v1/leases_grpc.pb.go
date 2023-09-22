// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/leases/v1/leases.proto

package leases

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LeasesClient is the client API for Leases service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeasesClient interface {
	// Create creates a new lease for managing changes to metadata. A lease
	// can be used to protect objects from being removed.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Delete deletes the lease and makes any unreferenced objects created
	// during the lease eligible for garbage collection if not referenced
	// or retained by other resources during the lease.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List lists all active leases, returning the full list of
	// leases and optionally including the referenced resources.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// AddResource references the resource by the provided lease.
	AddResource(ctx context.Context, in *AddResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteResource dereferences the resource by the provided lease.
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListResources lists all the resources referenced by the lease.
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
}

type leasesClient struct {
	cc grpc.ClientConnInterface
}

func NewLeasesClient(cc grpc.ClientConnInterface) LeasesClient {
	return &leasesClient{cc}
}

func (c *leasesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) AddResource(ctx context.Context, in *AddResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/AddResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeasesServer is the server API for Leases service.
// All implementations must embed UnimplementedLeasesServer
// for forward compatibility
type LeasesServer interface {
	// Create creates a new lease for managing changes to metadata. A lease
	// can be used to protect objects from being removed.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Delete deletes the lease and makes any unreferenced objects created
	// during the lease eligible for garbage collection if not referenced
	// or retained by other resources during the lease.
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	// List lists all active leases, returning the full list of
	// leases and optionally including the referenced resources.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// AddResource references the resource by the provided lease.
	AddResource(context.Context, *AddResourceRequest) (*emptypb.Empty, error)
	// DeleteResource dereferences the resource by the provided lease.
	DeleteResource(context.Context, *DeleteResourceRequest) (*emptypb.Empty, error)
	// ListResources lists all the resources referenced by the lease.
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
	mustEmbedUnimplementedLeasesServer()
}

// UnimplementedLeasesServer must be embedded to have forward compatible implementations.
type UnimplementedLeasesServer struct {
}

func (UnimplementedLeasesServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLeasesServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLeasesServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLeasesServer) AddResource(context.Context, *AddResourceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddResource not implemented")
}
func (UnimplementedLeasesServer) DeleteResource(context.Context, *DeleteResourceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedLeasesServer) ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedLeasesServer) mustEmbedUnimplementedLeasesServer() {}

// UnsafeLeasesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeasesServer will
// result in compilation errors.
type UnsafeLeasesServer interface {
	mustEmbedUnimplementedLeasesServer()
}

func RegisterLeasesServer(s grpc.ServiceRegistrar, srv LeasesServer) {
	s.RegisterService(&Leases_ServiceDesc, srv)
}

func _Leases_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_AddResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).AddResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/AddResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).AddResource(ctx, req.(*AddResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Leases_ServiceDesc is the grpc.ServiceDesc for Leases service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leases_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.leases.v1.Leases",
	HandlerType: (*LeasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Leases_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Leases_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Leases_List_Handler,
		},
		{
			MethodName: "AddResource",
			Handler:    _Leases_AddResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _Leases_DeleteResource_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _Leases_ListResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/leases/v1/leases.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/leases/v1/leases.proto

package leases

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LeasesClient is the client API for Leases service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeasesClient interface {
	// Create creates a new lease for managing changes to metadata. A lease
	// can be used to protect objects from being removed.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Delete deletes the lease and makes any unreferenced objects created
	// during the lease eligible for garbage collection if not referenced
	// or retained by other resources during the lease.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List lists all active leases, returning the full list of
	// leases and optionally including the referenced resources.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// AddResource references the resource by the provided lease.
	AddResource(ctx context.Context, in *AddResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteResource dereferences the resource by the provided lease.
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListResources lists all the resources referenced by the lease.
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
}

type leasesClient struct {
	cc grpc.ClientConnInterface
}

func NewLeasesClient(cc grpc.ClientConnInterface) LeasesClient {
	return &leasesClient{cc}
}

func (c *leasesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) AddResource(ctx context.Context, in *AddResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/AddResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leasesClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.leases.v1.Leases/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeasesServer is the server API for Leases service.
// All implementations must embed UnimplementedLeasesServer
// for forward compatibility
type LeasesServer interface {
	// Create creates a new lease for managing changes to metadata. A lease
	// can be used to protect objects from being removed.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Delete deletes the lease and makes any unreferenced objects created
	// during the lease eligible for garbage collection if not referenced
	// or retained by other resources during the lease.
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	// List lists all active leases, returning the full list of
	// leases and optionally including the referenced resources.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// AddResource references the resource by the provided lease.
	AddResource(context.Context, *AddResourceRequest) (*emptypb.Empty, error)
	// DeleteResource dereferences the resource by the provided lease.
	DeleteResource(context.Context, *DeleteResourceRequest) (*emptypb.Empty, error)
	// ListResources lists all the resources referenced by the lease.
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
	mustEmbedUnimplementedLeasesServer()
}

// UnimplementedLeasesServer must be embedded to have forward compatible implementations.
type UnimplementedLeasesServer struct {
}

func (UnimplementedLeasesServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLeasesServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLeasesServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLeasesServer) AddResource(context.Context, *AddResourceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddResource not implemented")
}
func (UnimplementedLeasesServer) DeleteResource(context.Context, *DeleteResourceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedLeasesServer) ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedLeasesServer) mustEmbedUnimplementedLeasesServer() {}

// UnsafeLeasesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeasesServer will
// result in compilation errors.
type UnsafeLeasesServer interface {
	mustEmbedUnimplementedLeasesServer()
}

func RegisterLeasesServer(s grpc.ServiceRegistrar, srv LeasesServer) {
	s.RegisterService(&Leases_ServiceDesc, srv)
}

func _Leases_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_AddResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).AddResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/AddResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).AddResource(ctx, req.(*AddResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leases_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeasesServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.leases.v1.Leases/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeasesServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Leases_ServiceDesc is the grpc.ServiceDesc for Leases service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leases_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.leases.v1.Leases",
	HandlerType: (*LeasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Leases_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Leases_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Leases_List_Handler,
		},
		{
			MethodName: "AddResource",
			Handler:    _Leases_AddResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _Leases_DeleteResource_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _Leases_ListResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/leases/v1/leases.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: octoback/core/v1/service.proto

package corev1

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

const (
	Service_GetGroceryList_FullMethodName    = "/octoback.core.v1.Service/GetGroceryList"
	Service_CreateGroceryList_FullMethodName = "/octoback.core.v1.Service/CreateGroceryList"
	Service_UpdateGroceryList_FullMethodName = "/octoback.core.v1.Service/UpdateGroceryList"
	Service_ListGroceryLists_FullMethodName  = "/octoback.core.v1.Service/ListGroceryLists"
	Service_DeleteGroceryList_FullMethodName = "/octoback.core.v1.Service/DeleteGroceryList"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Get the grocery list.
	GetGroceryList(ctx context.Context, in *GetGroceryListRequest, opts ...grpc.CallOption) (*GetGroceryListResponse, error)
	// Create a grocery list
	CreateGroceryList(ctx context.Context, in *CreateGroceryListRequest, opts ...grpc.CallOption) (*CreateGroceryListResponse, error)
	// Update a grocery list
	UpdateGroceryList(ctx context.Context, in *UpdateGroceryListRequest, opts ...grpc.CallOption) (*UpdateGroceryListResponse, error)
	// List grocery lists
	ListGroceryLists(ctx context.Context, in *ListGroceryListsRequest, opts ...grpc.CallOption) (*ListGroceryListsResponse, error)
	// Delete a grocery list
	DeleteGroceryList(ctx context.Context, in *DeleteGroceryListRequest, opts ...grpc.CallOption) (*DeleteGroceryListResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetGroceryList(ctx context.Context, in *GetGroceryListRequest, opts ...grpc.CallOption) (*GetGroceryListResponse, error) {
	out := new(GetGroceryListResponse)
	err := c.cc.Invoke(ctx, Service_GetGroceryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateGroceryList(ctx context.Context, in *CreateGroceryListRequest, opts ...grpc.CallOption) (*CreateGroceryListResponse, error) {
	out := new(CreateGroceryListResponse)
	err := c.cc.Invoke(ctx, Service_CreateGroceryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateGroceryList(ctx context.Context, in *UpdateGroceryListRequest, opts ...grpc.CallOption) (*UpdateGroceryListResponse, error) {
	out := new(UpdateGroceryListResponse)
	err := c.cc.Invoke(ctx, Service_UpdateGroceryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListGroceryLists(ctx context.Context, in *ListGroceryListsRequest, opts ...grpc.CallOption) (*ListGroceryListsResponse, error) {
	out := new(ListGroceryListsResponse)
	err := c.cc.Invoke(ctx, Service_ListGroceryLists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteGroceryList(ctx context.Context, in *DeleteGroceryListRequest, opts ...grpc.CallOption) (*DeleteGroceryListResponse, error) {
	out := new(DeleteGroceryListResponse)
	err := c.cc.Invoke(ctx, Service_DeleteGroceryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Get the grocery list.
	GetGroceryList(context.Context, *GetGroceryListRequest) (*GetGroceryListResponse, error)
	// Create a grocery list
	CreateGroceryList(context.Context, *CreateGroceryListRequest) (*CreateGroceryListResponse, error)
	// Update a grocery list
	UpdateGroceryList(context.Context, *UpdateGroceryListRequest) (*UpdateGroceryListResponse, error)
	// List grocery lists
	ListGroceryLists(context.Context, *ListGroceryListsRequest) (*ListGroceryListsResponse, error)
	// Delete a grocery list
	DeleteGroceryList(context.Context, *DeleteGroceryListRequest) (*DeleteGroceryListResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GetGroceryList(context.Context, *GetGroceryListRequest) (*GetGroceryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroceryList not implemented")
}
func (UnimplementedServiceServer) CreateGroceryList(context.Context, *CreateGroceryListRequest) (*CreateGroceryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroceryList not implemented")
}
func (UnimplementedServiceServer) UpdateGroceryList(context.Context, *UpdateGroceryListRequest) (*UpdateGroceryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroceryList not implemented")
}
func (UnimplementedServiceServer) ListGroceryLists(context.Context, *ListGroceryListsRequest) (*ListGroceryListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroceryLists not implemented")
}
func (UnimplementedServiceServer) DeleteGroceryList(context.Context, *DeleteGroceryListRequest) (*DeleteGroceryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroceryList not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_GetGroceryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroceryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetGroceryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetGroceryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetGroceryList(ctx, req.(*GetGroceryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateGroceryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroceryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateGroceryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateGroceryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateGroceryList(ctx, req.(*CreateGroceryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateGroceryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroceryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateGroceryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateGroceryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateGroceryList(ctx, req.(*UpdateGroceryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListGroceryLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroceryListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListGroceryLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ListGroceryLists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListGroceryLists(ctx, req.(*ListGroceryListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteGroceryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroceryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteGroceryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteGroceryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteGroceryList(ctx, req.(*DeleteGroceryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "octoback.core.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroceryList",
			Handler:    _Service_GetGroceryList_Handler,
		},
		{
			MethodName: "CreateGroceryList",
			Handler:    _Service_CreateGroceryList_Handler,
		},
		{
			MethodName: "UpdateGroceryList",
			Handler:    _Service_UpdateGroceryList_Handler,
		},
		{
			MethodName: "ListGroceryLists",
			Handler:    _Service_ListGroceryLists_Handler,
		},
		{
			MethodName: "DeleteGroceryList",
			Handler:    _Service_DeleteGroceryList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "octoback/core/v1/service.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: supervisor_proto/supervisor.proto

package supervisor

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

// AssignClient is the client API for Assign service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssignClient interface {
	AssignTenant(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*Result, error)
	UnassignTenant(ctx context.Context, in *UnassignRequest, opts ...grpc.CallOption) (*Result, error)
	GetCurrentTenant(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTenantResponse, error)
}

type assignClient struct {
	cc grpc.ClientConnInterface
}

func NewAssignClient(cc grpc.ClientConnInterface) AssignClient {
	return &assignClient{cc}
}

func (c *assignClient) AssignTenant(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/supervisor.Assign/AssignTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignClient) UnassignTenant(ctx context.Context, in *UnassignRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/supervisor.Assign/UnassignTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assignClient) GetCurrentTenant(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, "/supervisor.Assign/GetCurrentTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssignServer is the server API for Assign service.
// All implementations must embed UnimplementedAssignServer
// for forward compatibility
type AssignServer interface {
	AssignTenant(context.Context, *AssignRequest) (*Result, error)
	UnassignTenant(context.Context, *UnassignRequest) (*Result, error)
	GetCurrentTenant(context.Context, *emptypb.Empty) (*GetTenantResponse, error)
	mustEmbedUnimplementedAssignServer()
}

// UnimplementedAssignServer must be embedded to have forward compatible implementations.
type UnimplementedAssignServer struct {
}

func (UnimplementedAssignServer) AssignTenant(context.Context, *AssignRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignTenant not implemented")
}
func (UnimplementedAssignServer) UnassignTenant(context.Context, *UnassignRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnassignTenant not implemented")
}
func (UnimplementedAssignServer) GetCurrentTenant(context.Context, *emptypb.Empty) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTenant not implemented")
}
func (UnimplementedAssignServer) mustEmbedUnimplementedAssignServer() {}

// UnsafeAssignServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssignServer will
// result in compilation errors.
type UnsafeAssignServer interface {
	mustEmbedUnimplementedAssignServer()
}

func RegisterAssignServer(s grpc.ServiceRegistrar, srv AssignServer) {
	s.RegisterService(&Assign_ServiceDesc, srv)
}

func _Assign_AssignTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignServer).AssignTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Assign/AssignTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignServer).AssignTenant(ctx, req.(*AssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assign_UnassignTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnassignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignServer).UnassignTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Assign/UnassignTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignServer).UnassignTenant(ctx, req.(*UnassignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assign_GetCurrentTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssignServer).GetCurrentTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.Assign/GetCurrentTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssignServer).GetCurrentTenant(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Assign_ServiceDesc is the grpc.ServiceDesc for Assign service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Assign_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.Assign",
	HandlerType: (*AssignServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignTenant",
			Handler:    _Assign_AssignTenant_Handler,
		},
		{
			MethodName: "UnassignTenant",
			Handler:    _Assign_UnassignTenant_Handler,
		},
		{
			MethodName: "GetCurrentTenant",
			Handler:    _Assign_GetCurrentTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor_proto/supervisor.proto",
}

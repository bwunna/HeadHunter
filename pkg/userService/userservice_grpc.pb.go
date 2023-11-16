// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: userservice.proto

package userService

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
	UserService_GetEmployeeByEmail_FullMethodName    = "/UserService/GetEmployeeByEmail"
	UserService_AddEmployee_FullMethodName           = "/UserService/AddEmployee"
	UserService_DeleteEmployeeByEmail_FullMethodName = "/UserService/DeleteEmployeeByEmail"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetEmployeeByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Employee, error)
	AddEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Basic, error)
	DeleteEmployeeByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Basic, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetEmployeeByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, UserService_GetEmployeeByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Basic, error) {
	out := new(Basic)
	err := c.cc.Invoke(ctx, UserService_AddEmployee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteEmployeeByEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*Basic, error) {
	out := new(Basic)
	err := c.cc.Invoke(ctx, UserService_DeleteEmployeeByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetEmployeeByEmail(context.Context, *EmailRequest) (*Employee, error)
	AddEmployee(context.Context, *Employee) (*Basic, error)
	DeleteEmployeeByEmail(context.Context, *EmailRequest) (*Basic, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetEmployeeByEmail(context.Context, *EmailRequest) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeByEmail not implemented")
}
func (UnimplementedUserServiceServer) AddEmployee(context.Context, *Employee) (*Basic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedUserServiceServer) DeleteEmployeeByEmail(context.Context, *EmailRequest) (*Basic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployeeByEmail not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetEmployeeByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetEmployeeByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetEmployeeByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetEmployeeByEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteEmployeeByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteEmployeeByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteEmployeeByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteEmployeeByEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployeeByEmail",
			Handler:    _UserService_GetEmployeeByEmail_Handler,
		},
		{
			MethodName: "AddEmployee",
			Handler:    _UserService_AddEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployeeByEmail",
			Handler:    _UserService_DeleteEmployeeByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice.proto",
}

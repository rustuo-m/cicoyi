// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc_service/user.proto

package pb

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
	UserService_GenerateTokenByName_FullMethodName    = "/user.UserService/GenerateTokenByName"
	UserService_GenerateTokenByEmail_FullMethodName   = "/user.UserService/GenerateTokenByEmail"
	UserService_SetPassword_FullMethodName            = "/user.UserService/SetPassword"
	UserService_GetUserListByCondition_FullMethodName = "/user.UserService/GetUserListByCondition"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 用户名密码注册
	GenerateTokenByName(ctx context.Context, in *RegisterByNameReq, opts ...grpc.CallOption) (*RegisterResponse, error)
	// 邮箱验证码登录
	GenerateTokenByEmail(ctx context.Context, in *RegisterByEmailCodeReq, opts ...grpc.CallOption) (*RegisterResponse, error)
	// 设置密码
	SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error)
	// 根据条件查询用户列表
	GetUserListByCondition(ctx context.Context, in *GetUserListByConditionReq, opts ...grpc.CallOption) (*GetUserListByConditionResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GenerateTokenByName(ctx context.Context, in *RegisterByNameReq, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_GenerateTokenByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GenerateTokenByEmail(ctx context.Context, in *RegisterByEmailCodeReq, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_GenerateTokenByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetPassword(ctx context.Context, in *SetPasswordReq, opts ...grpc.CallOption) (*SetPasswordResp, error) {
	out := new(SetPasswordResp)
	err := c.cc.Invoke(ctx, UserService_SetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserListByCondition(ctx context.Context, in *GetUserListByConditionReq, opts ...grpc.CallOption) (*GetUserListByConditionResp, error) {
	out := new(GetUserListByConditionResp)
	err := c.cc.Invoke(ctx, UserService_GetUserListByCondition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// 用户名密码注册
	GenerateTokenByName(context.Context, *RegisterByNameReq) (*RegisterResponse, error)
	// 邮箱验证码登录
	GenerateTokenByEmail(context.Context, *RegisterByEmailCodeReq) (*RegisterResponse, error)
	// 设置密码
	SetPassword(context.Context, *SetPasswordReq) (*SetPasswordResp, error)
	// 根据条件查询用户列表
	GetUserListByCondition(context.Context, *GetUserListByConditionReq) (*GetUserListByConditionResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GenerateTokenByName(context.Context, *RegisterByNameReq) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTokenByName not implemented")
}
func (UnimplementedUserServiceServer) GenerateTokenByEmail(context.Context, *RegisterByEmailCodeReq) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTokenByEmail not implemented")
}
func (UnimplementedUserServiceServer) SetPassword(context.Context, *SetPasswordReq) (*SetPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedUserServiceServer) GetUserListByCondition(context.Context, *GetUserListByConditionReq) (*GetUserListByConditionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserListByCondition not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GenerateTokenByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenerateTokenByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GenerateTokenByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenerateTokenByName(ctx, req.(*RegisterByNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GenerateTokenByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterByEmailCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenerateTokenByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GenerateTokenByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenerateTokenByEmail(ctx, req.(*RegisterByEmailCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetPassword(ctx, req.(*SetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserListByCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListByConditionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserListByCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserListByCondition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserListByCondition(ctx, req.(*GetUserListByConditionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTokenByName",
			Handler:    _UserService_GenerateTokenByName_Handler,
		},
		{
			MethodName: "GenerateTokenByEmail",
			Handler:    _UserService_GenerateTokenByEmail_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _UserService_SetPassword_Handler,
		},
		{
			MethodName: "GetUserListByCondition",
			Handler:    _UserService_GetUserListByCondition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_service/user.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: sso/sso.info.proto

package ssov1

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

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	User(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error)
	UserByID(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UserRoles(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRolesResponse, error)
	UserRolesByApp(ctx context.Context, in *UserRolesByAppRequest, opts ...grpc.CallOption) (*UserRolesByAppResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) User(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/info.UserInfo/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UserByID(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/info.UserInfo/UserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UserRoles(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRolesResponse, error) {
	out := new(UserRolesResponse)
	err := c.cc.Invoke(ctx, "/info.UserInfo/UserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UserRolesByApp(ctx context.Context, in *UserRolesByAppRequest, opts ...grpc.CallOption) (*UserRolesByAppResponse, error) {
	out := new(UserRolesByAppResponse)
	err := c.cc.Invoke(ctx, "/info.UserInfo/UserRolesByApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
// All implementations must embed UnimplementedUserInfoServer
// for forward compatibility
type UserInfoServer interface {
	User(context.Context, *emptypb.Empty) (*UserResponse, error)
	UserByID(context.Context, *UserRequest) (*UserResponse, error)
	UserRoles(context.Context, *UserRequest) (*UserRolesResponse, error)
	UserRolesByApp(context.Context, *UserRolesByAppRequest) (*UserRolesByAppResponse, error)
	mustEmbedUnimplementedUserInfoServer()
}

// UnimplementedUserInfoServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (UnimplementedUserInfoServer) User(context.Context, *emptypb.Empty) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedUserInfoServer) UserByID(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserByID not implemented")
}
func (UnimplementedUserInfoServer) UserRoles(context.Context, *UserRequest) (*UserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRoles not implemented")
}
func (UnimplementedUserInfoServer) UserRolesByApp(context.Context, *UserRolesByAppRequest) (*UserRolesByAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRolesByApp not implemented")
}
func (UnimplementedUserInfoServer) mustEmbedUnimplementedUserInfoServer() {}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc.ServiceRegistrar, srv UserInfoServer) {
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.UserInfo/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).User(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.UserInfo/UserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserByID(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.UserInfo/UserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserRoles(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UserRolesByApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRolesByAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserRolesByApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.UserInfo/UserRolesByApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserRolesByApp(ctx, req.(*UserRolesByAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "info.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "User",
			Handler:    _UserInfo_User_Handler,
		},
		{
			MethodName: "UserByID",
			Handler:    _UserInfo_UserByID_Handler,
		},
		{
			MethodName: "UserRoles",
			Handler:    _UserInfo_UserRoles_Handler,
		},
		{
			MethodName: "UserRolesByApp",
			Handler:    _UserInfo_UserRolesByApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.info.proto",
}

// AppInfoClient is the client API for AppInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppInfoClient interface {
	App(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AppResponse, error)
	AppByID(ctx context.Context, in *AppByIdRequest, opts ...grpc.CallOption) (*AppResponse, error)
}

type appInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewAppInfoClient(cc grpc.ClientConnInterface) AppInfoClient {
	return &appInfoClient{cc}
}

func (c *appInfoClient) App(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/info.AppInfo/App", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appInfoClient) AppByID(ctx context.Context, in *AppByIdRequest, opts ...grpc.CallOption) (*AppResponse, error) {
	out := new(AppResponse)
	err := c.cc.Invoke(ctx, "/info.AppInfo/AppByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppInfoServer is the server API for AppInfo service.
// All implementations must embed UnimplementedAppInfoServer
// for forward compatibility
type AppInfoServer interface {
	App(context.Context, *emptypb.Empty) (*AppResponse, error)
	AppByID(context.Context, *AppByIdRequest) (*AppResponse, error)
	mustEmbedUnimplementedAppInfoServer()
}

// UnimplementedAppInfoServer must be embedded to have forward compatible implementations.
type UnimplementedAppInfoServer struct {
}

func (UnimplementedAppInfoServer) App(context.Context, *emptypb.Empty) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method App not implemented")
}
func (UnimplementedAppInfoServer) AppByID(context.Context, *AppByIdRequest) (*AppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppByID not implemented")
}
func (UnimplementedAppInfoServer) mustEmbedUnimplementedAppInfoServer() {}

// UnsafeAppInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppInfoServer will
// result in compilation errors.
type UnsafeAppInfoServer interface {
	mustEmbedUnimplementedAppInfoServer()
}

func RegisterAppInfoServer(s grpc.ServiceRegistrar, srv AppInfoServer) {
	s.RegisterService(&AppInfo_ServiceDesc, srv)
}

func _AppInfo_App_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInfoServer).App(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.AppInfo/App",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInfoServer).App(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppInfo_AppByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppInfoServer).AppByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.AppInfo/AppByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppInfoServer).AppByID(ctx, req.(*AppByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppInfo_ServiceDesc is the grpc.ServiceDesc for AppInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "info.AppInfo",
	HandlerType: (*AppInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "App",
			Handler:    _AppInfo_App_Handler,
		},
		{
			MethodName: "AppByID",
			Handler:    _AppInfo_AppByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/sso.info.proto",
}

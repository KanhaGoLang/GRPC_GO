// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/proto.proto

package proto

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
	UserService_CreateUser_FullMethodName        = "/grpcService.UserService/CreateUser"
	UserService_ReadUserById_FullMethodName      = "/grpcService.UserService/ReadUserById"
	UserService_UpdateUser_FullMethodName        = "/grpcService.UserService/UpdateUser"
	UserService_DeleteUser_FullMethodName        = "/grpcService.UserService/DeleteUser"
	UserService_GetAllUsers_FullMethodName       = "/grpcService.UserService/GetAllUsers"
	UserService_GetUserPosts_FullMethodName      = "/grpcService.UserService/GetUserPosts"
	UserService_SaveMultipleUsers_FullMethodName = "/grpcService.UserService/SaveMultipleUsers"
	UserService_AuthUser_FullMethodName          = "/grpcService.UserService/AuthUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	ReadUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserSuccess, error)
	GetAllUsers(ctx context.Context, in *NoParameter, opts ...grpc.CallOption) (*Users, error)
	GetUserPosts(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Posts, error)
	// New client streaming RPC for saving multiple users
	SaveMultipleUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_SaveMultipleUsersClient, error)
	// auth user
	AuthUser(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ReadUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_ReadUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserSuccess, error) {
	out := new(UserSuccess)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *NoParameter, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, UserService_GetAllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserPosts(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Posts, error) {
	out := new(Posts)
	err := c.cc.Invoke(ctx, UserService_GetUserPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SaveMultipleUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_SaveMultipleUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], UserService_SaveMultipleUsers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceSaveMultipleUsersClient{stream}
	return x, nil
}

type UserService_SaveMultipleUsersClient interface {
	Send(*User) error
	CloseAndRecv() (*UserSuccess, error)
	grpc.ClientStream
}

type userServiceSaveMultipleUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceSaveMultipleUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceSaveMultipleUsersClient) CloseAndRecv() (*UserSuccess, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UserSuccess)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) AuthUser(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, UserService_AuthUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*User, error)
	ReadUserById(context.Context, *UserId) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *UserId) (*UserSuccess, error)
	GetAllUsers(context.Context, *NoParameter) (*Users, error)
	GetUserPosts(context.Context, *UserId) (*Posts, error)
	// New client streaming RPC for saving multiple users
	SaveMultipleUsers(UserService_SaveMultipleUsersServer) error
	// auth user
	AuthUser(context.Context, *AuthRequest) (*TokenResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ReadUserById(context.Context, *UserId) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUserById not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *UserId) (*UserSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *NoParameter) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUserPosts(context.Context, *UserId) (*Posts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPosts not implemented")
}
func (UnimplementedUserServiceServer) SaveMultipleUsers(UserService_SaveMultipleUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveMultipleUsers not implemented")
}
func (UnimplementedUserServiceServer) AuthUser(context.Context, *AuthRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUser not implemented")
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

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ReadUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ReadUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ReadUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ReadUserById(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*NoParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserPosts(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SaveMultipleUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).SaveMultipleUsers(&userServiceSaveMultipleUsersServer{stream})
}

type UserService_SaveMultipleUsersServer interface {
	SendAndClose(*UserSuccess) error
	Recv() (*User, error)
	grpc.ServerStream
}

type userServiceSaveMultipleUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceSaveMultipleUsersServer) SendAndClose(m *UserSuccess) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceSaveMultipleUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_AuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AuthUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AuthUser(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcService.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ReadUserById",
			Handler:    _UserService_ReadUserById_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "GetUserPosts",
			Handler:    _UserService_GetUserPosts_Handler,
		},
		{
			MethodName: "AuthUser",
			Handler:    _UserService_AuthUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveMultipleUsers",
			Handler:       _UserService_SaveMultipleUsers_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/proto.proto",
}

const (
	PostService_Create_FullMethodName   = "/grpcService.PostService/Create"
	PostService_ReadById_FullMethodName = "/grpcService.PostService/ReadById"
	PostService_Update_FullMethodName   = "/grpcService.PostService/Update"
	PostService_Delete_FullMethodName   = "/grpcService.PostService/Delete"
	PostService_GetAll_FullMethodName   = "/grpcService.PostService/GetAll"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	Create(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	ReadById(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error)
	Update(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	Delete(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*PostSuccess, error)
	GetAll(ctx context.Context, in *NoPostParameter, opts ...grpc.CallOption) (*Posts, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) Create(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ReadById(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_ReadById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Update(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) Delete(ctx context.Context, in *PostId, opts ...grpc.CallOption) (*PostSuccess, error) {
	out := new(PostSuccess)
	err := c.cc.Invoke(ctx, PostService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAll(ctx context.Context, in *NoPostParameter, opts ...grpc.CallOption) (*Posts, error) {
	out := new(Posts)
	err := c.cc.Invoke(ctx, PostService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	Create(context.Context, *Post) (*Post, error)
	ReadById(context.Context, *PostId) (*Post, error)
	Update(context.Context, *Post) (*Post, error)
	Delete(context.Context, *PostId) (*PostSuccess, error)
	GetAll(context.Context, *NoPostParameter) (*Posts, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) Create(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPostServiceServer) ReadById(context.Context, *PostId) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadById not implemented")
}
func (UnimplementedPostServiceServer) Update(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPostServiceServer) Delete(context.Context, *PostId) (*PostSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPostServiceServer) GetAll(context.Context, *NoPostParameter) (*Posts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Create(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ReadById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ReadById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ReadById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ReadById(ctx, req.(*PostId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Update(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Delete(ctx, req.(*PostId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoPostParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAll(ctx, req.(*NoPostParameter))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcService.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PostService_Create_Handler,
		},
		{
			MethodName: "ReadById",
			Handler:    _PostService_ReadById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PostService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PostService_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PostService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}

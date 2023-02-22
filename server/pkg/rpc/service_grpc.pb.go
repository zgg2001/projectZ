// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: service.proto

package rpc

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

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	LicencePlateCheck(ctx context.Context, in *LPCheckRequest, opts ...grpc.CallOption) (*LPCheckResponse, error)
	UploadParkingInfo(ctx context.Context, in *UploadInfoRequest, opts ...grpc.CallOption) (*UploadInfoResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) LicencePlateCheck(ctx context.Context, in *LPCheckRequest, opts ...grpc.CallOption) (*LPCheckResponse, error) {
	out := new(LPCheckResponse)
	err := c.cc.Invoke(ctx, "/ProjectService/LicencePlateCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UploadParkingInfo(ctx context.Context, in *UploadInfoRequest, opts ...grpc.CallOption) (*UploadInfoResponse, error) {
	out := new(UploadInfoResponse)
	err := c.cc.Invoke(ctx, "/ProjectService/UploadParkingInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/ProjectService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error) {
	out := new(GetUserDataResponse)
	err := c.cc.Invoke(ctx, "/ProjectService/GetUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	LicencePlateCheck(context.Context, *LPCheckRequest) (*LPCheckResponse, error)
	UploadParkingInfo(context.Context, *UploadInfoRequest) (*UploadInfoResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) LicencePlateCheck(context.Context, *LPCheckRequest) (*LPCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LicencePlateCheck not implemented")
}
func (UnimplementedProjectServiceServer) UploadParkingInfo(context.Context, *UploadInfoRequest) (*UploadInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadParkingInfo not implemented")
}
func (UnimplementedProjectServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedProjectServiceServer) GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_LicencePlateCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LPCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).LicencePlateCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectService/LicencePlateCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).LicencePlateCheck(ctx, req.(*LPCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UploadParkingInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UploadParkingInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectService/UploadParkingInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UploadParkingInfo(ctx, req.(*UploadInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProjectService/GetUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetUserData(ctx, req.(*GetUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LicencePlateCheck",
			Handler:    _ProjectService_LicencePlateCheck_Handler,
		},
		{
			MethodName: "UploadParkingInfo",
			Handler:    _ProjectService_UploadParkingInfo_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _ProjectService_UserLogin_Handler,
		},
		{
			MethodName: "GetUserData",
			Handler:    _ProjectService_GetUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

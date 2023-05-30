// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: account.proto

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	GetAccountList(ctx context.Context, in *PagingRequest, opts ...grpc.CallOption) (*AccountListRes, error)
	GetAccountByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*AccountRes, error)
	GetAccountByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*AccountRes, error)
	AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AccountRes, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountRes, error)
	CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*CheckPasswordRes, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetAccountList(ctx context.Context, in *PagingRequest, opts ...grpc.CallOption) (*AccountListRes, error) {
	out := new(AccountListRes)
	err := c.cc.Invoke(ctx, "/AccountService/GetAccountList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*AccountRes, error) {
	out := new(AccountRes)
	err := c.cc.Invoke(ctx, "/AccountService/GetAccountByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*AccountRes, error) {
	out := new(AccountRes)
	err := c.cc.Invoke(ctx, "/AccountService/GetAccountByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) AddAccount(ctx context.Context, in *AddAccountRequest, opts ...grpc.CallOption) (*AccountRes, error) {
	out := new(AccountRes)
	err := c.cc.Invoke(ctx, "/AccountService/AddAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountRes, error) {
	out := new(UpdateAccountRes)
	err := c.cc.Invoke(ctx, "/AccountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*CheckPasswordRes, error) {
	out := new(CheckPasswordRes)
	err := c.cc.Invoke(ctx, "/AccountService/CheckPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	GetAccountList(context.Context, *PagingRequest) (*AccountListRes, error)
	GetAccountByMobile(context.Context, *MobileRequest) (*AccountRes, error)
	GetAccountByID(context.Context, *IDRequest) (*AccountRes, error)
	AddAccount(context.Context, *AddAccountRequest) (*AccountRes, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountRes, error)
	CheckPassword(context.Context, *CheckPasswordRequest) (*CheckPasswordRes, error)
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) GetAccountList(context.Context, *PagingRequest) (*AccountListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountList not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByMobile(context.Context, *MobileRequest) (*AccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByMobile not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByID(context.Context, *IDRequest) (*AccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByID not implemented")
}
func (UnimplementedAccountServiceServer) AddAccount(context.Context, *AddAccountRequest) (*AccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (UnimplementedAccountServiceServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServiceServer) CheckPassword(context.Context, *CheckPasswordRequest) (*CheckPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_GetAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/GetAccountList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountList(ctx, req.(*PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/GetAccountByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/GetAccountByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/AddAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).AddAccount(ctx, req.(*AddAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CheckPassword(ctx, req.(*CheckPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountList",
			Handler:    _AccountService_GetAccountList_Handler,
		},
		{
			MethodName: "GetAccountByMobile",
			Handler:    _AccountService_GetAccountByMobile_Handler,
		},
		{
			MethodName: "GetAccountByID",
			Handler:    _AccountService_GetAccountByID_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _AccountService_AddAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _AccountService_CheckPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

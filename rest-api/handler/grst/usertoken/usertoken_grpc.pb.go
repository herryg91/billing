// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package usertoken

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UsertokenAPIClient is the client API for UsertokenAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsertokenAPIClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*UserToken, error)
	RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*UserToken, error)
	ValidateToken(ctx context.Context, in *ValidateTokenReq, opts ...grpc.CallOption) (*UserTokenClaim, error)
}

type usertokenAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUsertokenAPIClient(cc grpc.ClientConnInterface) UsertokenAPIClient {
	return &usertokenAPIClient{cc}
}

func (c *usertokenAPIClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*UserToken, error) {
	out := new(UserToken)
	err := c.cc.Invoke(ctx, "/usertoken.UsertokenAPI/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usertokenAPIClient) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*UserToken, error) {
	out := new(UserToken)
	err := c.cc.Invoke(ctx, "/usertoken.UsertokenAPI/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usertokenAPIClient) ValidateToken(ctx context.Context, in *ValidateTokenReq, opts ...grpc.CallOption) (*UserTokenClaim, error) {
	out := new(UserTokenClaim)
	err := c.cc.Invoke(ctx, "/usertoken.UsertokenAPI/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsertokenAPIServer is the server API for UsertokenAPI service.
// All implementations must embed UnimplementedUsertokenAPIServer
// for forward compatibility
type UsertokenAPIServer interface {
	GenerateToken(context.Context, *GenerateTokenReq) (*UserToken, error)
	RefreshToken(context.Context, *RefreshTokenReq) (*UserToken, error)
	ValidateToken(context.Context, *ValidateTokenReq) (*UserTokenClaim, error)
	mustEmbedUnimplementedUsertokenAPIServer()
}

// UnimplementedUsertokenAPIServer must be embedded to have forward compatible implementations.
type UnimplementedUsertokenAPIServer struct {
}

func (UnimplementedUsertokenAPIServer) GenerateToken(context.Context, *GenerateTokenReq) (*UserToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUsertokenAPIServer) RefreshToken(context.Context, *RefreshTokenReq) (*UserToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedUsertokenAPIServer) ValidateToken(context.Context, *ValidateTokenReq) (*UserTokenClaim, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedUsertokenAPIServer) mustEmbedUnimplementedUsertokenAPIServer() {}

// UnsafeUsertokenAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsertokenAPIServer will
// result in compilation errors.
type UnsafeUsertokenAPIServer interface {
	mustEmbedUnimplementedUsertokenAPIServer()
}

func RegisterUsertokenAPIServer(s *grpc.Server, srv UsertokenAPIServer) {
	s.RegisterService(&_UsertokenAPI_serviceDesc, srv)
}

func _UsertokenAPI_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsertokenAPIServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usertoken.UsertokenAPI/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsertokenAPIServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsertokenAPI_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsertokenAPIServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usertoken.UsertokenAPI/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsertokenAPIServer).RefreshToken(ctx, req.(*RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsertokenAPI_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsertokenAPIServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usertoken.UsertokenAPI/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsertokenAPIServer).ValidateToken(ctx, req.(*ValidateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsertokenAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usertoken.UsertokenAPI",
	HandlerType: (*UsertokenAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _UsertokenAPI_GenerateToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _UsertokenAPI_RefreshToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _UsertokenAPI_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usertoken.proto",
}

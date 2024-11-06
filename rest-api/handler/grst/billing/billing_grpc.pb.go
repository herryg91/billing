// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package billing

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BillingApiClient is the client API for BillingApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingApiClient interface {
	GetBillingByLoanCode(ctx context.Context, in *GetBillingByLoanCodeReq, opts ...grpc.CallOption) (*Billings, error)
	GetBillingOverDue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Billings, error)
	GenerateBillingPayment(ctx context.Context, in *GenerateBillingPaymentReq, opts ...grpc.CallOption) (*Billing, error)
	SettleBillingPayment(ctx context.Context, in *GenerateBillingPaymentReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type billingApiClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingApiClient(cc grpc.ClientConnInterface) BillingApiClient {
	return &billingApiClient{cc}
}

func (c *billingApiClient) GetBillingByLoanCode(ctx context.Context, in *GetBillingByLoanCodeReq, opts ...grpc.CallOption) (*Billings, error) {
	out := new(Billings)
	err := c.cc.Invoke(ctx, "/billing.BillingApi/GetBillingByLoanCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingApiClient) GetBillingOverDue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Billings, error) {
	out := new(Billings)
	err := c.cc.Invoke(ctx, "/billing.BillingApi/GetBillingOverDue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingApiClient) GenerateBillingPayment(ctx context.Context, in *GenerateBillingPaymentReq, opts ...grpc.CallOption) (*Billing, error) {
	out := new(Billing)
	err := c.cc.Invoke(ctx, "/billing.BillingApi/GenerateBillingPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingApiClient) SettleBillingPayment(ctx context.Context, in *GenerateBillingPaymentReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/billing.BillingApi/SettleBillingPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingApiServer is the server API for BillingApi service.
// All implementations must embed UnimplementedBillingApiServer
// for forward compatibility
type BillingApiServer interface {
	GetBillingByLoanCode(context.Context, *GetBillingByLoanCodeReq) (*Billings, error)
	GetBillingOverDue(context.Context, *empty.Empty) (*Billings, error)
	GenerateBillingPayment(context.Context, *GenerateBillingPaymentReq) (*Billing, error)
	SettleBillingPayment(context.Context, *GenerateBillingPaymentReq) (*empty.Empty, error)
	mustEmbedUnimplementedBillingApiServer()
}

// UnimplementedBillingApiServer must be embedded to have forward compatible implementations.
type UnimplementedBillingApiServer struct {
}

func (UnimplementedBillingApiServer) GetBillingByLoanCode(context.Context, *GetBillingByLoanCodeReq) (*Billings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingByLoanCode not implemented")
}
func (UnimplementedBillingApiServer) GetBillingOverDue(context.Context, *empty.Empty) (*Billings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBillingOverDue not implemented")
}
func (UnimplementedBillingApiServer) GenerateBillingPayment(context.Context, *GenerateBillingPaymentReq) (*Billing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateBillingPayment not implemented")
}
func (UnimplementedBillingApiServer) SettleBillingPayment(context.Context, *GenerateBillingPaymentReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SettleBillingPayment not implemented")
}
func (UnimplementedBillingApiServer) mustEmbedUnimplementedBillingApiServer() {}

// UnsafeBillingApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingApiServer will
// result in compilation errors.
type UnsafeBillingApiServer interface {
	mustEmbedUnimplementedBillingApiServer()
}

func RegisterBillingApiServer(s *grpc.Server, srv BillingApiServer) {
	s.RegisterService(&_BillingApi_serviceDesc, srv)
}

func _BillingApi_GetBillingByLoanCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingByLoanCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingApiServer).GetBillingByLoanCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.BillingApi/GetBillingByLoanCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingApiServer).GetBillingByLoanCode(ctx, req.(*GetBillingByLoanCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingApi_GetBillingOverDue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingApiServer).GetBillingOverDue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.BillingApi/GetBillingOverDue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingApiServer).GetBillingOverDue(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingApi_GenerateBillingPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateBillingPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingApiServer).GenerateBillingPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.BillingApi/GenerateBillingPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingApiServer).GenerateBillingPayment(ctx, req.(*GenerateBillingPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingApi_SettleBillingPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateBillingPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingApiServer).SettleBillingPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.BillingApi/SettleBillingPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingApiServer).SettleBillingPayment(ctx, req.(*GenerateBillingPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BillingApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.BillingApi",
	HandlerType: (*BillingApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillingByLoanCode",
			Handler:    _BillingApi_GetBillingByLoanCode_Handler,
		},
		{
			MethodName: "GetBillingOverDue",
			Handler:    _BillingApi_GetBillingOverDue_Handler,
		},
		{
			MethodName: "GenerateBillingPayment",
			Handler:    _BillingApi_GenerateBillingPayment_Handler,
		},
		{
			MethodName: "SettleBillingPayment",
			Handler:    _BillingApi_SettleBillingPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "billing.proto",
}

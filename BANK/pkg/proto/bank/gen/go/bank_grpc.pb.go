// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: bank.proto

package frgaed_bank_v1

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
	Bank_RequestForPaymentProcessing_FullMethodName = "/bank.Bank/RequestForPaymentProcessing"
	Bank_RequestForPaymentStatus_FullMethodName     = "/bank.Bank/RequestForPaymentStatus"
)

// BankClient is the client API for Bank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankClient interface {
	RequestForPaymentProcessing(ctx context.Context, in *RFPPRequest, opts ...grpc.CallOption) (*RFPPResponse, error)
	RequestForPaymentStatus(ctx context.Context, in *RFPSRequest, opts ...grpc.CallOption) (*RFPSResponse, error)
}

type bankClient struct {
	cc grpc.ClientConnInterface
}

func NewBankClient(cc grpc.ClientConnInterface) BankClient {
	return &bankClient{cc}
}

func (c *bankClient) RequestForPaymentProcessing(ctx context.Context, in *RFPPRequest, opts ...grpc.CallOption) (*RFPPResponse, error) {
	out := new(RFPPResponse)
	err := c.cc.Invoke(ctx, Bank_RequestForPaymentProcessing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) RequestForPaymentStatus(ctx context.Context, in *RFPSRequest, opts ...grpc.CallOption) (*RFPSResponse, error) {
	out := new(RFPSResponse)
	err := c.cc.Invoke(ctx, Bank_RequestForPaymentStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServer is the server API for Bank service.
// All implementations must embed UnimplementedBankServer
// for forward compatibility
type BankServer interface {
	RequestForPaymentProcessing(context.Context, *RFPPRequest) (*RFPPResponse, error)
	RequestForPaymentStatus(context.Context, *RFPSRequest) (*RFPSResponse, error)
	mustEmbedUnimplementedBankServer()
}

// UnimplementedBankServer must be embedded to have forward compatible implementations.
type UnimplementedBankServer struct {
}

func (UnimplementedBankServer) RequestForPaymentProcessing(context.Context, *RFPPRequest) (*RFPPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestForPaymentProcessing not implemented")
}
func (UnimplementedBankServer) RequestForPaymentStatus(context.Context, *RFPSRequest) (*RFPSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestForPaymentStatus not implemented")
}
func (UnimplementedBankServer) mustEmbedUnimplementedBankServer() {}

// UnsafeBankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServer will
// result in compilation errors.
type UnsafeBankServer interface {
	mustEmbedUnimplementedBankServer()
}

func RegisterBankServer(s grpc.ServiceRegistrar, srv BankServer) {
	s.RegisterService(&Bank_ServiceDesc, srv)
}

func _Bank_RequestForPaymentProcessing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RFPPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).RequestForPaymentProcessing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_RequestForPaymentProcessing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).RequestForPaymentProcessing(ctx, req.(*RFPPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_RequestForPaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RFPSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).RequestForPaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bank_RequestForPaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).RequestForPaymentStatus(ctx, req.(*RFPSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bank_ServiceDesc is the grpc.ServiceDesc for Bank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.Bank",
	HandlerType: (*BankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestForPaymentProcessing",
			Handler:    _Bank_RequestForPaymentProcessing_Handler,
		},
		{
			MethodName: "RequestForPaymentStatus",
			Handler:    _Bank_RequestForPaymentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bank.proto",
}

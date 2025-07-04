// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: transfer/transfer.proto

package transferrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Transfer_Transfer_FullMethodName                 = "/balance.Transfer/Transfer"
	Transfer_GetPublicKey_FullMethodName             = "/balance.Transfer/GetPublicKey"
	Transfer_GetMinimumBalanceForRent_FullMethodName = "/balance.Transfer/GetMinimumBalanceForRent"
)

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetPublicKey(ctx context.Context, in *PrivateKeyRequest, opts ...grpc.CallOption) (*PublicKeyResponse, error)
	GetMinimumBalanceForRent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMinimumBalanceForRentResponse, error)
}

type transferClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferClient(cc grpc.ClientConnInterface) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, Transfer_Transfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferClient) GetPublicKey(ctx context.Context, in *PrivateKeyRequest, opts ...grpc.CallOption) (*PublicKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicKeyResponse)
	err := c.cc.Invoke(ctx, Transfer_GetPublicKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferClient) GetMinimumBalanceForRent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMinimumBalanceForRentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMinimumBalanceForRentResponse)
	err := c.cc.Invoke(ctx, Transfer_GetMinimumBalanceForRent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServer is the server API for Transfer service.
// All implementations must embed UnimplementedTransferServer
// for forward compatibility.
type TransferServer interface {
	Transfer(context.Context, *TransferRequest) (*TransactionResponse, error)
	GetPublicKey(context.Context, *PrivateKeyRequest) (*PublicKeyResponse, error)
	GetMinimumBalanceForRent(context.Context, *emptypb.Empty) (*GetMinimumBalanceForRentResponse, error)
	mustEmbedUnimplementedTransferServer()
}

// UnimplementedTransferServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransferServer struct{}

func (UnimplementedTransferServer) Transfer(context.Context, *TransferRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedTransferServer) GetPublicKey(context.Context, *PrivateKeyRequest) (*PublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedTransferServer) GetMinimumBalanceForRent(context.Context, *emptypb.Empty) (*GetMinimumBalanceForRentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinimumBalanceForRent not implemented")
}
func (UnimplementedTransferServer) mustEmbedUnimplementedTransferServer() {}
func (UnimplementedTransferServer) testEmbeddedByValue()                  {}

// UnsafeTransferServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServer will
// result in compilation errors.
type UnsafeTransferServer interface {
	mustEmbedUnimplementedTransferServer()
}

func RegisterTransferServer(s grpc.ServiceRegistrar, srv TransferServer) {
	// If the following call pancis, it indicates UnimplementedTransferServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Transfer_ServiceDesc, srv)
}

func _Transfer_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfer_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfer_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfer_GetPublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServer).GetPublicKey(ctx, req.(*PrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transfer_GetMinimumBalanceForRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServer).GetMinimumBalanceForRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transfer_GetMinimumBalanceForRent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServer).GetMinimumBalanceForRent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Transfer_ServiceDesc is the grpc.ServiceDesc for Transfer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transfer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "balance.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Transfer_Transfer_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _Transfer_GetPublicKey_Handler,
		},
		{
			MethodName: "GetMinimumBalanceForRent",
			Handler:    _Transfer_GetMinimumBalanceForRent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer/transfer.proto",
}

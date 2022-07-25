// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nova/gal/v1/tx.proto

package galv1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
	Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error)
	Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error)
	PendingUndelegateRecord(ctx context.Context, in *MsgPendingUndelegateRecord, opts ...grpc.CallOption) (*MsgPendingUndelegateRecordResponse, error)
	Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
	ClaimSnAsset(ctx context.Context, in *MsgClaimSnAsset, opts ...grpc.CallOption) (*MsgClaimSnAssetResponse, error)
	PendingWithdraw(ctx context.Context, in *MsgPendingWithdraw, opts ...grpc.CallOption) (*MsgPendingWithdrawResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, "/nova.gal.v1.Msg/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Delegate(ctx context.Context, in *MsgDelegate, opts ...grpc.CallOption) (*MsgDelegateResponse, error) {
	out := new(MsgDelegateResponse)
	err := c.cc.Invoke(ctx, "/nova.gal.v1.Msg/Delegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Undelegate(ctx context.Context, in *MsgUndelegate, opts ...grpc.CallOption) (*MsgUndelegateResponse, error) {
	out := new(MsgUndelegateResponse)
	err := c.cc.Invoke(ctx, "/nova.gal.v1.Msg/Undelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PendingUndelegateRecord(ctx context.Context, in *MsgPendingUndelegateRecord, opts ...grpc.CallOption) (*MsgPendingUndelegateRecordResponse, error) {
	out := new(MsgPendingUndelegateRecordResponse)
	err := c.cc.Invoke(ctx, "/nova.gal.v1.Msg/PendingUndelegateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, "/nova.gal.v1.Msg/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimSnAsset(ctx context.Context, in *MsgClaimSnAsset, opts ...grpc.CallOption) (*MsgClaimSnAssetResponse, error) {
	out := new(MsgClaimSnAssetResponse)
	err := c.cc.Invoke(ctx, "/nova.gal.v1.Msg/ClaimSnAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PendingWithdraw(ctx context.Context, in *MsgPendingWithdraw, opts ...grpc.CallOption) (*MsgPendingWithdrawResponse, error) {
	out := new(MsgPendingWithdrawResponse)
	err := c.cc.Invoke(ctx, "/nova.gal.v1.Msg/PendingWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error)
	Delegate(context.Context, *MsgDelegate) (*MsgDelegateResponse, error)
	Undelegate(context.Context, *MsgUndelegate) (*MsgUndelegateResponse, error)
	PendingUndelegateRecord(context.Context, *MsgPendingUndelegateRecord) (*MsgPendingUndelegateRecordResponse, error)
	Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error)
	ClaimSnAsset(context.Context, *MsgClaimSnAsset) (*MsgClaimSnAssetResponse, error)
	PendingWithdraw(context.Context, *MsgPendingWithdraw) (*MsgPendingWithdrawResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedMsgServer) Delegate(context.Context, *MsgDelegate) (*MsgDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delegate not implemented")
}
func (UnimplementedMsgServer) Undelegate(context.Context, *MsgUndelegate) (*MsgUndelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Undelegate not implemented")
}
func (UnimplementedMsgServer) PendingUndelegateRecord(context.Context, *MsgPendingUndelegateRecord) (*MsgPendingUndelegateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingUndelegateRecord not implemented")
}
func (UnimplementedMsgServer) Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedMsgServer) ClaimSnAsset(context.Context, *MsgClaimSnAsset) (*MsgClaimSnAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimSnAsset not implemented")
}
func (UnimplementedMsgServer) PendingWithdraw(context.Context, *MsgPendingWithdraw) (*MsgPendingWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingWithdraw not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.gal.v1.Msg/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Delegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Delegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.gal.v1.Msg/Delegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Delegate(ctx, req.(*MsgDelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Undelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUndelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Undelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.gal.v1.Msg/Undelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Undelegate(ctx, req.(*MsgUndelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PendingUndelegateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPendingUndelegateRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PendingUndelegateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.gal.v1.Msg/PendingUndelegateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PendingUndelegateRecord(ctx, req.(*MsgPendingUndelegateRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdraw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.gal.v1.Msg/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Withdraw(ctx, req.(*MsgWithdraw))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimSnAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimSnAsset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimSnAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.gal.v1.Msg/ClaimSnAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimSnAsset(ctx, req.(*MsgClaimSnAsset))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PendingWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPendingWithdraw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PendingWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.gal.v1.Msg/PendingWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PendingWithdraw(ctx, req.(*MsgPendingWithdraw))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nova.gal.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
		{
			MethodName: "Delegate",
			Handler:    _Msg_Delegate_Handler,
		},
		{
			MethodName: "Undelegate",
			Handler:    _Msg_Undelegate_Handler,
		},
		{
			MethodName: "PendingUndelegateRecord",
			Handler:    _Msg_PendingUndelegateRecord_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Msg_Withdraw_Handler,
		},
		{
			MethodName: "ClaimSnAsset",
			Handler:    _Msg_ClaimSnAsset_Handler,
		},
		{
			MethodName: "PendingWithdraw",
			Handler:    _Msg_PendingWithdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nova/gal/v1/tx.proto",
}

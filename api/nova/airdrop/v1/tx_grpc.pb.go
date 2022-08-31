// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nova/airdrop/v1/tx.proto

package airdropv1

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
	ClaimAirdrop(ctx context.Context, in *MsgClaimAirdropRequest, opts ...grpc.CallOption) (*MsgClaimAirdropResponse, error)
	MarkSocialQuestPerformed(ctx context.Context, in *MsgMarkSocialQuestPerformedRequest, opts ...grpc.CallOption) (*MsgMarkSocialQuestPerformedResponse, error)
	MarkUserProvidedLiquidity(ctx context.Context, in *MsgMarkUserProvidedLiquidityRequest, opts ...grpc.CallOption) (*MsgMarkUserProvidedLiquidityResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ClaimAirdrop(ctx context.Context, in *MsgClaimAirdropRequest, opts ...grpc.CallOption) (*MsgClaimAirdropResponse, error) {
	out := new(MsgClaimAirdropResponse)
	err := c.cc.Invoke(ctx, "/nova.airdrop.v1.Msg/ClaimAirdrop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MarkSocialQuestPerformed(ctx context.Context, in *MsgMarkSocialQuestPerformedRequest, opts ...grpc.CallOption) (*MsgMarkSocialQuestPerformedResponse, error) {
	out := new(MsgMarkSocialQuestPerformedResponse)
	err := c.cc.Invoke(ctx, "/nova.airdrop.v1.Msg/MarkSocialQuestPerformed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MarkUserProvidedLiquidity(ctx context.Context, in *MsgMarkUserProvidedLiquidityRequest, opts ...grpc.CallOption) (*MsgMarkUserProvidedLiquidityResponse, error) {
	out := new(MsgMarkUserProvidedLiquidityResponse)
	err := c.cc.Invoke(ctx, "/nova.airdrop.v1.Msg/MarkUserProvidedLiquidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	ClaimAirdrop(context.Context, *MsgClaimAirdropRequest) (*MsgClaimAirdropResponse, error)
	MarkSocialQuestPerformed(context.Context, *MsgMarkSocialQuestPerformedRequest) (*MsgMarkSocialQuestPerformedResponse, error)
	MarkUserProvidedLiquidity(context.Context, *MsgMarkUserProvidedLiquidityRequest) (*MsgMarkUserProvidedLiquidityResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ClaimAirdrop(context.Context, *MsgClaimAirdropRequest) (*MsgClaimAirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimAirdrop not implemented")
}
func (UnimplementedMsgServer) MarkSocialQuestPerformed(context.Context, *MsgMarkSocialQuestPerformedRequest) (*MsgMarkSocialQuestPerformedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkSocialQuestPerformed not implemented")
}
func (UnimplementedMsgServer) MarkUserProvidedLiquidity(context.Context, *MsgMarkUserProvidedLiquidityRequest) (*MsgMarkUserProvidedLiquidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkUserProvidedLiquidity not implemented")
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

func _Msg_ClaimAirdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimAirdropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimAirdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.airdrop.v1.Msg/ClaimAirdrop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimAirdrop(ctx, req.(*MsgClaimAirdropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MarkSocialQuestPerformed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMarkSocialQuestPerformedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MarkSocialQuestPerformed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.airdrop.v1.Msg/MarkSocialQuestPerformed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MarkSocialQuestPerformed(ctx, req.(*MsgMarkSocialQuestPerformedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MarkUserProvidedLiquidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMarkUserProvidedLiquidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MarkUserProvidedLiquidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.airdrop.v1.Msg/MarkUserProvidedLiquidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MarkUserProvidedLiquidity(ctx, req.(*MsgMarkUserProvidedLiquidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nova.airdrop.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimAirdrop",
			Handler:    _Msg_ClaimAirdrop_Handler,
		},
		{
			MethodName: "MarkSocialQuestPerformed",
			Handler:    _Msg_MarkSocialQuestPerformed_Handler,
		},
		{
			MethodName: "MarkUserProvidedLiquidity",
			Handler:    _Msg_MarkUserProvidedLiquidity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nova/airdrop/v1/tx.proto",
}

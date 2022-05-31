// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nova/intertx/v1/tx.proto

package intertxv1

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
	// Register defines a rpc handler for MsgRegisterZone
	RegisterZone(ctx context.Context, in *MsgRegisterZone, opts ...grpc.CallOption) (*MsgRegisterZoneResponse, error)
	// IcaDelegate defines a rpc handler for MsgIcaDelegate
	IcaDelegate(ctx context.Context, in *MsgIcaDelegate, opts ...grpc.CallOption) (*MsgIcaDelegateResponse, error)
	// IcaUnDelegate defines a rpc handler for MsgIcaUnDelegate
	IcaUndelegate(ctx context.Context, in *MsgIcaUndelegate, opts ...grpc.CallOption) (*MsgIcaUndelegateResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterZone(ctx context.Context, in *MsgRegisterZone, opts ...grpc.CallOption) (*MsgRegisterZoneResponse, error) {
	out := new(MsgRegisterZoneResponse)
	err := c.cc.Invoke(ctx, "/nova.intertx.v1.Msg/RegisterZone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) IcaDelegate(ctx context.Context, in *MsgIcaDelegate, opts ...grpc.CallOption) (*MsgIcaDelegateResponse, error) {
	out := new(MsgIcaDelegateResponse)
	err := c.cc.Invoke(ctx, "/nova.intertx.v1.Msg/IcaDelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) IcaUndelegate(ctx context.Context, in *MsgIcaUndelegate, opts ...grpc.CallOption) (*MsgIcaUndelegateResponse, error) {
	out := new(MsgIcaUndelegateResponse)
	err := c.cc.Invoke(ctx, "/nova.intertx.v1.Msg/IcaUndelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// Register defines a rpc handler for MsgRegisterZone
	RegisterZone(context.Context, *MsgRegisterZone) (*MsgRegisterZoneResponse, error)
	// IcaDelegate defines a rpc handler for MsgIcaDelegate
	IcaDelegate(context.Context, *MsgIcaDelegate) (*MsgIcaDelegateResponse, error)
	// IcaUnDelegate defines a rpc handler for MsgIcaUnDelegate
	IcaUndelegate(context.Context, *MsgIcaUndelegate) (*MsgIcaUndelegateResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RegisterZone(context.Context, *MsgRegisterZone) (*MsgRegisterZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterZone not implemented")
}
func (UnimplementedMsgServer) IcaDelegate(context.Context, *MsgIcaDelegate) (*MsgIcaDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IcaDelegate not implemented")
}
func (UnimplementedMsgServer) IcaUndelegate(context.Context, *MsgIcaUndelegate) (*MsgIcaUndelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IcaUndelegate not implemented")
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

func _Msg_RegisterZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterZone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.intertx.v1.Msg/RegisterZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterZone(ctx, req.(*MsgRegisterZone))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_IcaDelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIcaDelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IcaDelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.intertx.v1.Msg/IcaDelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IcaDelegate(ctx, req.(*MsgIcaDelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_IcaUndelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIcaUndelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IcaUndelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.intertx.v1.Msg/IcaUndelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IcaUndelegate(ctx, req.(*MsgIcaUndelegate))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nova.intertx.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterZone",
			Handler:    _Msg_RegisterZone_Handler,
		},
		{
			MethodName: "IcaDelegate",
			Handler:    _Msg_IcaDelegate_Handler,
		},
		{
			MethodName: "IcaUndelegate",
			Handler:    _Msg_IcaUndelegate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nova/intertx/v1/tx.proto",
}

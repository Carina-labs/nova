// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/oracle/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgUpdateChainState struct {
	// coin refers to the sum of owned, staked and claimable quantity of the coin
	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coin" yaml:"coin"`
	// decimal of the native currency in host chain.
	Decimal uint32 `protobuf:"varint,2,opt,name=decimal,proto3" json:"decimal,omitempty" yaml:"decimal"`
	// address of the oracle
	Operator string `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty" yaml:"operator"`
	// block_height of the block fetched by oracle from host_chain
	BlockHeight int64 `protobuf:"varint,4,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty" yaml:"block_height"`
	// app_hash of the block fetched by oracle from host chain
	AppHash []byte `protobuf:"bytes,5,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty" yaml:"app_hash"`
	// chain_id of the host chain
	ChainId string `protobuf:"bytes,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
}

func (m *MsgUpdateChainState) Reset()         { *m = MsgUpdateChainState{} }
func (m *MsgUpdateChainState) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateChainState) ProtoMessage()    {}
func (*MsgUpdateChainState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ebac1664c7de357, []int{0}
}
func (m *MsgUpdateChainState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateChainState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateChainState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateChainState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateChainState.Merge(m, src)
}
func (m *MsgUpdateChainState) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateChainState) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateChainState.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateChainState proto.InternalMessageInfo

func (m *MsgUpdateChainState) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *MsgUpdateChainState) GetDecimal() uint32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *MsgUpdateChainState) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *MsgUpdateChainState) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MsgUpdateChainState) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *MsgUpdateChainState) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type MsgUpdateChainStateResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty" yaml:"success"`
}

func (m *MsgUpdateChainStateResponse) Reset()         { *m = MsgUpdateChainStateResponse{} }
func (m *MsgUpdateChainStateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateChainStateResponse) ProtoMessage()    {}
func (*MsgUpdateChainStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ebac1664c7de357, []int{1}
}
func (m *MsgUpdateChainStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateChainStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateChainStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateChainStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateChainStateResponse.Merge(m, src)
}
func (m *MsgUpdateChainStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateChainStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateChainStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateChainStateResponse proto.InternalMessageInfo

func (m *MsgUpdateChainStateResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*MsgUpdateChainState)(nil), "nova.oracle.v1.MsgUpdateChainState")
	proto.RegisterType((*MsgUpdateChainStateResponse)(nil), "nova.oracle.v1.MsgUpdateChainStateResponse")
}

func init() { proto.RegisterFile("nova/oracle/v1/tx.proto", fileDescriptor_9ebac1664c7de357) }

var fileDescriptor_9ebac1664c7de357 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xc7, 0xeb, 0x5f, 0xf7, 0x5b, 0x8b, 0x3b, 0x06, 0x4a, 0x91, 0x16, 0x8a, 0x94, 0x44, 0xe6,
	0x12, 0x01, 0xb3, 0xd5, 0x71, 0xdb, 0x05, 0x29, 0xe5, 0x30, 0x84, 0x76, 0x09, 0xe2, 0xc2, 0xa5,
	0x72, 0x1c, 0x2b, 0x89, 0x9a, 0xc6, 0x56, 0xec, 0x55, 0xdb, 0x8d, 0x97, 0xc0, 0xeb, 0xe0, 0x95,
	0xec, 0xb8, 0x23, 0xa7, 0x80, 0xda, 0x77, 0xd0, 0x57, 0x80, 0xec, 0x34, 0xa8, 0xc0, 0x0e, 0x9c,
	0xf2, 0x38, 0xdf, 0xcf, 0xf3, 0x47, 0xcf, 0xf3, 0x85, 0x27, 0x95, 0x58, 0x51, 0x22, 0x6a, 0xca,
	0x4a, 0x4e, 0x56, 0x53, 0xa2, 0xaf, 0xb1, 0xac, 0x85, 0x16, 0xce, 0xb1, 0x11, 0x70, 0x2b, 0xe0,
	0xd5, 0x74, 0xe2, 0x31, 0xa1, 0x96, 0x42, 0x91, 0x84, 0x2a, 0x03, 0x26, 0x5c, 0xd3, 0x29, 0x61,
	0xa2, 0xa8, 0x5a, 0x7e, 0xf2, 0x24, 0x13, 0x99, 0xb0, 0x21, 0x31, 0x51, 0xfb, 0x17, 0x7d, 0xee,
	0xc3, 0xf1, 0xa5, 0xca, 0x3e, 0xca, 0x94, 0x6a, 0x3e, 0xcb, 0x69, 0x51, 0x7d, 0xd0, 0x54, 0x73,
	0xa7, 0x82, 0x07, 0x26, 0xd7, 0x05, 0x01, 0x08, 0x47, 0x67, 0x4f, 0x71, 0x5b, 0x1c, 0x9b, 0xe2,
	0x78, 0x57, 0x1c, 0xcf, 0x44, 0x51, 0x45, 0x6f, 0x6e, 0x1b, 0xbf, 0xb7, 0x6d, 0xfc, 0xd1, 0x0d,
	0x5d, 0x96, 0xe7, 0xc8, 0x24, 0xa1, 0xaf, 0xdf, 0xfd, 0x30, 0x2b, 0x74, 0x7e, 0x95, 0x60, 0x26,
	0x96, 0x64, 0x37, 0x58, 0xfb, 0x39, 0x55, 0xe9, 0x82, 0xe8, 0x1b, 0xc9, 0x95, 0xcd, 0x57, 0xb1,
	0xed, 0xe3, 0xbc, 0x82, 0x83, 0x94, 0xb3, 0x62, 0x49, 0x4b, 0xf7, 0xbf, 0x00, 0x84, 0x0f, 0x23,
	0x67, 0xdb, 0xf8, 0xc7, 0x6d, 0xcd, 0x9d, 0x80, 0xe2, 0x0e, 0x71, 0x08, 0x1c, 0x0a, 0xc9, 0x6b,
	0xaa, 0x45, 0xed, 0xf6, 0x03, 0x10, 0x3e, 0x88, 0xc6, 0xdb, 0xc6, 0x7f, 0xd4, 0xe2, 0x9d, 0x82,
	0xe2, 0x5f, 0x90, 0x73, 0x0e, 0x8f, 0x92, 0x52, 0xb0, 0xc5, 0x3c, 0xe7, 0x45, 0x96, 0x6b, 0xf7,
	0x20, 0x00, 0x61, 0x3f, 0x3a, 0xd9, 0x36, 0xfe, 0xb8, 0x4d, 0xda, 0x57, 0x51, 0x3c, 0xb2, 0xcf,
	0x0b, 0xfb, 0x72, 0x30, 0x1c, 0x52, 0x29, 0xe7, 0x39, 0x55, 0xb9, 0xfb, 0x7f, 0x00, 0xc2, 0xa3,
	0xfd, 0x66, 0x9d, 0x82, 0xe2, 0x01, 0x95, 0xf2, 0x82, 0xaa, 0xdc, 0xf0, 0xcc, 0x2c, 0x72, 0x5e,
	0xa4, 0xee, 0xe1, 0x9f, 0xc3, 0x75, 0x0a, 0x8a, 0x07, 0x36, 0x7c, 0x97, 0xa2, 0xf7, 0xf0, 0xd9,
	0x3d, 0x17, 0x88, 0xb9, 0x92, 0xa2, 0x52, 0xdc, 0x6c, 0x46, 0x5d, 0x31, 0xc6, 0x95, 0xb2, 0xc7,
	0x18, 0xee, 0x6f, 0x66, 0x27, 0xa0, 0xb8, 0x43, 0xce, 0x16, 0xb0, 0x7f, 0xa9, 0x32, 0x27, 0x85,
	0x8f, 0xff, 0x3a, 0xe9, 0x73, 0xfc, 0xbb, 0x63, 0xf0, 0x3d, 0x5d, 0x27, 0x2f, 0xff, 0x01, 0xea,
	0x46, 0x8b, 0xde, 0xde, 0xae, 0x3d, 0x70, 0xb7, 0xf6, 0xc0, 0x8f, 0xb5, 0x07, 0xbe, 0x6c, 0xbc,
	0xde, 0xdd, 0xc6, 0xeb, 0x7d, 0xdb, 0x78, 0xbd, 0x4f, 0x2f, 0xf6, 0xce, 0x3f, 0xa3, 0x75, 0x51,
	0xd1, 0xd3, 0x92, 0x26, 0x8a, 0x58, 0x33, 0x5f, 0x77, 0x76, 0xb6, 0x36, 0x48, 0x0e, 0xad, 0x13,
	0x5f, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x51, 0x42, 0x82, 0xe0, 0xea, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	UpdateChainState(ctx context.Context, in *MsgUpdateChainState, opts ...grpc.CallOption) (*MsgUpdateChainStateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateChainState(ctx context.Context, in *MsgUpdateChainState, opts ...grpc.CallOption) (*MsgUpdateChainStateResponse, error) {
	out := new(MsgUpdateChainStateResponse)
	err := c.cc.Invoke(ctx, "/nova.oracle.v1.Msg/UpdateChainState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdateChainState(context.Context, *MsgUpdateChainState) (*MsgUpdateChainStateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateChainState(ctx context.Context, req *MsgUpdateChainState) (*MsgUpdateChainStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChainState not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateChainState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateChainState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateChainState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.oracle.v1.Msg/UpdateChainState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateChainState(ctx, req.(*MsgUpdateChainState))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nova.oracle.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateChainState",
			Handler:    _Msg_UpdateChainState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nova/oracle/v1/tx.proto",
}

func (m *MsgUpdateChainState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateChainState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateChainState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Decimal != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgUpdateChainStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateChainStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateChainStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateChainState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coin.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.Decimal != 0 {
		n += 1 + sovTx(uint64(m.Decimal))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTx(uint64(m.BlockHeight))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateChainStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateChainState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateChainState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateChainState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateChainStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateChainStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateChainStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

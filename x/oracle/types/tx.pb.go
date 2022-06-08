// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/oracle/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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
	Operator      string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	ChainDenom    string `protobuf:"bytes,2,opt,name=chain_denom,json=chainDenom,proto3" json:"chain_denom,omitempty"`
	StakedBalance uint64 `protobuf:"varint,3,opt,name=staked_balance,json=stakedBalance,proto3" json:"staked_balance,omitempty"`
	Decimal       uint64 `protobuf:"varint,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
	BlockHeight   uint64 `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
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

func (m *MsgUpdateChainState) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *MsgUpdateChainState) GetChainDenom() string {
	if m != nil {
		return m.ChainDenom
	}
	return ""
}

func (m *MsgUpdateChainState) GetStakedBalance() uint64 {
	if m != nil {
		return m.StakedBalance
	}
	return 0
}

func (m *MsgUpdateChainState) GetDecimal() uint64 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *MsgUpdateChainState) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type MsgUpdateChainStateResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
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
	proto.RegisterType((*MsgUpdateChainState)(nil), "novachain.oracle.v1.MsgUpdateChainState")
	proto.RegisterType((*MsgUpdateChainStateResponse)(nil), "novachain.oracle.v1.MsgUpdateChainStateResponse")
}

func init() { proto.RegisterFile("nova/oracle/v1/tx.proto", fileDescriptor_9ebac1664c7de357) }

var fileDescriptor_9ebac1664c7de357 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0x02, 0x31,
	0x18, 0xc7, 0xa9, 0xa0, 0x62, 0x51, 0x63, 0xca, 0xe0, 0x05, 0x93, 0x8a, 0x24, 0x26, 0xc4, 0xc4,
	0x9e, 0xe8, 0xe0, 0x0e, 0x0c, 0x2e, 0x2c, 0x67, 0x5c, 0x5c, 0xc8, 0x77, 0xbd, 0x2f, 0x77, 0x17,
	0x8e, 0xf6, 0x72, 0x2d, 0x04, 0xdf, 0xc2, 0xa7, 0xf1, 0x19, 0x1c, 0x19, 0x1d, 0x0d, 0xbc, 0x88,
	0xa1, 0xe4, 0x5c, 0x64, 0x70, 0xfc, 0x7e, 0xbf, 0xe6, 0xeb, 0xbf, 0xfd, 0xd3, 0x73, 0xa5, 0xe7,
	0xe0, 0xeb, 0x02, 0x64, 0x86, 0xfe, 0xbc, 0xe7, 0xdb, 0x85, 0xc8, 0x0b, 0x6d, 0x35, 0x6b, 0x6e,
	0x84, 0x4c, 0x20, 0x55, 0x62, 0x6b, 0xc5, 0xbc, 0xd7, 0xf9, 0x20, 0xb4, 0x39, 0x32, 0xf1, 0x4b,
	0x1e, 0x81, 0xc5, 0xc1, 0x46, 0x3e, 0x5b, 0xb0, 0xc8, 0x5a, 0xb4, 0xae, 0x73, 0x2c, 0xc0, 0xea,
	0xc2, 0x23, 0x6d, 0xd2, 0x3d, 0x0a, 0x7e, 0x67, 0x76, 0x49, 0x1b, 0x6e, 0xcd, 0x38, 0x42, 0xa5,
	0xa7, 0xde, 0x9e, 0xd3, 0xd4, 0xa1, 0xe1, 0x86, 0xb0, 0x6b, 0x7a, 0x6a, 0x2c, 0x4c, 0x30, 0x1a,
	0x87, 0x90, 0x81, 0x92, 0xe8, 0x55, 0xdb, 0xa4, 0x5b, 0x0b, 0x4e, 0xb6, 0xb4, 0xbf, 0x85, 0xcc,
	0xa3, 0x87, 0x11, 0xca, 0x74, 0x0a, 0x99, 0x57, 0x73, 0xbe, 0x1c, 0xd9, 0x15, 0x3d, 0x0e, 0x33,
	0x2d, 0x27, 0xe3, 0x04, 0xd3, 0x38, 0xb1, 0xde, 0xbe, 0xd3, 0x0d, 0xc7, 0x9e, 0x1c, 0xea, 0x3c,
	0xd2, 0x8b, 0x1d, 0xb9, 0x03, 0x34, 0xb9, 0x56, 0xc6, 0xed, 0x36, 0x33, 0x29, 0xd1, 0x18, 0x17,
	0xbf, 0x1e, 0x94, 0xe3, 0xfd, 0x8c, 0x56, 0x47, 0x26, 0x66, 0x8a, 0x9e, 0xfd, 0x79, 0x74, 0x57,
	0xec, 0xf8, 0x22, 0xb1, 0xe3, 0x9a, 0xd6, 0xdd, 0x7f, 0x4f, 0x96, 0x81, 0xfa, 0xc3, 0xcf, 0x15,
	0x27, 0xcb, 0x15, 0x27, 0xdf, 0x2b, 0x4e, 0xde, 0xd7, 0xbc, 0xb2, 0x5c, 0xf3, 0xca, 0xd7, 0x9a,
	0x57, 0x5e, 0x6f, 0xe2, 0xd4, 0x26, 0xb3, 0x50, 0x48, 0x3d, 0xf5, 0x07, 0x50, 0xa4, 0x0a, 0x6e,
	0x33, 0x08, 0x8d, 0xef, 0x7a, 0x5c, 0x94, 0x4d, 0xda, 0xb7, 0x1c, 0x4d, 0x78, 0xe0, 0xaa, 0x7c,
	0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x80, 0x4d, 0x27, 0xe5, 0x01, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/novachain.oracle.v1.Msg/UpdateChainState", in, out, opts...)
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
		FullMethod: "/novachain.oracle.v1.Msg/UpdateChainState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateChainState(ctx, req.(*MsgUpdateChainState))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "novachain.oracle.v1.Msg",
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
	if m.BlockHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.Decimal != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x20
	}
	if m.StakedBalance != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StakedBalance))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainDenom) > 0 {
		i -= len(m.ChainDenom)
		copy(dAtA[i:], m.ChainDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0xa
	}
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
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StakedBalance != 0 {
		n += 1 + sovTx(uint64(m.StakedBalance))
	}
	if m.Decimal != 0 {
		n += 1 + sovTx(uint64(m.Decimal))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTx(uint64(m.BlockHeight))
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainDenom", wireType)
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
			m.ChainDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedBalance", wireType)
			}
			m.StakedBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakedBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
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
				m.Decimal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
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
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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

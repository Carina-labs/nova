// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type GenesisState struct {
	// params defines all the parameters of module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// states is an amount of coins on the host chain.
	States []ChainInfo `protobuf:"bytes,2,rep,name=states,proto3" json:"states"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d9ea6305f2a803f, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetStates() []ChainInfo {
	if m != nil {
		return m.States
	}
	return nil
}

type ChainInfo struct {
	// chain_denom defines the denom of host chain's native currency.
	ChainDenom string `protobuf:"bytes,1,opt,name=chain_denom,json=chainDenom,proto3" json:"chain_denom,omitempty"`
	// operator_address is an oracle operator's address
	OperatorAddress string `protobuf:"bytes,2,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// last_block_height is the block height observed by the operator on the host chain.
	LastBlockHeight uint64 `protobuf:"varint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	// total_staked_balance is total staked asset on the host chain.
	TotalStakedBalance uint64 `protobuf:"varint,5,opt,name=total_staked_balance,json=totalStakedBalance,proto3" json:"total_staked_balance,omitempty"`
	// decimal of the native currency in host chain.
	Decimal uint64 `protobuf:"varint,6,opt,name=decimal,proto3" json:"decimal,omitempty"`
}

func (m *ChainInfo) Reset()         { *m = ChainInfo{} }
func (m *ChainInfo) String() string { return proto.CompactTextString(m) }
func (*ChainInfo) ProtoMessage()    {}
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d9ea6305f2a803f, []int{1}
}
func (m *ChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainInfo.Merge(m, src)
}
func (m *ChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChainInfo proto.InternalMessageInfo

func (m *ChainInfo) GetChainDenom() string {
	if m != nil {
		return m.ChainDenom
	}
	return ""
}

func (m *ChainInfo) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *ChainInfo) GetLastBlockHeight() uint64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *ChainInfo) GetTotalStakedBalance() uint64 {
	if m != nil {
		return m.TotalStakedBalance
	}
	return 0
}

func (m *ChainInfo) GetDecimal() uint64 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nova.oracle.v1.GenesisState")
	proto.RegisterType((*ChainInfo)(nil), "nova.oracle.v1.ChainInfo")
}

func init() { proto.RegisterFile("nova/oracle/v1/genesis.proto", fileDescriptor_4d9ea6305f2a803f) }

var fileDescriptor_4d9ea6305f2a803f = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x5b, 0xc1, 0x1a, 0x16, 0x23, 0xba, 0x21, 0xa6, 0xa2, 0x29, 0x84, 0x13, 0x92, 0xd8,
	0x0a, 0x9a, 0x78, 0xb6, 0x90, 0xa8, 0x37, 0x03, 0x37, 0x2f, 0xcd, 0xb4, 0x5d, 0xdb, 0x86, 0xb6,
	0xdb, 0x74, 0x57, 0xa2, 0x07, 0xdf, 0xc1, 0xc7, 0xe2, 0x26, 0x47, 0x4f, 0xc6, 0xc0, 0x8b, 0x98,
	0xdd, 0x52, 0x13, 0xb9, 0xed, 0xfc, 0xdf, 0xf7, 0x77, 0x9a, 0x0c, 0x3a, 0x4b, 0xe9, 0x1c, 0x2c,
	0x9a, 0x83, 0x17, 0x13, 0x6b, 0x3e, 0xb0, 0x02, 0x92, 0x12, 0x16, 0x31, 0x33, 0xcb, 0x29, 0xa7,
	0xf8, 0x40, 0x50, 0xb3, 0xa0, 0xe6, 0x7c, 0xd0, 0x3a, 0xdd, 0xb2, 0x33, 0xc8, 0x21, 0xd9, 0xc8,
	0xad, 0x66, 0x40, 0x03, 0x2a, 0x9f, 0x96, 0x78, 0x15, 0x69, 0xf7, 0x1d, 0xed, 0xdf, 0x15, 0xdf,
	0x9c, 0x72, 0xe0, 0x04, 0x5f, 0x23, 0xad, 0x68, 0xe9, 0x6a, 0x47, 0xed, 0xd5, 0x87, 0xc7, 0xe6,
	0xff, 0x1d, 0xe6, 0xa3, 0xa4, 0x76, 0x75, 0xf1, 0xdd, 0x56, 0x26, 0x1b, 0x17, 0xdf, 0x20, 0x8d,
	0x89, 0x3a, 0xd3, 0x77, 0x3a, 0x95, 0x5e, 0x7d, 0x78, 0xb2, 0xdd, 0x1a, 0x85, 0x10, 0xa5, 0x0f,
	0xe9, 0x33, 0x2d, 0x8b, 0x85, 0xde, 0xfd, 0x54, 0x51, 0xed, 0x8f, 0xe1, 0x36, 0xaa, 0x7b, 0x62,
	0x70, 0x7c, 0x92, 0xd2, 0x44, 0xfe, 0x41, 0x6d, 0x82, 0x64, 0x34, 0x16, 0x09, 0x3e, 0x47, 0x87,
	0x34, 0x23, 0x39, 0x70, 0x9a, 0x3b, 0xe0, 0xfb, 0x39, 0x61, 0x62, 0xa3, 0xb0, 0x1a, 0x65, 0x7e,
	0x5b, 0xc4, 0xb8, 0x8f, 0x8e, 0x62, 0x60, 0xdc, 0x71, 0x63, 0xea, 0xcd, 0x9c, 0x90, 0x44, 0x41,
	0xc8, 0xf5, 0x4a, 0x47, 0xed, 0x55, 0x27, 0x0d, 0x01, 0x6c, 0x91, 0xdf, 0xcb, 0x18, 0x5f, 0xa2,
	0x26, 0xa7, 0x1c, 0x62, 0x87, 0x71, 0x98, 0x11, 0xdf, 0x71, 0x21, 0x86, 0xd4, 0x23, 0xfa, 0xae,
	0xd4, 0xb1, 0x64, 0x53, 0x89, 0xec, 0x82, 0x60, 0x1d, 0xed, 0xf9, 0xc4, 0x8b, 0x12, 0x88, 0x75,
	0x4d, 0x4a, 0xe5, 0x68, 0x8f, 0x17, 0x2b, 0x43, 0x5d, 0xae, 0x0c, 0xf5, 0x67, 0x65, 0xa8, 0x1f,
	0x6b, 0x43, 0x59, 0xae, 0x0d, 0xe5, 0x6b, 0x6d, 0x28, 0x4f, 0xfd, 0x20, 0xe2, 0xe1, 0x8b, 0x6b,
	0x7a, 0x34, 0xb1, 0x46, 0x90, 0x47, 0x29, 0x5c, 0xc4, 0xe0, 0x32, 0x4b, 0x1e, 0xed, 0xb5, 0x3c,
	0x1b, 0x7f, 0xcb, 0x08, 0x73, 0x35, 0x79, 0x9d, 0xab, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59,
	0xd7, 0x87, 0x25, 0x00, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.States) > 0 {
		for iNdEx := len(m.States) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.States[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimal != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x30
	}
	if m.TotalStakedBalance != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TotalStakedBalance))
		i--
		dAtA[i] = 0x28
	}
	if m.LastBlockHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastBlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainDenom) > 0 {
		i -= len(m.ChainDenom)
		copy(dAtA[i:], m.ChainDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.States) > 0 {
		for _, e := range m.States {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastBlockHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastBlockHeight))
	}
	if m.TotalStakedBalance != 0 {
		n += 1 + sovGenesis(uint64(m.TotalStakedBalance))
	}
	if m.Decimal != 0 {
		n += 1 + sovGenesis(uint64(m.Decimal))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field States", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.States = append(m.States, ChainInfo{})
			if err := m.States[len(m.States)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeight", wireType)
			}
			m.LastBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStakedBalance", wireType)
			}
			m.TotalStakedBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalStakedBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

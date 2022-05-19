// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: novachain/gal/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the gal module's genesis state.
type GenesisState struct {
	Params          Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DepositAccounts []*DepositAccount `protobuf:"bytes,2,rep,name=depositAccounts,proto3" json:"depositAccounts,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3cd3a9ee3c69e5e, []int{0}
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

func (m *GenesisState) GetDepositAccounts() []*DepositAccount {
	if m != nil {
		return m.DepositAccounts
	}
	return nil
}

// DepositAccount defines snToken's total share and deposit information.
type DepositAccount struct {
	Denom           string         `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	DepositInfos    []*DepositInfo `protobuf:"bytes,2,rep,name=depositInfos,proto3" json:"depositInfos,omitempty"`
	TotalShare      int64          `protobuf:"varint,3,opt,name=totalShare,proto3" json:"totalShare,omitempty"`
	LastBlockUpdate int64          `protobuf:"varint,4,opt,name=lastBlockUpdate,proto3" json:"lastBlockUpdate,omitempty"`
}

func (m *DepositAccount) Reset()         { *m = DepositAccount{} }
func (m *DepositAccount) String() string { return proto.CompactTextString(m) }
func (*DepositAccount) ProtoMessage()    {}
func (*DepositAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3cd3a9ee3c69e5e, []int{1}
}
func (m *DepositAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositAccount.Merge(m, src)
}
func (m *DepositAccount) XXX_Size() int {
	return m.Size()
}
func (m *DepositAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositAccount.DiscardUnknown(m)
}

var xxx_messageInfo_DepositAccount proto.InternalMessageInfo

func (m *DepositAccount) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DepositAccount) GetDepositInfos() []*DepositInfo {
	if m != nil {
		return m.DepositInfos
	}
	return nil
}

func (m *DepositAccount) GetTotalShare() int64 {
	if m != nil {
		return m.TotalShare
	}
	return 0
}

func (m *DepositAccount) GetLastBlockUpdate() int64 {
	if m != nil {
		return m.LastBlockUpdate
	}
	return 0
}

// DepositInfo defines user address, share and debt.
type DepositInfo struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Share   int64  `protobuf:"varint,2,opt,name=share,proto3" json:"share,omitempty"`
	Debt    int64  `protobuf:"varint,3,opt,name=debt,proto3" json:"debt,omitempty"`
}

func (m *DepositInfo) Reset()         { *m = DepositInfo{} }
func (m *DepositInfo) String() string { return proto.CompactTextString(m) }
func (*DepositInfo) ProtoMessage()    {}
func (*DepositInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3cd3a9ee3c69e5e, []int{2}
}
func (m *DepositInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositInfo.Merge(m, src)
}
func (m *DepositInfo) XXX_Size() int {
	return m.Size()
}
func (m *DepositInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DepositInfo proto.InternalMessageInfo

func (m *DepositInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DepositInfo) GetShare() int64 {
	if m != nil {
		return m.Share
	}
	return 0
}

func (m *DepositInfo) GetDebt() int64 {
	if m != nil {
		return m.Debt
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "novachain.gal.v1.GenesisState")
	proto.RegisterType((*DepositAccount)(nil), "novachain.gal.v1.DepositAccount")
	proto.RegisterType((*DepositInfo)(nil), "novachain.gal.v1.DepositInfo")
}

func init() { proto.RegisterFile("novachain/gal/v1/genesis.proto", fileDescriptor_c3cd3a9ee3c69e5e) }

var fileDescriptor_c3cd3a9ee3c69e5e = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0xf4, 0x36, 0xa1, 0x88, 0x4d, 0x45, 0xd1, 0xaa, 0x07, 0x53, 0xa9, 0x8b, 0x95, 0x93, 0x0f,
	0xe0, 0x95, 0x8b, 0xc4, 0xbd, 0x01, 0x81, 0xe0, 0x04, 0xae, 0xb8, 0x70, 0x41, 0xcf, 0xf6, 0x76,
	0x63, 0x61, 0xef, 0xb3, 0xbc, 0x1b, 0x8b, 0xfe, 0x05, 0xe2, 0x57, 0xf8, 0x89, 0x1e, 0x7b, 0xe4,
	0x84, 0x50, 0xf2, 0x23, 0x28, 0xbb, 0x0e, 0x34, 0x89, 0xb8, 0xbd, 0x79, 0x33, 0xb3, 0x33, 0x96,
	0x1f, 0xe5, 0x1a, 0x7b, 0x28, 0xe6, 0x50, 0x69, 0xa1, 0xa0, 0x16, 0x7d, 0x2a, 0x94, 0xd4, 0xd2,
	0x54, 0x26, 0x69, 0x3b, 0xb4, 0xc8, 0x1e, 0xfd, 0xe5, 0x13, 0x05, 0x75, 0xd2, 0xa7, 0xa7, 0x67,
	0x7b, 0x8e, 0x16, 0x3a, 0x68, 0x06, 0xc3, 0xe9, 0x89, 0x42, 0x85, 0x6e, 0x14, 0xeb, 0x69, 0xd8,
	0x3e, 0x56, 0x88, 0xaa, 0x96, 0xc2, 0xa1, 0x7c, 0x71, 0x25, 0x40, 0x5f, 0x6f, 0xa8, 0x02, 0x4d,
	0x83, 0xe6, 0xb3, 0xf7, 0x78, 0x30, 0x50, 0xdc, 0x23, 0x91, 0x83, 0x91, 0xa2, 0x4f, 0x73, 0x69,
	0x21, 0x15, 0x05, 0x56, 0xda, 0xf3, 0xd3, 0xef, 0x84, 0x1e, 0xbd, 0xf1, 0x75, 0x2f, 0x2d, 0x58,
	0xc9, 0x5e, 0xd0, 0x43, 0x5f, 0x26, 0x24, 0x11, 0x89, 0x27, 0xe7, 0x61, 0xb2, 0x5b, 0x3f, 0x79,
	0xef, 0xf8, 0xd9, 0xf8, 0xe6, 0xd7, 0x93, 0x20, 0x1b, 0xd4, 0xec, 0x1d, 0x3d, 0x2e, 0x65, 0x8b,
	0xa6, 0xb2, 0x17, 0x45, 0x81, 0x0b, 0x6d, 0x4d, 0x78, 0x10, 0x8d, 0xe2, 0xc9, 0x79, 0xb4, 0xff,
	0xc0, 0xab, 0x2d, 0x61, 0xb6, 0x6b, 0x9c, 0xfe, 0x20, 0xf4, 0xe1, 0xb6, 0x86, 0x9d, 0xd0, 0x7b,
	0xa5, 0xd4, 0xd8, 0xb8, 0x56, 0x0f, 0x32, 0x0f, 0xd8, 0x05, 0x3d, 0x1a, 0xbc, 0x6f, 0xf5, 0x15,
	0x6e, 0x12, 0xcf, 0xfe, 0x9b, 0xb8, 0x56, 0x65, 0x5b, 0x16, 0xc6, 0x29, 0xb5, 0x68, 0xa1, 0xbe,
	0x9c, 0x43, 0x27, 0xc3, 0x51, 0x44, 0xe2, 0x51, 0x76, 0x67, 0xc3, 0x62, 0x7a, 0x5c, 0x83, 0xb1,
	0xb3, 0x1a, 0x8b, 0x2f, 0x1f, 0xdb, 0x12, 0xac, 0x0c, 0xc7, 0x4e, 0xb4, 0xbb, 0x9e, 0x7e, 0xa0,
	0x93, 0x3b, 0x31, 0x2c, 0xa4, 0xf7, 0xa1, 0x2c, 0x3b, 0x69, 0xcc, 0xd0, 0x79, 0x03, 0xd7, 0xdf,
	0x62, 0x5c, 0xda, 0x81, 0x7b, 0xc8, 0x03, 0xc6, 0xe8, 0xb8, 0x94, 0xb9, 0x1d, 0x2a, 0xb8, 0x79,
	0xf6, 0xfa, 0x66, 0xc9, 0xc9, 0xed, 0x92, 0x93, 0xdf, 0x4b, 0x4e, 0xbe, 0xad, 0x78, 0x70, 0xbb,
	0xe2, 0xc1, 0xcf, 0x15, 0x0f, 0x3e, 0x3d, 0x55, 0x95, 0x9d, 0x2f, 0xf2, 0xa4, 0xc0, 0x46, 0xbc,
	0x84, 0xae, 0xd2, 0xf0, 0xac, 0x86, 0xdc, 0x88, 0x7f, 0x97, 0xf5, 0xd5, 0xdd, 0x96, 0xbd, 0x6e,
	0xa5, 0xc9, 0x0f, 0xdd, 0xcf, 0x7e, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x10, 0x90, 0x10,
	0xab, 0x02, 0x00, 0x00,
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
	if len(m.DepositAccounts) > 0 {
		for iNdEx := len(m.DepositAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *DepositAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastBlockUpdate != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastBlockUpdate))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalShare != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TotalShare))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DepositInfos) > 0 {
		for iNdEx := len(m.DepositInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DepositInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Debt != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Debt))
		i--
		dAtA[i] = 0x18
	}
	if m.Share != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Share))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
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
	if len(m.DepositAccounts) > 0 {
		for _, e := range m.DepositAccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *DepositAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.DepositInfos) > 0 {
		for _, e := range m.DepositInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TotalShare != 0 {
		n += 1 + sovGenesis(uint64(m.TotalShare))
	}
	if m.LastBlockUpdate != 0 {
		n += 1 + sovGenesis(uint64(m.LastBlockUpdate))
	}
	return n
}

func (m *DepositInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Share != 0 {
		n += 1 + sovGenesis(uint64(m.Share))
	}
	if m.Debt != 0 {
		n += 1 + sovGenesis(uint64(m.Debt))
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
				return fmt.Errorf("proto: wrong wireType = %d for field DepositAccounts", wireType)
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
			m.DepositAccounts = append(m.DepositAccounts, &DepositAccount{})
			if err := m.DepositAccounts[len(m.DepositAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *DepositAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DepositAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositInfos", wireType)
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
			m.DepositInfos = append(m.DepositInfos, &DepositInfo{})
			if err := m.DepositInfos[len(m.DepositInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShare", wireType)
			}
			m.TotalShare = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalShare |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockUpdate", wireType)
			}
			m.LastBlockUpdate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBlockUpdate |= int64(b&0x7F) << shift
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
func (m *DepositInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DepositInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Share", wireType)
			}
			m.Share = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Share |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debt", wireType)
			}
			m.Debt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Debt |= int64(b&0x7F) << shift
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

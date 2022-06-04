// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/gal/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the gal module's genesis state.
type GenesisState struct {
	Params          Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DepositAccounts []*DepositAccount `protobuf:"bytes,2,rep,name=depositAccounts,proto3" json:"depositAccounts,omitempty"`
	WithdrawInfo    []*WithdrawInfo   `protobuf:"bytes,3,rep,name=withdrawInfo,proto3" json:"withdrawInfo,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc8782bb5a532a, []int{0}
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

func (m *GenesisState) GetWithdrawInfo() []*WithdrawInfo {
	if m != nil {
		return m.WithdrawInfo
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
	return fileDescriptor_44fc8782bb5a532a, []int{1}
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
	return fileDescriptor_44fc8782bb5a532a, []int{2}
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

type WithdrawInfo struct {
	Address        string    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Denom          string    `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount         int64     `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	CompletionTime time.Time `protobuf:"bytes,4,opt,name=completion_time,json=completionTime,proto3,stdtime" json:"completion_time"`
}

func (m *WithdrawInfo) Reset()         { *m = WithdrawInfo{} }
func (m *WithdrawInfo) String() string { return proto.CompactTextString(m) }
func (*WithdrawInfo) ProtoMessage()    {}
func (*WithdrawInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_44fc8782bb5a532a, []int{3}
}
func (m *WithdrawInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WithdrawInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WithdrawInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WithdrawInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawInfo.Merge(m, src)
}
func (m *WithdrawInfo) XXX_Size() int {
	return m.Size()
}
func (m *WithdrawInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawInfo proto.InternalMessageInfo

func (m *WithdrawInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WithdrawInfo) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *WithdrawInfo) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WithdrawInfo) GetCompletionTime() time.Time {
	if m != nil {
		return m.CompletionTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nova.gal.v1.GenesisState")
	proto.RegisterType((*DepositAccount)(nil), "nova.gal.v1.DepositAccount")
	proto.RegisterType((*DepositInfo)(nil), "nova.gal.v1.DepositInfo")
	proto.RegisterType((*WithdrawInfo)(nil), "nova.gal.v1.WithdrawInfo")
}

func init() { proto.RegisterFile("nova/gal/v1/genesis.proto", fileDescriptor_44fc8782bb5a532a) }

var fileDescriptor_44fc8782bb5a532a = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xad, 0xdb, 0x52, 0xc0, 0xa9, 0xb6, 0x92, 0x59, 0xa1, 0xb4, 0x48, 0x69, 0xd5, 0x0b, 0xbd,
	0x10, 0x2b, 0xe5, 0x0a, 0x87, 0x2d, 0x20, 0xc4, 0x01, 0x09, 0xb2, 0x20, 0x24, 0x2e, 0xab, 0x49,
	0xe2, 0x4d, 0x23, 0x92, 0x38, 0x8a, 0xdd, 0x2e, 0xfb, 0x17, 0xfb, 0x13, 0x5c, 0xf9, 0x0d, 0xf6,
	0xb8, 0x47, 0x4e, 0x80, 0xda, 0x1f, 0x41, 0xb1, 0x1d, 0x91, 0x54, 0x68, 0x6f, 0x7e, 0xf3, 0xde,
	0xcc, 0x9b, 0x19, 0x8d, 0xf1, 0x38, 0xe7, 0x5b, 0xa0, 0x31, 0xa4, 0x74, 0xeb, 0xd1, 0x98, 0xe5,
	0x4c, 0x24, 0xc2, 0x2d, 0x4a, 0x2e, 0x39, 0xb1, 0x2a, 0xca, 0x8d, 0x21, 0x75, 0xb7, 0xde, 0xc4,
	0x6e, 0xea, 0x0a, 0x28, 0x21, 0x33, 0xb2, 0xc9, 0x71, 0xcc, 0x63, 0xae, 0x9e, 0xb4, 0x7a, 0x99,
	0xe8, 0x38, 0xe6, 0x3c, 0x4e, 0x19, 0x55, 0x28, 0xd8, 0x9c, 0x53, 0xc8, 0x2f, 0x0d, 0x35, 0x3d,
	0xa4, 0x64, 0x92, 0x31, 0x21, 0x21, 0x2b, 0xea, 0xdc, 0x90, 0x8b, 0x8c, 0x8b, 0x33, 0x5d, 0x54,
	0x03, 0x43, 0x39, 0x1a, 0xd1, 0x00, 0x04, 0xa3, 0x5b, 0x2f, 0x60, 0x12, 0x3c, 0x1a, 0xf2, 0x24,
	0xd7, 0xfc, 0xfc, 0x07, 0xc2, 0xc3, 0xd7, 0x7a, 0x8a, 0x53, 0x09, 0x92, 0x11, 0x0f, 0x0f, 0x74,
	0xb7, 0x36, 0x9a, 0xa1, 0x85, 0xb5, 0x7c, 0xe0, 0x36, 0xa6, 0x72, 0xdf, 0x29, 0x6a, 0xd5, 0xbf,
	0xfe, 0x35, 0xed, 0xf8, 0x46, 0x48, 0x5e, 0xe1, 0x51, 0xc4, 0x0a, 0x2e, 0x12, 0x79, 0x12, 0x86,
	0x7c, 0x93, 0x4b, 0x61, 0x77, 0x67, 0xbd, 0x85, 0xb5, 0x7c, 0xd4, 0xca, 0x7d, 0xd9, 0xd2, 0xf8,
	0x87, 0x39, 0xe4, 0x39, 0x1e, 0x5e, 0x24, 0x72, 0x1d, 0x95, 0x70, 0xf1, 0x26, 0x3f, 0xe7, 0x76,
	0x4f, 0xd5, 0x18, 0xb7, 0x6a, 0x7c, 0x6a, 0x08, 0xfc, 0x96, 0x7c, 0xfe, 0x1d, 0xe1, 0xa3, 0xb6,
	0x05, 0x39, 0xc6, 0x77, 0x22, 0x96, 0xf3, 0x4c, 0x8d, 0x72, 0xdf, 0xd7, 0x80, 0x3c, 0xc3, 0x43,
	0x63, 0x5d, 0xe5, 0xd5, 0xbd, 0xda, 0xff, 0xeb, 0x55, 0xdb, 0x34, 0xd5, 0xc4, 0xc1, 0x58, 0x72,
	0x09, 0xe9, 0xe9, 0x1a, 0x4a, 0x66, 0xf7, 0x66, 0x68, 0xd1, 0xf3, 0x1b, 0x11, 0xb2, 0xc0, 0xa3,
	0x14, 0x84, 0x5c, 0xa5, 0x3c, 0xfc, 0xf2, 0xb1, 0x88, 0x40, 0x32, 0xbb, 0xaf, 0x44, 0x87, 0xe1,
	0xf9, 0x7b, 0x6c, 0x35, 0x6c, 0x88, 0x8d, 0xef, 0x42, 0x14, 0x95, 0x4c, 0x08, 0xd3, 0x6e, 0x0d,
	0xab, 0x31, 0x84, 0x72, 0xeb, 0xaa, 0x42, 0x1a, 0x10, 0x82, 0xfb, 0x11, 0x0b, 0xa4, 0x69, 0x41,
	0xbd, 0xe7, 0xdf, 0x10, 0x1e, 0x36, 0x57, 0x74, 0x7b, 0x51, 0xbd, 0x9b, 0x6e, 0x73, 0x37, 0x0f,
	0xf1, 0x00, 0xb2, 0x6a, 0x77, 0xa6, 0xac, 0x41, 0xe4, 0x2d, 0x1e, 0x85, 0x3c, 0x2b, 0x52, 0x26,
	0x13, 0x9e, 0x9f, 0x55, 0xf7, 0xa7, 0xa6, 0xb2, 0x96, 0x13, 0x57, 0x1f, 0xa7, 0x5b, 0x1f, 0xa7,
	0xfb, 0xa1, 0x3e, 0xce, 0xd5, 0xbd, 0xea, 0x4a, 0xae, 0x7e, 0x4f, 0x91, 0x7f, 0xf4, 0x2f, 0xb9,
	0xa2, 0x57, 0x27, 0xd7, 0x3b, 0x07, 0xdd, 0xec, 0x1c, 0xf4, 0x67, 0xe7, 0xa0, 0xab, 0xbd, 0xd3,
	0xb9, 0xd9, 0x3b, 0x9d, 0x9f, 0x7b, 0xa7, 0xf3, 0xf9, 0x71, 0x9c, 0xc8, 0xf5, 0x26, 0x70, 0x43,
	0x9e, 0xd1, 0x17, 0x50, 0x26, 0x39, 0x3c, 0x49, 0x21, 0x10, 0x54, 0xfd, 0xa6, 0xaf, 0xea, 0x3f,
	0xc9, 0xcb, 0x82, 0x89, 0x60, 0xa0, 0x0c, 0x9f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xe6,
	0xb5, 0x27, 0x90, 0x03, 0x00, 0x00,
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
	if len(m.WithdrawInfo) > 0 {
		for iNdEx := len(m.WithdrawInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
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

func (m *WithdrawInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WithdrawInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WithdrawInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if m.Amount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
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
	if len(m.WithdrawInfo) > 0 {
		for _, e := range m.WithdrawInfo {
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

func (m *WithdrawInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovGenesis(uint64(m.Amount))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovGenesis(uint64(l))
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawInfo", wireType)
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
			m.WithdrawInfo = append(m.WithdrawInfo, &WithdrawInfo{})
			if err := m.WithdrawInfo[len(m.WithdrawInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *WithdrawInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WithdrawInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WithdrawInfo: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
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

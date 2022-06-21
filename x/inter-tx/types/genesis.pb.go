// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/intertx/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

type RegisteredZone struct {
	ZoneId                 string                  `protobuf:"bytes,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	IcaConnectionInfo      *IcaConnectionInfo      `protobuf:"bytes,2,opt,name=ica_connection_info,json=icaConnectionInfo,proto3" json:"ica_connection_info,omitempty"`
	TransferConnectionInfo *TransferConnectionInfo `protobuf:"bytes,3,opt,name=transfer_connection_info,json=transferConnectionInfo,proto3" json:"transfer_connection_info,omitempty"`
	IcaAccount             *IcaAccount             `protobuf:"bytes,4,opt,name=ica_account,json=icaAccount,proto3" json:"ica_account,omitempty"`
	ValidatorAddress       string                  `protobuf:"bytes,5,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	BaseDenom              string                  `protobuf:"bytes,6,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	SnDenom                string                  `protobuf:"bytes,7,opt,name=sn_denom,json=snDenom,proto3" json:"sn_denom,omitempty"`
	StDenom                string                  `protobuf:"bytes,8,opt,name=st_denom,json=stDenom,proto3" json:"st_denom,omitempty"`
}

func (m *RegisteredZone) Reset()         { *m = RegisteredZone{} }
func (m *RegisteredZone) String() string { return proto.CompactTextString(m) }
func (*RegisteredZone) ProtoMessage()    {}
func (*RegisteredZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_f37fac6afa23cfb3, []int{0}
}
func (m *RegisteredZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredZone.Merge(m, src)
}
func (m *RegisteredZone) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredZone) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredZone.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredZone proto.InternalMessageInfo

func (m *RegisteredZone) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *RegisteredZone) GetIcaConnectionInfo() *IcaConnectionInfo {
	if m != nil {
		return m.IcaConnectionInfo
	}
	return nil
}

func (m *RegisteredZone) GetTransferConnectionInfo() *TransferConnectionInfo {
	if m != nil {
		return m.TransferConnectionInfo
	}
	return nil
}

func (m *RegisteredZone) GetIcaAccount() *IcaAccount {
	if m != nil {
		return m.IcaAccount
	}
	return nil
}

func (m *RegisteredZone) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *RegisteredZone) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *RegisteredZone) GetSnDenom() string {
	if m != nil {
		return m.SnDenom
	}
	return ""
}

func (m *RegisteredZone) GetStDenom() string {
	if m != nil {
		return m.StDenom
	}
	return ""
}

type IcaAccount struct {
	OwnerAddress string     `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	HostAddress  string     `protobuf:"bytes,2,opt,name=host_address,json=hostAddress,proto3" json:"host_address,omitempty"`
	Balance      types.Coin `protobuf:"bytes,3,opt,name=balance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balance"`
}

func (m *IcaAccount) Reset()         { *m = IcaAccount{} }
func (m *IcaAccount) String() string { return proto.CompactTextString(m) }
func (*IcaAccount) ProtoMessage()    {}
func (*IcaAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_f37fac6afa23cfb3, []int{1}
}
func (m *IcaAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IcaAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IcaAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IcaAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcaAccount.Merge(m, src)
}
func (m *IcaAccount) XXX_Size() int {
	return m.Size()
}
func (m *IcaAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_IcaAccount.DiscardUnknown(m)
}

var xxx_messageInfo_IcaAccount proto.InternalMessageInfo

func (m *IcaAccount) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *IcaAccount) GetHostAddress() string {
	if m != nil {
		return m.HostAddress
	}
	return ""
}

func (m *IcaAccount) GetBalance() types.Coin {
	if m != nil {
		return m.Balance
	}
	return types.Coin{}
}

//zone name, connection id, portID(owner address)
type IcaConnectionInfo struct {
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	PortId       string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
}

func (m *IcaConnectionInfo) Reset()         { *m = IcaConnectionInfo{} }
func (m *IcaConnectionInfo) String() string { return proto.CompactTextString(m) }
func (*IcaConnectionInfo) ProtoMessage()    {}
func (*IcaConnectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f37fac6afa23cfb3, []int{2}
}
func (m *IcaConnectionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IcaConnectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IcaConnectionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IcaConnectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcaConnectionInfo.Merge(m, src)
}
func (m *IcaConnectionInfo) XXX_Size() int {
	return m.Size()
}
func (m *IcaConnectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IcaConnectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IcaConnectionInfo proto.InternalMessageInfo

func (m *IcaConnectionInfo) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *IcaConnectionInfo) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

type TransferConnectionInfo struct {
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	PortId       string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	ChannelId    string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *TransferConnectionInfo) Reset()         { *m = TransferConnectionInfo{} }
func (m *TransferConnectionInfo) String() string { return proto.CompactTextString(m) }
func (*TransferConnectionInfo) ProtoMessage()    {}
func (*TransferConnectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f37fac6afa23cfb3, []int{3}
}
func (m *TransferConnectionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferConnectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferConnectionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferConnectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferConnectionInfo.Merge(m, src)
}
func (m *TransferConnectionInfo) XXX_Size() int {
	return m.Size()
}
func (m *TransferConnectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferConnectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransferConnectionInfo proto.InternalMessageInfo

func (m *TransferConnectionInfo) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *TransferConnectionInfo) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *TransferConnectionInfo) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisteredZone)(nil), "nova.intertx.v1.RegisteredZone")
	proto.RegisterType((*IcaAccount)(nil), "nova.intertx.v1.IcaAccount")
	proto.RegisterType((*IcaConnectionInfo)(nil), "nova.intertx.v1.IcaConnectionInfo")
	proto.RegisterType((*TransferConnectionInfo)(nil), "nova.intertx.v1.TransferConnectionInfo")
}

func init() { proto.RegisterFile("nova/intertx/v1/genesis.proto", fileDescriptor_f37fac6afa23cfb3) }

var fileDescriptor_f37fac6afa23cfb3 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x92, 0xd2, 0x4e,
	0x10, 0xc7, 0xc9, 0xf2, 0xfb, 0xc1, 0xee, 0x80, 0x7f, 0x88, 0xd6, 0x1a, 0xd6, 0xda, 0xec, 0x8a,
	0x07, 0xa9, 0x52, 0x12, 0xd1, 0xab, 0x97, 0x5d, 0x2c, 0xab, 0x72, 0x34, 0xe5, 0x69, 0x2f, 0xa9,
	0x49, 0x32, 0xc0, 0x94, 0xd0, 0x4d, 0x65, 0x06, 0x44, 0x9f, 0xc2, 0x47, 0xf0, 0xec, 0x2b, 0xf8,
	0x02, 0x7b, 0xdc, 0xa3, 0x27, 0xb5, 0xe0, 0x45, 0xac, 0xf9, 0x03, 0x6c, 0x01, 0x37, 0x4f, 0x99,
	0xfe, 0x7e, 0x7a, 0xba, 0xbe, 0xe9, 0xee, 0x21, 0xa7, 0x80, 0x33, 0x1a, 0x72, 0x90, 0xac, 0x90,
	0xf3, 0x70, 0xd6, 0x0d, 0x07, 0x0c, 0x98, 0xe0, 0x22, 0x98, 0x14, 0x28, 0xd1, 0xbd, 0xa7, 0x70,
	0x60, 0x71, 0x30, 0xeb, 0x9e, 0x3c, 0x1c, 0xe0, 0x00, 0x35, 0x0b, 0xd5, 0xc9, 0xa4, 0x9d, 0x34,
	0x33, 0x14, 0x63, 0x14, 0x89, 0x01, 0x26, 0xb0, 0xc8, 0x37, 0x51, 0x98, 0x52, 0xc1, 0xc2, 0x59,
	0x37, 0x65, 0x92, 0x76, 0xc3, 0x0c, 0x39, 0x18, 0xde, 0xfa, 0x56, 0x26, 0x77, 0x63, 0x36, 0xe0,
	0x42, 0xb2, 0x82, 0xe5, 0x57, 0x08, 0xcc, 0x7d, 0x44, 0xaa, 0x5f, 0x10, 0x58, 0xc2, 0x73, 0xcf,
	0x39, 0x77, 0xda, 0x47, 0x71, 0x45, 0x85, 0x51, 0xee, 0xc6, 0xe4, 0x01, 0xcf, 0x68, 0x92, 0x21,
	0x00, 0xcb, 0x24, 0x47, 0x48, 0x38, 0xf4, 0xd1, 0x3b, 0x38, 0x77, 0xda, 0xb5, 0x57, 0xad, 0x60,
	0xcb, 0x6b, 0x10, 0x65, 0xb4, 0xb7, 0x4e, 0x8d, 0xa0, 0x8f, 0x71, 0x83, 0x6f, 0x4b, 0x2e, 0x25,
	0x9e, 0x2c, 0x28, 0x88, 0x3e, 0x2b, 0x76, 0x0a, 0x97, 0x75, 0xe1, 0x67, 0x3b, 0x85, 0x3f, 0xd8,
	0x0b, 0x5b, 0xd5, 0x8f, 0xe5, 0x5e, 0xdd, 0x7d, 0x43, 0x6a, 0xca, 0x36, 0xcd, 0x32, 0x9c, 0x82,
	0xf4, 0xfe, 0xd3, 0x55, 0x1f, 0xef, 0xb3, 0x7b, 0x61, 0x52, 0x62, 0xc2, 0xd7, 0x67, 0xf7, 0x39,
	0x69, 0xcc, 0xe8, 0x88, 0xe7, 0x54, 0x62, 0x91, 0xd0, 0x3c, 0x2f, 0x98, 0x10, 0xde, 0xff, 0xba,
	0x2f, 0xf7, 0xd7, 0xe0, 0xc2, 0xe8, 0xee, 0x29, 0x21, 0xaa, 0xd1, 0x49, 0xce, 0x00, 0xc7, 0x5e,
	0x45, 0x67, 0x1d, 0x29, 0xe5, 0xad, 0x12, 0xdc, 0x26, 0x39, 0x14, 0x60, 0x61, 0x55, 0xc3, 0xaa,
	0x80, 0x0d, 0x92, 0x16, 0x1d, 0x5a, 0x24, 0x35, 0x6a, 0xfd, 0x70, 0x08, 0xd9, 0x98, 0x73, 0x9f,
	0x92, 0x3b, 0xf8, 0x09, 0xd8, 0xc6, 0x8c, 0x19, 0x52, 0x5d, 0x8b, 0x2b, 0x23, 0x4f, 0x48, 0x7d,
	0x88, 0x42, 0xae, 0x73, 0x0e, 0x74, 0x4e, 0x4d, 0x69, 0xab, 0x14, 0x46, 0xaa, 0x29, 0x1d, 0x51,
	0xc8, 0x98, 0x6d, 0x74, 0x33, 0xb0, 0x9b, 0xa3, 0x0c, 0x07, 0x76, 0x57, 0x82, 0x1e, 0x72, 0xb8,
	0x7c, 0x79, 0xfd, 0xeb, 0xac, 0xf4, 0xfd, 0xf7, 0x59, 0x7b, 0xc0, 0xe5, 0x70, 0x9a, 0x06, 0x19,
	0x8e, 0xed, 0x9a, 0xd9, 0x4f, 0x47, 0xe4, 0x1f, 0x43, 0xf9, 0x79, 0xc2, 0x84, 0xbe, 0x20, 0xe2,
	0x55, 0xed, 0xd6, 0x7b, 0xd2, 0xd8, 0x59, 0x04, 0xf5, 0x0f, 0xb7, 0x87, 0xbd, 0x5a, 0xb4, 0xfa,
	0x46, 0x8c, 0x72, 0xb5, 0x87, 0x13, 0x2c, 0xa4, 0xc2, 0xc6, 0x7e, 0x45, 0x85, 0x51, 0xde, 0x9a,
	0x92, 0xe3, 0xfd, 0x2b, 0xf0, 0x6f, 0x75, 0xd5, 0xf4, 0xb2, 0x21, 0x05, 0x60, 0x23, 0xc5, 0xca,
	0x66, 0x7a, 0x56, 0x89, 0xf2, 0xcb, 0x77, 0xd7, 0x0b, 0xdf, 0xb9, 0x59, 0xf8, 0xce, 0x9f, 0x85,
	0xef, 0x7c, 0x5d, 0xfa, 0xa5, 0x9b, 0xa5, 0x5f, 0xfa, 0xb9, 0xf4, 0x4b, 0x57, 0x2f, 0x6e, 0xb5,
	0xa5, 0x47, 0x0b, 0x0e, 0xb4, 0x33, 0xa2, 0xa9, 0x08, 0xf5, 0xe3, 0x9e, 0x9b, 0xe7, 0xdd, 0x91,
	0x73, 0xd3, 0xa0, 0xb4, 0xa2, 0x5f, 0xde, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0x1c,
	0x50, 0xa8, 0xfc, 0x03, 0x00, 0x00,
}

func (m *RegisteredZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StDenom) > 0 {
		i -= len(m.StDenom)
		copy(dAtA[i:], m.StDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StDenom)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SnDenom) > 0 {
		i -= len(m.SnDenom)
		copy(dAtA[i:], m.SnDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.SnDenom)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.IcaAccount != nil {
		{
			size, err := m.IcaAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.TransferConnectionInfo != nil {
		{
			size, err := m.TransferConnectionInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.IcaConnectionInfo != nil {
		{
			size, err := m.IcaConnectionInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ZoneId) > 0 {
		i -= len(m.ZoneId)
		copy(dAtA[i:], m.ZoneId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ZoneId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IcaAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IcaAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IcaAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.HostAddress) > 0 {
		i -= len(m.HostAddress)
		copy(dAtA[i:], m.HostAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.HostAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IcaConnectionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IcaConnectionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IcaConnectionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TransferConnectionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferConnectionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferConnectionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
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
func (m *RegisteredZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ZoneId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.IcaConnectionInfo != nil {
		l = m.IcaConnectionInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TransferConnectionInfo != nil {
		l = m.TransferConnectionInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.IcaAccount != nil {
		l = m.IcaAccount.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.SnDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.StDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *IcaAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.HostAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *IcaConnectionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *TransferConnectionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisteredZone) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisteredZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneId", wireType)
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
			m.ZoneId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaConnectionInfo", wireType)
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
			if m.IcaConnectionInfo == nil {
				m.IcaConnectionInfo = &IcaConnectionInfo{}
			}
			if err := m.IcaConnectionInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferConnectionInfo", wireType)
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
			if m.TransferConnectionInfo == nil {
				m.TransferConnectionInfo = &TransferConnectionInfo{}
			}
			if err := m.TransferConnectionInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaAccount", wireType)
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
			if m.IcaAccount == nil {
				m.IcaAccount = &IcaAccount{}
			}
			if err := m.IcaAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
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
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnDenom", wireType)
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
			m.SnDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StDenom", wireType)
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
			m.StDenom = string(dAtA[iNdEx:postIndex])
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
func (m *IcaAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IcaAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IcaAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
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
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostAddress", wireType)
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
			m.HostAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
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
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *IcaConnectionInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IcaConnectionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IcaConnectionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
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
func (m *TransferConnectionInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TransferConnectionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferConnectionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
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
			m.ChannelId = string(dAtA[iNdEx:postIndex])
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

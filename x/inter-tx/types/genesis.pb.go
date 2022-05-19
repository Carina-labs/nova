// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: novachain/intertx/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

type RegisteredZone struct {
	ConnectionInfo   *IcaConnectionInfo `protobuf:"bytes,1,opt,name=connection_info,json=connectionInfo,proto3" json:"connection_info,omitempty"`
	ValidatorAddress string             `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	LocalDenom       string             `protobuf:"bytes,3,opt,name=local_denom,json=localDenom,proto3" json:"local_denom,omitempty"`
	AuthzAddress     string             `protobuf:"bytes,4,opt,name=authz_address,json=authzAddress,proto3" json:"authz_address,omitempty"`
}

func (m *RegisteredZone) Reset()         { *m = RegisteredZone{} }
func (m *RegisteredZone) String() string { return proto.CompactTextString(m) }
func (*RegisteredZone) ProtoMessage()    {}
func (*RegisteredZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_910780db9c0a0633, []int{0}
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

func (m *RegisteredZone) GetConnectionInfo() *IcaConnectionInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

func (m *RegisteredZone) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *RegisteredZone) GetLocalDenom() string {
	if m != nil {
		return m.LocalDenom
	}
	return ""
}

func (m *RegisteredZone) GetAuthzAddress() string {
	if m != nil {
		return m.AuthzAddress
	}
	return ""
}

type IcaAccount struct {
	ZoneName     string     `protobuf:"bytes,1,opt,name=zone_name,json=zoneName,proto3" json:"zone_name,omitempty"`
	OwnerAddress string     `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Balance      types.Coin `protobuf:"bytes,3,opt,name=balance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balance"`
}

func (m *IcaAccount) Reset()         { *m = IcaAccount{} }
func (m *IcaAccount) String() string { return proto.CompactTextString(m) }
func (*IcaAccount) ProtoMessage()    {}
func (*IcaAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_910780db9c0a0633, []int{1}
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

func (m *IcaAccount) GetZoneName() string {
	if m != nil {
		return m.ZoneName
	}
	return ""
}

func (m *IcaAccount) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
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
	ZoneName     string `protobuf:"bytes,1,opt,name=zone_name,json=zoneName,proto3" json:"zone_name,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	OwnerAddress string `protobuf:"bytes,3,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *IcaConnectionInfo) Reset()         { *m = IcaConnectionInfo{} }
func (m *IcaConnectionInfo) String() string { return proto.CompactTextString(m) }
func (*IcaConnectionInfo) ProtoMessage()    {}
func (*IcaConnectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_910780db9c0a0633, []int{2}
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

func (m *IcaConnectionInfo) GetZoneName() string {
	if m != nil {
		return m.ZoneName
	}
	return ""
}

func (m *IcaConnectionInfo) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *IcaConnectionInfo) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisteredZone)(nil), "novachain.intertx.v1.RegisteredZone")
	proto.RegisterType((*IcaAccount)(nil), "novachain.intertx.v1.IcaAccount")
	proto.RegisterType((*IcaConnectionInfo)(nil), "novachain.intertx.v1.IcaConnectionInfo")
}

func init() {
	proto.RegisterFile("novachain/intertx/v1/genesis.proto", fileDescriptor_910780db9c0a0633)
}

var fileDescriptor_910780db9c0a0633 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x29, 0x02, 0xb2, 0x29, 0x85, 0x5a, 0x3d, 0xa4, 0x45, 0x72, 0xaa, 0x70, 0x20, 0x12,
	0x8a, 0x4d, 0xca, 0x17, 0xb4, 0xe1, 0x12, 0x21, 0x21, 0xe4, 0x63, 0x2f, 0xd6, 0x78, 0x3d, 0x75,
	0x56, 0xd8, 0x33, 0x95, 0x77, 0x63, 0x42, 0xbe, 0x82, 0xef, 0xe0, 0xce, 0x3f, 0xf4, 0xd8, 0x0b,
	0x12, 0x27, 0x40, 0xc9, 0x8f, 0x20, 0xaf, 0x9d, 0x10, 0x08, 0x52, 0x4f, 0xf6, 0xbc, 0x79, 0x3b,
	0x7a, 0xef, 0xcd, 0x88, 0x3e, 0x71, 0x09, 0x72, 0x0a, 0x8a, 0x02, 0x45, 0x06, 0x0b, 0x33, 0x0f,
	0xca, 0x51, 0x90, 0x22, 0xa1, 0x56, 0xda, 0xbf, 0x2e, 0xd8, 0xb0, 0x7b, 0xb4, 0xe1, 0xf8, 0x0d,
	0xc7, 0x2f, 0x47, 0x27, 0x47, 0x29, 0xa7, 0x6c, 0x09, 0x41, 0xf5, 0x57, 0x73, 0x4f, 0x8e, 0x25,
	0xeb, 0x9c, 0x75, 0x54, 0x37, 0xea, 0xa2, 0x69, 0x79, 0x75, 0x15, 0xc4, 0xa0, 0x31, 0x28, 0x47,
	0x31, 0x1a, 0x18, 0x05, 0x92, 0x15, 0xd5, 0xfd, 0xfe, 0x37, 0x47, 0x1c, 0x84, 0x98, 0x2a, 0x6d,
	0xb0, 0xc0, 0xe4, 0x92, 0x09, 0xdd, 0xf7, 0xe2, 0x89, 0x64, 0x22, 0x94, 0x46, 0x31, 0x45, 0x8a,
	0xae, 0xb8, 0xeb, 0x9c, 0x3a, 0x83, 0xce, 0xd9, 0x0b, 0xff, 0x7f, 0x9a, 0xfc, 0x89, 0x84, 0xf1,
	0x86, 0x3f, 0xa1, 0x2b, 0x0e, 0x0f, 0xe4, 0x5f, 0xb5, 0xfb, 0x52, 0x1c, 0x96, 0x90, 0xa9, 0x04,
	0x0c, 0x17, 0x11, 0x24, 0x49, 0x81, 0x5a, 0x77, 0xef, 0x9d, 0x3a, 0x83, 0x76, 0xf8, 0x74, 0xd3,
	0x38, 0xaf, 0x71, 0xb7, 0x27, 0x3a, 0x19, 0x4b, 0xc8, 0xa2, 0x04, 0x89, 0xf3, 0xee, 0x9e, 0xa5,
	0x09, 0x0b, 0xbd, 0xa9, 0x10, 0xf7, 0xb9, 0x78, 0x0c, 0x33, 0x33, 0x5d, 0x6c, 0x26, 0xdd, 0xb7,
	0x94, 0x7d, 0x0b, 0x36, 0x53, 0xfa, 0x5f, 0x1d, 0x21, 0x26, 0x12, 0xce, 0xa5, 0xe4, 0x19, 0x19,
	0xf7, 0x99, 0x68, 0x2f, 0x98, 0x30, 0x22, 0xc8, 0xd1, 0xba, 0x69, 0x87, 0x8f, 0x2a, 0xe0, 0x1d,
	0xe4, 0x58, 0x0d, 0xe4, 0x8f, 0x84, 0xff, 0x4a, 0xdb, 0xb7, 0xe0, 0x5a, 0x16, 0x8a, 0x87, 0x31,
	0x64, 0x40, 0x12, 0xad, 0xa4, 0xce, 0xd9, 0xb1, 0xdf, 0x04, 0x5d, 0x45, 0xeb, 0x37, 0xd1, 0xfa,
	0x63, 0x56, 0x74, 0xf1, 0xea, 0xe6, 0x47, 0xaf, 0xf5, 0xe5, 0x67, 0x6f, 0x90, 0x2a, 0x33, 0x9d,
	0xc5, 0xbe, 0xe4, 0xbc, 0xd9, 0x4a, 0xf3, 0x19, 0xea, 0xe4, 0x43, 0x60, 0x3e, 0x5d, 0xa3, 0xb6,
	0x0f, 0x74, 0xb8, 0x9e, 0xdd, 0x5f, 0x88, 0xc3, 0x9d, 0x3c, 0xef, 0x54, 0xbf, 0xbd, 0xae, 0x64,
	0xad, 0x7e, 0x6b, 0x07, 0xc9, 0xae, 0xc5, 0xbd, 0x5d, 0x8b, 0x17, 0x6f, 0x6f, 0x96, 0x9e, 0x73,
	0xbb, 0xf4, 0x9c, 0x5f, 0x4b, 0xcf, 0xf9, 0xbc, 0xf2, 0x5a, 0xb7, 0x2b, 0xaf, 0xf5, 0x7d, 0xe5,
	0xb5, 0x2e, 0x47, 0x5b, 0x46, 0xc6, 0x50, 0x28, 0x82, 0x61, 0x06, 0xb1, 0x0e, 0xfe, 0xdc, 0xf1,
	0xbc, 0xbe, 0xe4, 0xa1, 0x99, 0xd7, 0xbe, 0xe2, 0x07, 0xf6, 0xbe, 0x5e, 0xff, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x6a, 0x33, 0xd8, 0xe0, 0xec, 0x02, 0x00, 0x00,
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
	if len(m.AuthzAddress) > 0 {
		i -= len(m.AuthzAddress)
		copy(dAtA[i:], m.AuthzAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AuthzAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LocalDenom) > 0 {
		i -= len(m.LocalDenom)
		copy(dAtA[i:], m.LocalDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.LocalDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ConnectionInfo != nil {
		{
			size, err := m.ConnectionInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
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
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ZoneName) > 0 {
		i -= len(m.ZoneName)
		copy(dAtA[i:], m.ZoneName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ZoneName)))
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
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ZoneName) > 0 {
		i -= len(m.ZoneName)
		copy(dAtA[i:], m.ZoneName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ZoneName)))
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
	if m.ConnectionInfo != nil {
		l = m.ConnectionInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.LocalDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.AuthzAddress)
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
	l = len(m.ZoneName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.OwnerAddress)
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
	l = len(m.ZoneName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.OwnerAddress)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionInfo", wireType)
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
			if m.ConnectionInfo == nil {
				m.ConnectionInfo = &IcaConnectionInfo{}
			}
			if err := m.ConnectionInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalDenom", wireType)
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
			m.LocalDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthzAddress", wireType)
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
			m.AuthzAddress = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneName", wireType)
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
			m.ZoneName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneName", wireType)
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
			m.ZoneName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
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

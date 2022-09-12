// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// ChainInfo refers to the state of the counterpart chain to be stored on the Oracle module.
// Status includes the amount of coins delegated to Zone, AppHash, and block height.
type ChainInfo struct {
	// coin refers to the sum of owned, staked and claimable quantity of the coin
	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coin" yaml:"coin"`
	// operator_address is an oracle operator's address
	OperatorAddress string `protobuf:"bytes,2,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty" yaml:"operator_address"`
	// last_block_height is the block height observed by the operator on the host chain.
	LastBlockHeight int64 `protobuf:"varint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty" yaml:"last_block_height"`
	// app_hash of the block fetched by oracle from host chain
	AppHash []byte `protobuf:"bytes,4,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty" yaml:"app_hash"`
	// chain_id of the host chain
	ChainId       string `protobuf:"bytes,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	OracleVersion uint64 `protobuf:"varint,6,opt,name=oracle_version,json=oracleVersion,proto3" json:"oracle_version,omitempty"`
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

func (m *ChainInfo) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *ChainInfo) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *ChainInfo) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *ChainInfo) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *ChainInfo) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChainInfo) GetOracleVersion() uint64 {
	if m != nil {
		return m.OracleVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nova.oracle.v1.GenesisState")
	proto.RegisterType((*ChainInfo)(nil), "nova.oracle.v1.ChainInfo")
}

func init() { proto.RegisterFile("nova/oracle/v1/genesis.proto", fileDescriptor_4d9ea6305f2a803f) }

var fileDescriptor_4d9ea6305f2a803f = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcd, 0x8a, 0xd3, 0x40,
	0x1c, 0x6f, 0x6c, 0xed, 0xba, 0xd3, 0x75, 0xab, 0x51, 0x34, 0xfb, 0x41, 0x52, 0x02, 0x42, 0x10,
	0x76, 0x42, 0x57, 0x41, 0xf0, 0x22, 0xa6, 0xa2, 0xdd, 0x9b, 0x44, 0xf0, 0xe0, 0x25, 0xfc, 0x93,
	0x8c, 0x49, 0xd8, 0x74, 0x26, 0x64, 0xc6, 0xe2, 0x1e, 0x7c, 0x07, 0x0f, 0x3e, 0x85, 0x4f, 0xb2,
	0xc7, 0x3d, 0x7a, 0x8a, 0xd2, 0xbe, 0x41, 0x9f, 0x40, 0x66, 0x26, 0x91, 0xb5, 0x7b, 0xca, 0xf0,
	0xfb, 0xfa, 0x7f, 0xe4, 0x8f, 0x8e, 0x29, 0x5b, 0x82, 0xcf, 0x6a, 0x48, 0x4a, 0xe2, 0x2f, 0xa7,
	0x7e, 0x46, 0x28, 0xe1, 0x05, 0xc7, 0x55, 0xcd, 0x04, 0x33, 0xf7, 0x25, 0x8b, 0x35, 0x8b, 0x97,
	0xd3, 0xc3, 0xa3, 0x2d, 0x75, 0x05, 0x35, 0x2c, 0x5a, 0xf1, 0xa1, 0x9d, 0x30, 0xbe, 0x60, 0xdc,
	0x8f, 0x81, 0x4b, 0x32, 0x26, 0x02, 0xa6, 0x7e, 0xc2, 0x0a, 0xda, 0xf2, 0x0f, 0x33, 0x96, 0x31,
	0xf5, 0xf4, 0xe5, 0x4b, 0xa3, 0xee, 0x37, 0xb4, 0xf7, 0x4e, 0xd7, 0xfc, 0x20, 0x40, 0x10, 0xf3,
	0x39, 0x1a, 0xea, 0x54, 0xcb, 0x98, 0x18, 0xde, 0xe8, 0xf4, 0x11, 0xfe, 0xbf, 0x07, 0xfc, 0x5e,
	0xb1, 0xc1, 0xe0, 0xb2, 0x71, 0x7a, 0x61, 0xab, 0x35, 0x5f, 0xa0, 0x21, 0x97, 0x76, 0x6e, 0xdd,
	0x9a, 0xf4, 0xbd, 0xd1, 0xe9, 0xc1, 0xb6, 0x6b, 0x96, 0x43, 0x41, 0xcf, 0xe8, 0x67, 0xd6, 0x19,
	0xb5, 0xdc, 0xfd, 0xd1, 0x47, 0xbb, 0xff, 0x38, 0x93, 0xa2, 0x81, 0x6c, 0xb8, 0x2d, 0x7d, 0x80,
	0xf5, 0x44, 0x58, 0x4e, 0x84, 0xdb, 0x89, 0xf0, 0x8c, 0x15, 0x34, 0x78, 0x25, 0x43, 0x36, 0x8d,
	0x33, 0xba, 0x80, 0x45, 0xf9, 0xd2, 0x95, 0x26, 0xf7, 0xe7, 0x6f, 0xc7, 0xcb, 0x0a, 0x91, 0x7f,
	0x89, 0x71, 0xc2, 0x16, 0x7e, 0xbb, 0x0d, 0xfd, 0x39, 0xe1, 0xe9, 0xb9, 0x2f, 0x2e, 0x2a, 0xc2,
	0x95, 0x9f, 0x87, 0xaa, 0x8e, 0xf9, 0x16, 0xdd, 0x63, 0x15, 0xa9, 0x41, 0xb0, 0x3a, 0x82, 0x34,
	0xad, 0x09, 0x97, 0x03, 0x18, 0xde, 0x6e, 0x70, 0xb4, 0x69, 0x9c, 0xc7, 0x3a, 0x7c, 0x5b, 0xe1,
	0x86, 0xe3, 0x0e, 0x7a, 0xad, 0x11, 0x73, 0x8e, 0xee, 0x97, 0xc0, 0x45, 0x14, 0x97, 0x2c, 0x39,
	0x8f, 0x72, 0x52, 0x64, 0xb9, 0xb0, 0xfa, 0x13, 0xc3, 0xeb, 0x07, 0xc7, 0x9b, 0xc6, 0xb1, 0x74,
	0xd0, 0x0d, 0x89, 0x1b, 0x8e, 0x25, 0x16, 0x48, 0x68, 0xae, 0x10, 0x13, 0xa3, 0x3b, 0x50, 0x55,
	0x51, 0x0e, 0x3c, 0xb7, 0x06, 0x13, 0xc3, 0xdb, 0x0b, 0x1e, 0x6c, 0x1a, 0x67, 0xac, 0x03, 0x3a,
	0xc6, 0x0d, 0x77, 0xa0, 0xaa, 0xe6, 0xc0, 0x73, 0xa9, 0x4f, 0xe4, 0xfa, 0xa2, 0x22, 0xb5, 0x6e,
	0xab, 0xce, 0xaf, 0xe9, 0x3b, 0xc6, 0x0d, 0x77, 0xd4, 0xf3, 0x2c, 0x35, 0x9f, 0xa0, 0x7d, 0xfd,
	0x53, 0xa2, 0x25, 0xa9, 0x79, 0xc1, 0xa8, 0x35, 0x9c, 0x18, 0xde, 0x20, 0xbc, 0xab, 0xd1, 0x8f,
	0x1a, 0x0c, 0xde, 0x5c, 0xae, 0x6c, 0xe3, 0x6a, 0x65, 0x1b, 0x7f, 0x56, 0xb6, 0xf1, 0x7d, 0x6d,
	0xf7, 0xae, 0xd6, 0x76, 0xef, 0xd7, 0xda, 0xee, 0x7d, 0x7a, 0x7a, 0x6d, 0xc5, 0x33, 0xa8, 0x0b,
	0x0a, 0x27, 0x25, 0xc4, 0xdc, 0x57, 0x97, 0xf9, 0xb5, 0xbb, 0x4d, 0xb5, 0xea, 0x78, 0xa8, 0x4e,
	0xec, 0xd9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x9a, 0x76, 0x52, 0xe5, 0x02, 0x00, 0x00,
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
	if m.OracleVersion != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.OracleVersion))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x22
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
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.Coin.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastBlockHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastBlockHeight))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.OracleVersion != 0 {
		n += 1 + sovGenesis(uint64(m.OracleVersion))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
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
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
				m.LastBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleVersion", wireType)
			}
			m.OracleVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OracleVersion |= uint64(b&0x7F) << shift
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

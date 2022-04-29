// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: novachain/gal/v1/genesis.proto

package types

import (
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the parameters for the gal module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novachain_gal_v1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_novachain_gal_v1_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_novachain_gal_v1_genesis_proto_rawDescGZIP(), []int{0}
}

// GenesisState defines the gal module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params          *Params           `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	DepositAccounts []*DepositAccount `protobuf:"bytes,2,rep,name=depositAccounts,proto3" json:"depositAccounts,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novachain_gal_v1_genesis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_novachain_gal_v1_genesis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_novachain_gal_v1_genesis_proto_rawDescGZIP(), []int{1}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetDepositAccounts() []*DepositAccount {
	if x != nil {
		return x.DepositAccounts
	}
	return nil
}

// DepositAccount defines snToken's total share and deposit information.
type DepositAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom           string         `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	DepositInfos    []*DepositInfo `protobuf:"bytes,2,rep,name=depositInfos,proto3" json:"depositInfos,omitempty"`
	TotalShare      int64          `protobuf:"varint,3,opt,name=totalShare,proto3" json:"totalShare,omitempty"`
	LastBlockUpdate int64          `protobuf:"varint,4,opt,name=lastBlockUpdate,proto3" json:"lastBlockUpdate,omitempty"`
}

func (x *DepositAccount) Reset() {
	*x = DepositAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novachain_gal_v1_genesis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositAccount) ProtoMessage() {}

func (x *DepositAccount) ProtoReflect() protoreflect.Message {
	mi := &file_novachain_gal_v1_genesis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositAccount.ProtoReflect.Descriptor instead.
func (*DepositAccount) Descriptor() ([]byte, []int) {
	return file_novachain_gal_v1_genesis_proto_rawDescGZIP(), []int{2}
}

func (x *DepositAccount) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *DepositAccount) GetDepositInfos() []*DepositInfo {
	if x != nil {
		return x.DepositInfos
	}
	return nil
}

func (x *DepositAccount) GetTotalShare() int64 {
	if x != nil {
		return x.TotalShare
	}
	return 0
}

func (x *DepositAccount) GetLastBlockUpdate() int64 {
	if x != nil {
		return x.LastBlockUpdate
	}
	return 0
}

// DepositInfo defines user address, share and debt.
type DepositInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Share   int64  `protobuf:"varint,2,opt,name=share,proto3" json:"share,omitempty"`
	Debt    int64  `protobuf:"varint,3,opt,name=debt,proto3" json:"debt,omitempty"`
}

func (x *DepositInfo) Reset() {
	*x = DepositInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novachain_gal_v1_genesis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositInfo) ProtoMessage() {}

func (x *DepositInfo) ProtoReflect() protoreflect.Message {
	mi := &file_novachain_gal_v1_genesis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositInfo.ProtoReflect.Descriptor instead.
func (*DepositInfo) Descriptor() ([]byte, []int) {
	return file_novachain_gal_v1_genesis_proto_rawDescGZIP(), []int{3}
}

func (x *DepositInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DepositInfo) GetShare() int64 {
	if x != nil {
		return x.Share
	}
	return 0
}

func (x *DepositInfo) GetDebt() int64 {
	if x != nil {
		return x.Debt
	}
	return 0
}

var File_novachain_gal_v1_genesis_proto protoreflect.FileDescriptor

var file_novachain_gal_v1_genesis_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x61, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x9c, 0x01,
	0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x67, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0f, 0x64, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xb8, 0x01, 0x0a,
	0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x46, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x6f,
	0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x51, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x62, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x65, 0x62, 0x74, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x72, 0x69, 0x6e, 0x61, 0x2d,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x78,
	0x2f, 0x67, 0x61, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_novachain_gal_v1_genesis_proto_rawDescOnce sync.Once
	file_novachain_gal_v1_genesis_proto_rawDescData = file_novachain_gal_v1_genesis_proto_rawDesc
)

func file_novachain_gal_v1_genesis_proto_rawDescGZIP() []byte {
	file_novachain_gal_v1_genesis_proto_rawDescOnce.Do(func() {
		file_novachain_gal_v1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_novachain_gal_v1_genesis_proto_rawDescData)
	})
	return file_novachain_gal_v1_genesis_proto_rawDescData
}

var file_novachain_gal_v1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_novachain_gal_v1_genesis_proto_goTypes = []interface{}{
	(*Params)(nil),         // 0: novachain.gal.v1beta1.Params
	(*GenesisState)(nil),   // 1: novachain.gal.v1beta1.GenesisState
	(*DepositAccount)(nil), // 2: novachain.gal.v1beta1.DepositAccount
	(*DepositInfo)(nil),    // 3: novachain.gal.v1beta1.DepositInfo
}
var file_novachain_gal_v1_genesis_proto_depIdxs = []int32{
	0, // 0: novachain.gal.v1beta1.GenesisState.params:type_name -> novachain.gal.v1beta1.Params
	2, // 1: novachain.gal.v1beta1.GenesisState.depositAccounts:type_name -> novachain.gal.v1beta1.DepositAccount
	3, // 2: novachain.gal.v1beta1.DepositAccount.depositInfos:type_name -> novachain.gal.v1beta1.DepositInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_novachain_gal_v1_genesis_proto_init() }
func file_novachain_gal_v1_genesis_proto_init() {
	if File_novachain_gal_v1_genesis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_novachain_gal_v1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novachain_gal_v1_genesis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novachain_gal_v1_genesis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novachain_gal_v1_genesis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_novachain_gal_v1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_novachain_gal_v1_genesis_proto_goTypes,
		DependencyIndexes: file_novachain_gal_v1_genesis_proto_depIdxs,
		MessageInfos:      file_novachain_gal_v1_genesis_proto_msgTypes,
	}.Build()
	File_novachain_gal_v1_genesis_proto = out.File
	file_novachain_gal_v1_genesis_proto_rawDesc = nil
	file_novachain_gal_v1_genesis_proto_goTypes = nil
	file_novachain_gal_v1_genesis_proto_depIdxs = nil
}

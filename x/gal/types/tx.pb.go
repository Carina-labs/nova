// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: novachain/gal/v1/tx.proto

package types

import (
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgDeposit defines user who deposit and amount of coins.
type MsgDeposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string depositor = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
	Amount []*types.Coin `protobuf:"bytes,2,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgDeposit) Reset() {
	*x = MsgDeposit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novachain_gal_v1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeposit) ProtoMessage() {}

func (x *MsgDeposit) ProtoReflect() protoreflect.Message {
	mi := &file_novachain_gal_v1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeposit.ProtoReflect.Descriptor instead.
func (*MsgDeposit) Descriptor() ([]byte, []int) {
	return file_novachain_gal_v1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgDeposit) GetAmount() []*types.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

type MsgDepositResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDepositResponse) Reset() {
	*x = MsgDepositResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novachain_gal_v1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDepositResponse) ProtoMessage() {}

func (x *MsgDepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_novachain_gal_v1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDepositResponse.ProtoReflect.Descriptor instead.
func (*MsgDepositResponse) Descriptor() ([]byte, []int) {
	return file_novachain_gal_v1_tx_proto_rawDescGZIP(), []int{1}
}

var File_novachain_gal_v1_tx_proto protoreflect.FileDescriptor

var file_novachain_gal_v1_tx_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x61, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x6f, 0x76,
	0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x5e, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x57, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x12, 0x21, 0x2e, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x67, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x67, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x61, 0x72, 0x69, 0x6e, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x6e, 0x6f, 0x76, 0x61,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x78, 0x2f, 0x67, 0x61, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_novachain_gal_v1_tx_proto_rawDescOnce sync.Once
	file_novachain_gal_v1_tx_proto_rawDescData = file_novachain_gal_v1_tx_proto_rawDesc
)

func file_novachain_gal_v1_tx_proto_rawDescGZIP() []byte {
	file_novachain_gal_v1_tx_proto_rawDescOnce.Do(func() {
		file_novachain_gal_v1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_novachain_gal_v1_tx_proto_rawDescData)
	})
	return file_novachain_gal_v1_tx_proto_rawDescData
}

var file_novachain_gal_v1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_novachain_gal_v1_tx_proto_goTypes = []interface{}{
	(*MsgDeposit)(nil),         // 0: novachain.gal.v1beta1.MsgDeposit
	(*MsgDepositResponse)(nil), // 1: novachain.gal.v1beta1.MsgDepositResponse
	(*types.Coin)(nil),         // 2: cosmos.base.v1beta1.Coin
}
var file_novachain_gal_v1_tx_proto_depIdxs = []int32{
	2, // 0: novachain.gal.v1beta1.MsgDeposit.amount:type_name -> cosmos.base.v1beta1.Coin
	0, // 1: novachain.gal.v1beta1.Msg.deposit:input_type -> novachain.gal.v1beta1.MsgDeposit
	1, // 2: novachain.gal.v1beta1.Msg.deposit:output_type -> novachain.gal.v1beta1.MsgDepositResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_novachain_gal_v1_tx_proto_init() }
func file_novachain_gal_v1_tx_proto_init() {
	if File_novachain_gal_v1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_novachain_gal_v1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeposit); i {
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
		file_novachain_gal_v1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDepositResponse); i {
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
			RawDescriptor: file_novachain_gal_v1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_novachain_gal_v1_tx_proto_goTypes,
		DependencyIndexes: file_novachain_gal_v1_tx_proto_depIdxs,
		MessageInfos:      file_novachain_gal_v1_tx_proto_msgTypes,
	}.Build()
	File_novachain_gal_v1_tx_proto = out.File
	file_novachain_gal_v1_tx_proto_rawDesc = nil
	file_novachain_gal_v1_tx_proto_goTypes = nil
	file_novachain_gal_v1_tx_proto_depIdxs = nil
}

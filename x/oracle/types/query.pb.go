// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/oracle/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4468790215c01e6, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4468790215c01e6, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryStateRequest struct {
	ChainDenom string `protobuf:"bytes,1,opt,name=chain_denom,json=chainDenom,proto3" json:"chain_denom,omitempty" yaml:"chain_denom"`
}

func (m *QueryStateRequest) Reset()         { *m = QueryStateRequest{} }
func (m *QueryStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStateRequest) ProtoMessage()    {}
func (*QueryStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4468790215c01e6, []int{2}
}
func (m *QueryStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStateRequest.Merge(m, src)
}
func (m *QueryStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStateRequest proto.InternalMessageInfo

func (m *QueryStateRequest) GetChainDenom() string {
	if m != nil {
		return m.ChainDenom
	}
	return ""
}

type QueryStateResponse struct {
	Coin            types.Coin `protobuf:"bytes,1,opt,name=coin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coin" yaml:"coin"`
	Operator        string     `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty" yaml:"operator"`
	Decimal         uint64     `protobuf:"varint,3,opt,name=decimal,proto3" json:"decimal,omitempty" yaml:"decimal"`
	LastBlockHeight uint64     `protobuf:"varint,4,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty" yaml:"last_block_height"`
	AppHash         string     `protobuf:"bytes,5,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty" yaml:"app_hash"`
	ChainId         string     `protobuf:"bytes,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	BlockProposer   string     `protobuf:"bytes,7,opt,name=block_proposer,json=blockProposer,proto3" json:"block_proposer,omitempty" yaml:"block_proopser"`
}

func (m *QueryStateResponse) Reset()         { *m = QueryStateResponse{} }
func (m *QueryStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStateResponse) ProtoMessage()    {}
func (*QueryStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4468790215c01e6, []int{3}
}
func (m *QueryStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStateResponse.Merge(m, src)
}
func (m *QueryStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStateResponse proto.InternalMessageInfo

func (m *QueryStateResponse) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *QueryStateResponse) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *QueryStateResponse) GetDecimal() uint64 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *QueryStateResponse) GetLastBlockHeight() uint64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *QueryStateResponse) GetAppHash() string {
	if m != nil {
		return m.AppHash
	}
	return ""
}

func (m *QueryStateResponse) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *QueryStateResponse) GetBlockProposer() string {
	if m != nil {
		return m.BlockProposer
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "nova.oracle.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "nova.oracle.v1.QueryParamsResponse")
	proto.RegisterType((*QueryStateRequest)(nil), "nova.oracle.v1.QueryStateRequest")
	proto.RegisterType((*QueryStateResponse)(nil), "nova.oracle.v1.QueryStateResponse")
}

func init() { proto.RegisterFile("nova/oracle/v1/query.proto", fileDescriptor_f4468790215c01e6) }

var fileDescriptor_f4468790215c01e6 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcf, 0x4f, 0xd4, 0x40,
	0x14, 0xde, 0xc2, 0xb2, 0x8b, 0x43, 0x84, 0x30, 0xfc, 0xb0, 0xac, 0xd8, 0xe2, 0x78, 0x21, 0x46,
	0xda, 0x80, 0x26, 0x26, 0x5e, 0x34, 0x0b, 0x07, 0x8c, 0x1e, 0xb0, 0xde, 0xbc, 0x6c, 0x66, 0xbb,
	0x93, 0x6d, 0x43, 0x3b, 0x6f, 0xe8, 0x0c, 0x44, 0xae, 0xfe, 0x05, 0x26, 0xfe, 0x17, 0xfe, 0x25,
	0xdc, 0x24, 0xf1, 0xe2, 0xa9, 0x1a, 0xf0, 0xe2, 0xb5, 0x7f, 0x81, 0xe9, 0xcc, 0x2c, 0x59, 0x16,
	0xc2, 0xa9, 0xd3, 0xf7, 0x7d, 0xef, 0x7d, 0x5f, 0xde, 0x7c, 0x83, 0x3a, 0x1c, 0x4e, 0x68, 0x08,
	0x05, 0x8d, 0x33, 0x16, 0x9e, 0x6c, 0x87, 0x47, 0xc7, 0xac, 0x38, 0x0d, 0x44, 0x01, 0x0a, 0xf0,
	0x7c, 0x8d, 0x05, 0x06, 0x0b, 0x4e, 0xb6, 0x3b, 0x5e, 0x0c, 0x32, 0x07, 0x19, 0xf6, 0xa9, 0xac,
	0xb9, 0x7d, 0xa6, 0xe8, 0x76, 0x18, 0x43, 0xca, 0x0d, 0xbf, 0xb3, 0x3e, 0x04, 0x18, 0x66, 0x2c,
	0xa4, 0x22, 0x0d, 0x29, 0xe7, 0xa0, 0xa8, 0x4a, 0x81, 0x4b, 0x8b, 0x2e, 0x0f, 0x61, 0x08, 0xfa,
	0x18, 0xd6, 0x27, 0x5b, 0x7d, 0x38, 0xa1, 0x2f, 0x68, 0x41, 0x73, 0xdb, 0x42, 0x96, 0x11, 0xfe,
	0x50, 0xfb, 0x39, 0xd0, 0xc5, 0x88, 0x1d, 0x1d, 0x33, 0xa9, 0xc8, 0x3b, 0xb4, 0x74, 0xad, 0x2a,
	0x05, 0x70, 0xc9, 0xf0, 0x0b, 0xd4, 0x32, 0xcd, 0xae, 0xb3, 0xe1, 0x6c, 0xce, 0xed, 0xac, 0x06,
	0xd7, 0xed, 0x07, 0x86, 0xdf, 0x6d, 0x9e, 0x95, 0x7e, 0x23, 0xb2, 0x5c, 0xf2, 0x1e, 0x2d, 0xea,
	0x61, 0x1f, 0x15, 0x55, 0xcc, 0x2a, 0xe0, 0x97, 0x68, 0x2e, 0x4e, 0x68, 0xca, 0x7b, 0x03, 0xc6,
	0x21, 0xd7, 0xf3, 0xee, 0x75, 0x57, 0xab, 0xd2, 0xc7, 0xa7, 0x34, 0xcf, 0x5e, 0x91, 0x31, 0x90,
	0x44, 0x48, 0xff, 0xed, 0xe9, 0x9f, 0x1f, 0xd3, 0xd6, 0xb1, 0x1d, 0x67, 0xad, 0x71, 0xd4, 0xac,
	0xd7, 0x64, 0x8d, 0xad, 0x05, 0x66, 0x8f, 0x41, 0xbd, 0xc7, 0xc0, 0xee, 0x31, 0xd8, 0x85, 0x94,
	0x77, 0x5f, 0xd7, 0xde, 0xaa, 0xd2, 0x9f, 0xb3, 0x3a, 0x90, 0x72, 0xf2, 0xfd, 0xb7, 0xbf, 0x39,
	0x4c, 0x55, 0x72, 0xdc, 0x0f, 0x62, 0xc8, 0x43, 0x7b, 0x07, 0xe6, 0xb3, 0x25, 0x07, 0x87, 0xa1,
	0x3a, 0x15, 0x4c, 0xea, 0x7e, 0x19, 0x69, 0x1d, 0x1c, 0xa2, 0x59, 0x10, 0xac, 0xa0, 0x0a, 0x0a,
	0x77, 0x4a, 0x9b, 0x5f, 0xaa, 0x4a, 0x7f, 0xc1, 0x0c, 0x1d, 0x21, 0x24, 0xba, 0x22, 0xe1, 0x67,
	0xa8, 0x3d, 0x60, 0x71, 0x9a, 0xd3, 0xcc, 0x9d, 0xde, 0x70, 0x36, 0x9b, 0x5d, 0x5c, 0x95, 0xfe,
	0xbc, 0xe1, 0x5b, 0x80, 0x44, 0x23, 0x0a, 0xde, 0x47, 0x8b, 0x19, 0x95, 0xaa, 0xd7, 0xcf, 0x20,
	0x3e, 0xec, 0x25, 0x2c, 0x1d, 0x26, 0xca, 0x6d, 0xea, 0xbe, 0xf5, 0xaa, 0xf4, 0x5d, 0xd3, 0x77,
	0x83, 0x42, 0xa2, 0x85, 0xba, 0xd6, 0xad, 0x4b, 0xfb, 0xba, 0x82, 0x03, 0x34, 0x4b, 0x85, 0xe8,
	0x25, 0x54, 0x26, 0xee, 0xcc, 0xa4, 0xd1, 0x11, 0x42, 0xa2, 0x36, 0x15, 0x62, 0x9f, 0xca, 0xa4,
	0xe6, 0x9b, 0xdd, 0xa7, 0x03, 0xb7, 0x35, 0xc9, 0x1f, 0x21, 0x24, 0x6a, 0xeb, 0xe3, 0xdb, 0x01,
	0x7e, 0x83, 0xe6, 0x8d, 0x03, 0x51, 0x80, 0x00, 0xc9, 0x0a, 0xb7, 0xad, 0xbb, 0xd6, 0xaa, 0xd2,
	0x5f, 0x31, 0x5d, 0x57, 0x38, 0x08, 0xc9, 0x0a, 0x12, 0xdd, 0xd7, 0x85, 0x03, 0xcb, 0xdf, 0xf9,
	0xe7, 0xa0, 0x19, 0x7d, 0xa3, 0xf8, 0x08, 0xb5, 0x4c, 0x82, 0x30, 0x99, 0x4c, 0xd6, 0xcd, 0x90,
	0x76, 0x9e, 0xdc, 0xc9, 0x31, 0xb9, 0x20, 0xde, 0x97, 0x9f, 0x7f, 0xbf, 0x4d, 0xb9, 0x78, 0x35,
	0xbc, 0xf5, 0x15, 0xe0, 0x1c, 0xcd, 0xe8, 0x20, 0xe1, 0xc7, 0xb7, 0x4e, 0x1b, 0xcf, 0x6c, 0x87,
	0xdc, 0x45, 0xb1, 0x7a, 0x8f, 0xb4, 0xde, 0x03, 0xbc, 0x32, 0xa9, 0x27, 0x6b, 0x5a, 0x77, 0xef,
	0xec, 0xc2, 0x73, 0xce, 0x2f, 0x3c, 0xe7, 0xcf, 0x85, 0xe7, 0x7c, 0xbd, 0xf4, 0x1a, 0xe7, 0x97,
	0x5e, 0xe3, 0xd7, 0xa5, 0xd7, 0xf8, 0xf4, 0x74, 0x2c, 0x80, 0xbb, 0xb4, 0x48, 0x39, 0xdd, 0xca,
	0x68, 0x5f, 0x9a, 0x31, 0x9f, 0x47, 0x83, 0x74, 0x10, 0xfb, 0x2d, 0xfd, 0x76, 0x9f, 0xff, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0x5f, 0xca, 0x38, 0x5a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the total set of minting parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	State(ctx context.Context, in *QueryStateRequest, opts ...grpc.CallOption) (*QueryStateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/nova.oracle.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) State(ctx context.Context, in *QueryStateRequest, opts ...grpc.CallOption) (*QueryStateResponse, error) {
	out := new(QueryStateResponse)
	err := c.cc.Invoke(ctx, "/nova.oracle.v1.Query/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of minting parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	State(context.Context, *QueryStateRequest) (*QueryStateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) State(ctx context.Context, req *QueryStateRequest) (*QueryStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.oracle.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nova.oracle.v1.Query/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).State(ctx, req.(*QueryStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nova.oracle.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "State",
			Handler:    _Query_State_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nova/oracle/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainDenom) > 0 {
		i -= len(m.ChainDenom)
		copy(dAtA[i:], m.ChainDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockProposer) > 0 {
		i -= len(m.BlockProposer)
		copy(dAtA[i:], m.BlockProposer)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BlockProposer)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.LastBlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.LastBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.Decimal != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coin.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Decimal != 0 {
		n += 1 + sovQuery(uint64(m.Decimal))
	}
	if m.LastBlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.LastBlockHeight))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BlockProposer)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeight", wireType)
			}
			m.LastBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockProposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockProposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

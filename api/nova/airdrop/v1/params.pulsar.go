// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package airdropv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Params                       protoreflect.MessageDescriptor
	fd_Params_snapshot_timestamp    protoreflect.FieldDescriptor
	fd_Params_claimable_timestamp   protoreflect.FieldDescriptor
	fd_Params_airdrop_end_timestamp protoreflect.FieldDescriptor
	fd_Params_airdrop_denom         protoreflect.FieldDescriptor
	fd_Params_quests_count          protoreflect.FieldDescriptor
	fd_Params_controller_address    protoreflect.FieldDescriptor
)

func init() {
	file_nova_airdrop_v1_params_proto_init()
	md_Params = File_nova_airdrop_v1_params_proto.Messages().ByName("Params")
	fd_Params_snapshot_timestamp = md_Params.Fields().ByName("snapshot_timestamp")
	fd_Params_claimable_timestamp = md_Params.Fields().ByName("claimable_timestamp")
	fd_Params_airdrop_end_timestamp = md_Params.Fields().ByName("airdrop_end_timestamp")
	fd_Params_airdrop_denom = md_Params.Fields().ByName("airdrop_denom")
	fd_Params_quests_count = md_Params.Fields().ByName("quests_count")
	fd_Params_controller_address = md_Params.Fields().ByName("controller_address")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_nova_airdrop_v1_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.SnapshotTimestamp != nil {
		value := protoreflect.ValueOfMessage(x.SnapshotTimestamp.ProtoReflect())
		if !f(fd_Params_snapshot_timestamp, value) {
			return
		}
	}
	if x.ClaimableTimestamp != nil {
		value := protoreflect.ValueOfMessage(x.ClaimableTimestamp.ProtoReflect())
		if !f(fd_Params_claimable_timestamp, value) {
			return
		}
	}
	if x.AirdropEndTimestamp != nil {
		value := protoreflect.ValueOfMessage(x.AirdropEndTimestamp.ProtoReflect())
		if !f(fd_Params_airdrop_end_timestamp, value) {
			return
		}
	}
	if x.AirdropDenom != "" {
		value := protoreflect.ValueOfString(x.AirdropDenom)
		if !f(fd_Params_airdrop_denom, value) {
			return
		}
	}
	if x.QuestsCount != int32(0) {
		value := protoreflect.ValueOfInt32(x.QuestsCount)
		if !f(fd_Params_quests_count, value) {
			return
		}
	}
	if x.ControllerAddress != "" {
		value := protoreflect.ValueOfString(x.ControllerAddress)
		if !f(fd_Params_controller_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "nova.airdrop.v1.Params.snapshot_timestamp":
		return x.SnapshotTimestamp != nil
	case "nova.airdrop.v1.Params.claimable_timestamp":
		return x.ClaimableTimestamp != nil
	case "nova.airdrop.v1.Params.airdrop_end_timestamp":
		return x.AirdropEndTimestamp != nil
	case "nova.airdrop.v1.Params.airdrop_denom":
		return x.AirdropDenom != ""
	case "nova.airdrop.v1.Params.quests_count":
		return x.QuestsCount != int32(0)
	case "nova.airdrop.v1.Params.controller_address":
		return x.ControllerAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nova.airdrop.v1.Params"))
		}
		panic(fmt.Errorf("message nova.airdrop.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "nova.airdrop.v1.Params.snapshot_timestamp":
		x.SnapshotTimestamp = nil
	case "nova.airdrop.v1.Params.claimable_timestamp":
		x.ClaimableTimestamp = nil
	case "nova.airdrop.v1.Params.airdrop_end_timestamp":
		x.AirdropEndTimestamp = nil
	case "nova.airdrop.v1.Params.airdrop_denom":
		x.AirdropDenom = ""
	case "nova.airdrop.v1.Params.quests_count":
		x.QuestsCount = int32(0)
	case "nova.airdrop.v1.Params.controller_address":
		x.ControllerAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nova.airdrop.v1.Params"))
		}
		panic(fmt.Errorf("message nova.airdrop.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "nova.airdrop.v1.Params.snapshot_timestamp":
		value := x.SnapshotTimestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "nova.airdrop.v1.Params.claimable_timestamp":
		value := x.ClaimableTimestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "nova.airdrop.v1.Params.airdrop_end_timestamp":
		value := x.AirdropEndTimestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "nova.airdrop.v1.Params.airdrop_denom":
		value := x.AirdropDenom
		return protoreflect.ValueOfString(value)
	case "nova.airdrop.v1.Params.quests_count":
		value := x.QuestsCount
		return protoreflect.ValueOfInt32(value)
	case "nova.airdrop.v1.Params.controller_address":
		value := x.ControllerAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nova.airdrop.v1.Params"))
		}
		panic(fmt.Errorf("message nova.airdrop.v1.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "nova.airdrop.v1.Params.snapshot_timestamp":
		x.SnapshotTimestamp = value.Message().Interface().(*timestamppb.Timestamp)
	case "nova.airdrop.v1.Params.claimable_timestamp":
		x.ClaimableTimestamp = value.Message().Interface().(*timestamppb.Timestamp)
	case "nova.airdrop.v1.Params.airdrop_end_timestamp":
		x.AirdropEndTimestamp = value.Message().Interface().(*timestamppb.Timestamp)
	case "nova.airdrop.v1.Params.airdrop_denom":
		x.AirdropDenom = value.Interface().(string)
	case "nova.airdrop.v1.Params.quests_count":
		x.QuestsCount = int32(value.Int())
	case "nova.airdrop.v1.Params.controller_address":
		x.ControllerAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nova.airdrop.v1.Params"))
		}
		panic(fmt.Errorf("message nova.airdrop.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "nova.airdrop.v1.Params.snapshot_timestamp":
		if x.SnapshotTimestamp == nil {
			x.SnapshotTimestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.SnapshotTimestamp.ProtoReflect())
	case "nova.airdrop.v1.Params.claimable_timestamp":
		if x.ClaimableTimestamp == nil {
			x.ClaimableTimestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.ClaimableTimestamp.ProtoReflect())
	case "nova.airdrop.v1.Params.airdrop_end_timestamp":
		if x.AirdropEndTimestamp == nil {
			x.AirdropEndTimestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.AirdropEndTimestamp.ProtoReflect())
	case "nova.airdrop.v1.Params.airdrop_denom":
		panic(fmt.Errorf("field airdrop_denom of message nova.airdrop.v1.Params is not mutable"))
	case "nova.airdrop.v1.Params.quests_count":
		panic(fmt.Errorf("field quests_count of message nova.airdrop.v1.Params is not mutable"))
	case "nova.airdrop.v1.Params.controller_address":
		panic(fmt.Errorf("field controller_address of message nova.airdrop.v1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nova.airdrop.v1.Params"))
		}
		panic(fmt.Errorf("message nova.airdrop.v1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "nova.airdrop.v1.Params.snapshot_timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "nova.airdrop.v1.Params.claimable_timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "nova.airdrop.v1.Params.airdrop_end_timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "nova.airdrop.v1.Params.airdrop_denom":
		return protoreflect.ValueOfString("")
	case "nova.airdrop.v1.Params.quests_count":
		return protoreflect.ValueOfInt32(int32(0))
	case "nova.airdrop.v1.Params.controller_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: nova.airdrop.v1.Params"))
		}
		panic(fmt.Errorf("message nova.airdrop.v1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in nova.airdrop.v1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.SnapshotTimestamp != nil {
			l = options.Size(x.SnapshotTimestamp)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ClaimableTimestamp != nil {
			l = options.Size(x.ClaimableTimestamp)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AirdropEndTimestamp != nil {
			l = options.Size(x.AirdropEndTimestamp)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AirdropDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.QuestsCount != 0 {
			n += 1 + runtime.Sov(uint64(x.QuestsCount))
		}
		l = len(x.ControllerAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ControllerAddress) > 0 {
			i -= len(x.ControllerAddress)
			copy(dAtA[i:], x.ControllerAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ControllerAddress)))
			i--
			dAtA[i] = 0x32
		}
		if x.QuestsCount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.QuestsCount))
			i--
			dAtA[i] = 0x28
		}
		if len(x.AirdropDenom) > 0 {
			i -= len(x.AirdropDenom)
			copy(dAtA[i:], x.AirdropDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AirdropDenom)))
			i--
			dAtA[i] = 0x22
		}
		if x.AirdropEndTimestamp != nil {
			encoded, err := options.Marshal(x.AirdropEndTimestamp)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if x.ClaimableTimestamp != nil {
			encoded, err := options.Marshal(x.ClaimableTimestamp)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.SnapshotTimestamp != nil {
			encoded, err := options.Marshal(x.SnapshotTimestamp)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SnapshotTimestamp", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.SnapshotTimestamp == nil {
					x.SnapshotTimestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SnapshotTimestamp); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClaimableTimestamp", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ClaimableTimestamp == nil {
					x.ClaimableTimestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ClaimableTimestamp); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AirdropEndTimestamp", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.AirdropEndTimestamp == nil {
					x.AirdropEndTimestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AirdropEndTimestamp); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AirdropDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AirdropDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field QuestsCount", wireType)
				}
				x.QuestsCount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.QuestsCount |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ControllerAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ControllerAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: nova/airdrop/v1/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The snapshot date based for the distribution of the airdrop.
	SnapshotTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=snapshot_timestamp,json=snapshotTimestamp,proto3" json:"snapshot_timestamp,omitempty"`
	// THe time when you can claim your airdrop nova tokens.
	ClaimableTimestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=claimable_timestamp,json=claimableTimestamp,proto3" json:"claimable_timestamp,omitempty"`
	// THe time when the user no longer can claim the airdrop tokens.
	AirdropEndTimestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=airdrop_end_timestamp,json=airdropEndTimestamp,proto3" json:"airdrop_end_timestamp,omitempty"`
	// The denom for the airdrop coin.
	AirdropDenom string `protobuf:"bytes,4,opt,name=airdrop_denom,json=airdropDenom,proto3" json:"airdrop_denom,omitempty"`
	// the number of quests user to do
	QuestsCount int32 `protobuf:"varint,5,opt,name=quests_count,json=questsCount,proto3" json:"quests_count,omitempty"`
	// controller address is responsible to check the user has performed the social quest (e.g. twitter, facebook or etc)
	ControllerAddress string `protobuf:"bytes,6,opt,name=controller_address,json=controllerAddress,proto3" json:"controller_address,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nova_airdrop_v1_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_nova_airdrop_v1_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetSnapshotTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.SnapshotTimestamp
	}
	return nil
}

func (x *Params) GetClaimableTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ClaimableTimestamp
	}
	return nil
}

func (x *Params) GetAirdropEndTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.AirdropEndTimestamp
	}
	return nil
}

func (x *Params) GetAirdropDenom() string {
	if x != nil {
		return x.AirdropDenom
	}
	return ""
}

func (x *Params) GetQuestsCount() int32 {
	if x != nil {
		return x.QuestsCount
	}
	return 0
}

func (x *Params) GetControllerAddress() string {
	if x != nil {
		return x.ControllerAddress
	}
	return ""
}

var File_nova_airdrop_v1_params_proto protoreflect.FileDescriptor

var file_nova_airdrop_v1_params_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6e, 0x6f, 0x76, 0x61, 0x2f, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x6e, 0x6f, 0x76, 0x61, 0x2e, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x53, 0x0a, 0x12, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90,
	0xdf, 0x1f, 0x01, 0x52, 0x11, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x55, 0x0a, 0x13, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x58, 0x0a,
	0x15, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf,
	0x1f, 0x01, 0x52, 0x13, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x69, 0x72, 0x64, 0x72,
	0x6f, 0x70, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x04,
	0x98, 0xa0, 0x1f, 0x00, 0x42, 0xbb, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x76,
	0x61, 0x2e, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x72, 0x69, 0x6e, 0x61, 0x2d, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x76,
	0x61, 0x2f, 0x61, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x69, 0x72,
	0x64, 0x72, 0x6f, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x41, 0x58, 0xaa, 0x02, 0x0f, 0x4e,
	0x6f, 0x76, 0x61, 0x2e, 0x41, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0f, 0x4e, 0x6f, 0x76, 0x61, 0x5c, 0x41, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1b, 0x4e, 0x6f, 0x76, 0x61, 0x5c, 0x41, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x4e, 0x6f, 0x76, 0x61, 0x3a, 0x3a, 0x41, 0x69, 0x72, 0x64, 0x72, 0x6f, 0x70, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nova_airdrop_v1_params_proto_rawDescOnce sync.Once
	file_nova_airdrop_v1_params_proto_rawDescData = file_nova_airdrop_v1_params_proto_rawDesc
)

func file_nova_airdrop_v1_params_proto_rawDescGZIP() []byte {
	file_nova_airdrop_v1_params_proto_rawDescOnce.Do(func() {
		file_nova_airdrop_v1_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_nova_airdrop_v1_params_proto_rawDescData)
	})
	return file_nova_airdrop_v1_params_proto_rawDescData
}

var file_nova_airdrop_v1_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nova_airdrop_v1_params_proto_goTypes = []interface{}{
	(*Params)(nil),                // 0: nova.airdrop.v1.Params
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_nova_airdrop_v1_params_proto_depIdxs = []int32{
	1, // 0: nova.airdrop.v1.Params.snapshot_timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: nova.airdrop.v1.Params.claimable_timestamp:type_name -> google.protobuf.Timestamp
	1, // 2: nova.airdrop.v1.Params.airdrop_end_timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nova_airdrop_v1_params_proto_init() }
func file_nova_airdrop_v1_params_proto_init() {
	if File_nova_airdrop_v1_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nova_airdrop_v1_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nova_airdrop_v1_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nova_airdrop_v1_params_proto_goTypes,
		DependencyIndexes: file_nova_airdrop_v1_params_proto_depIdxs,
		MessageInfos:      file_nova_airdrop_v1_params_proto_msgTypes,
	}.Build()
	File_nova_airdrop_v1_params_proto = out.File
	file_nova_airdrop_v1_params_proto_rawDesc = nil
	file_nova_airdrop_v1_params_proto_goTypes = nil
	file_nova_airdrop_v1_params_proto_depIdxs = nil
}
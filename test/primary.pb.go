// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: test/primary.proto

package test

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Mode int32

const (
	Mode_UNKNOWN Mode = 0
	Mode_ON      Mode = 1
	Mode_OFF     Mode = 2
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "UNKNOWN",
		1: "ON",
		2: "OFF",
	}
	Mode_value = map[string]int32{
		"UNKNOWN": 0,
		"ON":      1,
		"OFF":     2,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_test_primary_proto_enumTypes[0].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_test_primary_proto_enumTypes[0]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_test_primary_proto_rawDescGZIP(), []int{0}
}

// Test message definition.
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Str string field
	Str string `protobuf:"bytes,1,opt,name=Str,proto3" json:"Str,omitempty"`
	// Int32 int32 field
	Int32 int32 `protobuf:"varint,2,opt,name=Int32,proto3" json:"Int32,omitempty"`
	// Int64 int64 field
	Int64 int64 `protobuf:"varint,3,opt,name=Int64,proto3" json:"Int64,omitempty"`
	// Float float field
	Float float32 `protobuf:"fixed32,4,opt,name=Float,proto3" json:"Float,omitempty"`
	// Double double field
	Double float64 `protobuf:"fixed64,5,opt,name=Double,proto3" json:"Double,omitempty"`
	// Bool bool field
	Bool bool `protobuf:"varint,6,opt,name=Bool,proto3" json:"Bool,omitempty"`
	// Bytes byte[] field
	Bytes []byte `protobuf:"bytes,7,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	// StringList []string field
	StringList []string `protobuf:"bytes,16,rep,name=StringList,proto3" json:"StringList,omitempty"`
	// Nested nested message field, non-nullable
	Nested *Nested `protobuf:"bytes,22,opt,name=Nested,proto3" json:"Nested,omitempty"`
	// NestedList nested message array
	NestedList []*Nested `protobuf:"bytes,25,rep,name=NestedList,proto3" json:"NestedList,omitempty"`
	// Map normal map
	Map map[string]string `protobuf:"bytes,27,rep,name=Map,proto3" json:"Map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// MapObject is the object map
	NestedMap map[string]*Nested `protobuf:"bytes,29,rep,name=NestedMap,proto3" json:"NestedMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Mode is the enum value
	Mode Mode `protobuf:"varint,31,opt,name=Mode,proto3,enum=test.Mode" json:"Mode,omitempty"`
	// Types that are assignable to OneOf:
	//
	//	*Test_Branch1
	//	*Test_Branch2
	//	*Test_Branch3
	OneOf isTest_OneOf `protobuf_oneof:"OneOf"`
	// Required string field
	Required string `protobuf:"bytes,38,opt,name=required,proto3" json:"required,omitempty"`
	// Computed string field
	Computed string `protobuf:"bytes,39,opt,name=computed,proto3" json:"computed,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_primary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_test_primary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_test_primary_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Test) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Test) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Test) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Test) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Test) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Test) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Test) GetStringList() []string {
	if x != nil {
		return x.StringList
	}
	return nil
}

func (x *Test) GetNested() *Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *Test) GetNestedList() []*Nested {
	if x != nil {
		return x.NestedList
	}
	return nil
}

func (x *Test) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *Test) GetNestedMap() map[string]*Nested {
	if x != nil {
		return x.NestedMap
	}
	return nil
}

func (x *Test) GetMode() Mode {
	if x != nil {
		return x.Mode
	}
	return Mode_UNKNOWN
}

func (m *Test) GetOneOf() isTest_OneOf {
	if m != nil {
		return m.OneOf
	}
	return nil
}

func (x *Test) GetBranch1() *Branch1 {
	if x, ok := x.GetOneOf().(*Test_Branch1); ok {
		return x.Branch1
	}
	return nil
}

func (x *Test) GetBranch2() *Branch2 {
	if x, ok := x.GetOneOf().(*Test_Branch2); ok {
		return x.Branch2
	}
	return nil
}

func (x *Test) GetBranch3() string {
	if x, ok := x.GetOneOf().(*Test_Branch3); ok {
		return x.Branch3
	}
	return ""
}

func (x *Test) GetRequired() string {
	if x != nil {
		return x.Required
	}
	return ""
}

func (x *Test) GetComputed() string {
	if x != nil {
		return x.Computed
	}
	return ""
}

type isTest_OneOf interface {
	isTest_OneOf()
}

type Test_Branch1 struct {
	// Branch1 is the first oneOf branch
	Branch1 *Branch1 `protobuf:"bytes,33,opt,name=Branch1,proto3,oneof"`
}

type Test_Branch2 struct {
	// Branch2 is the second oneOf branch
	Branch2 *Branch2 `protobuf:"bytes,34,opt,name=Branch2,proto3,oneof"`
}

type Test_Branch3 struct {
	// Branch3 is the third branch which is simple string
	Branch3 string `protobuf:"bytes,35,opt,name=Branch3,proto3,oneof"`
}

func (*Test_Branch1) isTest_OneOf() {}

func (*Test_Branch2) isTest_OneOf() {}

func (*Test_Branch3) isTest_OneOf() {}

// EmptyMessageBranch message for empty oneof branch
type EmptyMessageBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessageBranch) Reset() {
	*x = EmptyMessageBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_primary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessageBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessageBranch) ProtoMessage() {}

func (x *EmptyMessageBranch) ProtoReflect() protoreflect.Message {
	mi := &file_test_primary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessageBranch.ProtoReflect.Descriptor instead.
func (*EmptyMessageBranch) Descriptor() ([]byte, []int) {
	return file_test_primary_proto_rawDescGZIP(), []int{1}
}

// Nested message definition
type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Str string field
	Str string `protobuf:"bytes,1,opt,name=Str,proto3" json:"Str,omitempty"`
	// Nested repeated nested messages
	OtherNestedList []*OtherNested `protobuf:"bytes,2,rep,name=OtherNestedList,proto3" json:"OtherNestedList,omitempty"`
	// Nested map repeated nested messages
	Map map[string]string `protobuf:"bytes,3,rep,name=Map,proto3" json:"Map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// MapObjectNested nested object map
	MapObjectNested map[string]*OtherNested `protobuf:"bytes,4,rep,name=MapObjectNested,proto3" json:"MapObjectNested,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_primary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_test_primary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_test_primary_proto_rawDescGZIP(), []int{2}
}

func (x *Nested) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Nested) GetOtherNestedList() []*OtherNested {
	if x != nil {
		return x.OtherNestedList
	}
	return nil
}

func (x *Nested) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *Nested) GetMapObjectNested() map[string]*OtherNested {
	if x != nil {
		return x.MapObjectNested
	}
	return nil
}

// OtherNested message nested into nested message
type OtherNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Str string field
	Str string `protobuf:"bytes,1,opt,name=Str,proto3" json:"Str,omitempty"`
}

func (x *OtherNested) Reset() {
	*x = OtherNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_primary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherNested) ProtoMessage() {}

func (x *OtherNested) ProtoReflect() protoreflect.Message {
	mi := &file_test_primary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherNested.ProtoReflect.Descriptor instead.
func (*OtherNested) Descriptor() ([]byte, []int) {
	return file_test_primary_proto_rawDescGZIP(), []int{3}
}

func (x *OtherNested) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

// Branch1 message is OneOf branch 1
type Branch1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Str string field
	Str string `protobuf:"bytes,1,opt,name=Str,proto3" json:"Str,omitempty"`
}

func (x *Branch1) Reset() {
	*x = Branch1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_primary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch1) ProtoMessage() {}

func (x *Branch1) ProtoReflect() protoreflect.Message {
	mi := &file_test_primary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch1.ProtoReflect.Descriptor instead.
func (*Branch1) Descriptor() ([]byte, []int) {
	return file_test_primary_proto_rawDescGZIP(), []int{4}
}

func (x *Branch1) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

// Branch2 message is OneOf branch 2
type Branch2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Int32 int field
	Int32 int32 `protobuf:"varint,1,opt,name=Int32,proto3" json:"Int32,omitempty"`
}

func (x *Branch2) Reset() {
	*x = Branch2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_primary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch2) ProtoMessage() {}

func (x *Branch2) ProtoReflect() protoreflect.Message {
	mi := &file_test_primary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch2.ProtoReflect.Descriptor instead.
func (*Branch2) Descriptor() ([]byte, []int) {
	return file_test_primary_proto_rawDescGZIP(), []int{5}
}

func (x *Branch2) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

var File_test_primary_proto protoreflect.FileDescriptor

var file_test_primary_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd1, 0x05, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x74, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x10, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x06, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x0a, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x37, 0x0a, 0x09, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x4d, 0x61, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x31, 0x18, 0x21,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x31, 0x48, 0x00, 0x52, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x31, 0x12, 0x29,
	0x0a, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x32, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x32, 0x48, 0x00,
	0x52, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x32, 0x12, 0x1a, 0x0a, 0x07, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x33, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x33, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x64, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x4a, 0x0a, 0x0e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x4f,
	0x6e, 0x65, 0x4f, 0x66, 0x22, 0x14, 0x0a, 0x12, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0xdc, 0x02, 0x0a, 0x06, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x53, 0x74, 0x72, 0x12, 0x3b, 0x0a, 0x0f, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x52, 0x0f, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x4b, 0x0a,
	0x0f, 0x4d, 0x61, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x2e, 0x4d, 0x61, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x4d, 0x61, 0x70, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x55, 0x0a, 0x14, 0x4d, 0x61, 0x70, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x0b, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x74, 0x72, 0x22, 0x1b, 0x0a, 0x07, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x53, 0x74, 0x72, 0x22, 0x1f, 0x0a, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x2a, 0x24, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x61,
	0x6d, 0x61, 0x77, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_primary_proto_rawDescOnce sync.Once
	file_test_primary_proto_rawDescData = file_test_primary_proto_rawDesc
)

func file_test_primary_proto_rawDescGZIP() []byte {
	file_test_primary_proto_rawDescOnce.Do(func() {
		file_test_primary_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_primary_proto_rawDescData)
	})
	return file_test_primary_proto_rawDescData
}

var file_test_primary_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_primary_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_test_primary_proto_goTypes = []interface{}{
	(Mode)(0),                  // 0: test.Mode
	(*Test)(nil),               // 1: test.Test
	(*EmptyMessageBranch)(nil), // 2: test.EmptyMessageBranch
	(*Nested)(nil),             // 3: test.Nested
	(*OtherNested)(nil),        // 4: test.OtherNested
	(*Branch1)(nil),            // 5: test.Branch1
	(*Branch2)(nil),            // 6: test.Branch2
	nil,                        // 7: test.Test.MapEntry
	nil,                        // 8: test.Test.NestedMapEntry
	nil,                        // 9: test.Nested.MapEntry
	nil,                        // 10: test.Nested.MapObjectNestedEntry
}
var file_test_primary_proto_depIdxs = []int32{
	3,  // 0: test.Test.Nested:type_name -> test.Nested
	3,  // 1: test.Test.NestedList:type_name -> test.Nested
	7,  // 2: test.Test.Map:type_name -> test.Test.MapEntry
	8,  // 3: test.Test.NestedMap:type_name -> test.Test.NestedMapEntry
	0,  // 4: test.Test.Mode:type_name -> test.Mode
	5,  // 5: test.Test.Branch1:type_name -> test.Branch1
	6,  // 6: test.Test.Branch2:type_name -> test.Branch2
	4,  // 7: test.Nested.OtherNestedList:type_name -> test.OtherNested
	9,  // 8: test.Nested.Map:type_name -> test.Nested.MapEntry
	10, // 9: test.Nested.MapObjectNested:type_name -> test.Nested.MapObjectNestedEntry
	3,  // 10: test.Test.NestedMapEntry.value:type_name -> test.Nested
	4,  // 11: test.Nested.MapObjectNestedEntry.value:type_name -> test.OtherNested
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_test_primary_proto_init() }
func file_test_primary_proto_init() {
	if File_test_primary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_primary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_test_primary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessageBranch); i {
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
		file_test_primary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested); i {
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
		file_test_primary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherNested); i {
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
		file_test_primary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch1); i {
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
		file_test_primary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch2); i {
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
	file_test_primary_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Test_Branch1)(nil),
		(*Test_Branch2)(nil),
		(*Test_Branch3)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_primary_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_primary_proto_goTypes,
		DependencyIndexes: file_test_primary_proto_depIdxs,
		EnumInfos:         file_test_primary_proto_enumTypes,
		MessageInfos:      file_test_primary_proto_msgTypes,
	}.Build()
	File_test_primary_proto = out.File
	file_test_primary_proto_rawDesc = nil
	file_test_primary_proto_goTypes = nil
	file_test_primary_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: ncbi/datasets/v1alpha1/mongo_options.proto

package v1alpha1

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// Field Options related to MongoDB fields
type MongoMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *MongoMessageOptions) Reset() {
	*x = MongoMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoMessageOptions) ProtoMessage() {}

func (x *MongoMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoMessageOptions.ProtoReflect.Descriptor instead.
func (*MongoMessageOptions) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescGZIP(), []int{0}
}

func (x *MongoMessageOptions) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *MongoMessageOptions) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Field Options related to MongoDB fields
type MongoFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Required bool                       `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
	Spec     *MongoFieldOptions_ObjSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *MongoFieldOptions) Reset() {
	*x = MongoFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoFieldOptions) ProtoMessage() {}

func (x *MongoFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoFieldOptions.ProtoReflect.Descriptor instead.
func (*MongoFieldOptions) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescGZIP(), []int{1}
}

func (x *MongoFieldOptions) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *MongoFieldOptions) GetSpec() *MongoFieldOptions_ObjSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type MongoFieldOptions_ObjSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BsonType string `protobuf:"bytes,2,opt,name=bson_type,json=bsonType,proto3" json:"bson_type,omitempty"`
	// Types that are assignable to Content:
	//	*MongoFieldOptions_ObjSpec_Properties
	//	*MongoFieldOptions_ObjSpec_Items
	//	*MongoFieldOptions_ObjSpec_AdditionalProperties
	Content isMongoFieldOptions_ObjSpec_Content `protobuf_oneof:"content"`
}

func (x *MongoFieldOptions_ObjSpec) Reset() {
	*x = MongoFieldOptions_ObjSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoFieldOptions_ObjSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoFieldOptions_ObjSpec) ProtoMessage() {}

func (x *MongoFieldOptions_ObjSpec) ProtoReflect() protoreflect.Message {
	mi := &file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoFieldOptions_ObjSpec.ProtoReflect.Descriptor instead.
func (*MongoFieldOptions_ObjSpec) Descriptor() ([]byte, []int) {
	return file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MongoFieldOptions_ObjSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MongoFieldOptions_ObjSpec) GetBsonType() string {
	if x != nil {
		return x.BsonType
	}
	return ""
}

func (m *MongoFieldOptions_ObjSpec) GetContent() isMongoFieldOptions_ObjSpec_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *MongoFieldOptions_ObjSpec) GetProperties() *MongoFieldOptions_ObjSpec {
	if x, ok := x.GetContent().(*MongoFieldOptions_ObjSpec_Properties); ok {
		return x.Properties
	}
	return nil
}

func (x *MongoFieldOptions_ObjSpec) GetItems() *MongoFieldOptions_ObjSpec {
	if x, ok := x.GetContent().(*MongoFieldOptions_ObjSpec_Items); ok {
		return x.Items
	}
	return nil
}

func (x *MongoFieldOptions_ObjSpec) GetAdditionalProperties() *MongoFieldOptions_ObjSpec {
	if x, ok := x.GetContent().(*MongoFieldOptions_ObjSpec_AdditionalProperties); ok {
		return x.AdditionalProperties
	}
	return nil
}

type isMongoFieldOptions_ObjSpec_Content interface {
	isMongoFieldOptions_ObjSpec_Content()
}

type MongoFieldOptions_ObjSpec_Properties struct {
	Properties *MongoFieldOptions_ObjSpec `protobuf:"bytes,3,opt,name=properties,proto3,oneof"`
}

type MongoFieldOptions_ObjSpec_Items struct {
	Items *MongoFieldOptions_ObjSpec `protobuf:"bytes,4,opt,name=items,proto3,oneof"`
}

type MongoFieldOptions_ObjSpec_AdditionalProperties struct {
	AdditionalProperties *MongoFieldOptions_ObjSpec `protobuf:"bytes,5,opt,name=additional_properties,json=additionalProperties,proto3,oneof"`
}

func (*MongoFieldOptions_ObjSpec_Properties) isMongoFieldOptions_ObjSpec_Content() {}

func (*MongoFieldOptions_ObjSpec_Items) isMongoFieldOptions_ObjSpec_Content() {}

func (*MongoFieldOptions_ObjSpec_AdditionalProperties) isMongoFieldOptions_ObjSpec_Content() {}

var file_ncbi_datasets_v1alpha1_mongo_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*MongoFieldOptions)(nil),
		Field:         52001,
		Name:          "ncbi.datasets.v1alpha1.mongo_field",
		Tag:           "bytes,52001,opt,name=mongo_field",
		Filename:      "ncbi/datasets/v1alpha1/mongo_options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*MongoMessageOptions)(nil),
		Field:         51000,
		Name:          "ncbi.datasets.v1alpha1.mongo_collection",
		Tag:           "bytes,51000,opt,name=mongo_collection",
		Filename:      "ncbi/datasets/v1alpha1/mongo_options.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional ncbi.datasets.v1alpha1.MongoFieldOptions mongo_field = 52001;
	E_MongoField = &file_ncbi_datasets_v1alpha1_mongo_options_proto_extTypes[0]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional ncbi.datasets.v1alpha1.MongoMessageOptions mongo_collection = 51000;
	E_MongoCollection = &file_ncbi_datasets_v1alpha1_mongo_options_proto_extTypes[1]
)

var File_ncbi_datasets_v1alpha1_mongo_options_proto protoreflect.FileDescriptor

var file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x63,
	0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x13, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0xc8, 0x03, 0x0a, 0x11, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x45, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x1a, 0xcf, 0x02, 0x0a, 0x07, 0x4f, 0x62, 0x6a,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x73, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x73, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x63, 0x62, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x63, 0x62, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x68, 0x0a, 0x15, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f,
	0x6e, 0x67, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4f, 0x62, 0x6a, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x14, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42,
	0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x6b, 0x0a, 0x0b, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa1, 0x96, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x6d, 0x6f, 0x6e,
	0x67, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x79, 0x0a, 0x10, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x8e, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x63, 0x62, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6f,
	0x6e, 0x67, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x1b, 0x5a, 0x16, 0x6e, 0x63, 0x62, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xf8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescOnce sync.Once
	file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescData = file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDesc
)

func file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescGZIP() []byte {
	file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescOnce.Do(func() {
		file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescData)
	})
	return file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDescData
}

var file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ncbi_datasets_v1alpha1_mongo_options_proto_goTypes = []interface{}{
	(*MongoMessageOptions)(nil),       // 0: ncbi.datasets.v1alpha1.MongoMessageOptions
	(*MongoFieldOptions)(nil),         // 1: ncbi.datasets.v1alpha1.MongoFieldOptions
	(*MongoFieldOptions_ObjSpec)(nil), // 2: ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec
	(*descriptor.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
	(*descriptor.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
}
var file_ncbi_datasets_v1alpha1_mongo_options_proto_depIdxs = []int32{
	2, // 0: ncbi.datasets.v1alpha1.MongoFieldOptions.spec:type_name -> ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec
	2, // 1: ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec.properties:type_name -> ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec
	2, // 2: ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec.items:type_name -> ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec
	2, // 3: ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec.additional_properties:type_name -> ncbi.datasets.v1alpha1.MongoFieldOptions.ObjSpec
	3, // 4: ncbi.datasets.v1alpha1.mongo_field:extendee -> google.protobuf.FieldOptions
	4, // 5: ncbi.datasets.v1alpha1.mongo_collection:extendee -> google.protobuf.MessageOptions
	1, // 6: ncbi.datasets.v1alpha1.mongo_field:type_name -> ncbi.datasets.v1alpha1.MongoFieldOptions
	0, // 7: ncbi.datasets.v1alpha1.mongo_collection:type_name -> ncbi.datasets.v1alpha1.MongoMessageOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	6, // [6:8] is the sub-list for extension type_name
	4, // [4:6] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ncbi_datasets_v1alpha1_mongo_options_proto_init() }
func file_ncbi_datasets_v1alpha1_mongo_options_proto_init() {
	if File_ncbi_datasets_v1alpha1_mongo_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoMessageOptions); i {
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
		file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoFieldOptions); i {
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
		file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoFieldOptions_ObjSpec); i {
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
	file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MongoFieldOptions_ObjSpec_Properties)(nil),
		(*MongoFieldOptions_ObjSpec_Items)(nil),
		(*MongoFieldOptions_ObjSpec_AdditionalProperties)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_ncbi_datasets_v1alpha1_mongo_options_proto_goTypes,
		DependencyIndexes: file_ncbi_datasets_v1alpha1_mongo_options_proto_depIdxs,
		MessageInfos:      file_ncbi_datasets_v1alpha1_mongo_options_proto_msgTypes,
		ExtensionInfos:    file_ncbi_datasets_v1alpha1_mongo_options_proto_extTypes,
	}.Build()
	File_ncbi_datasets_v1alpha1_mongo_options_proto = out.File
	file_ncbi_datasets_v1alpha1_mongo_options_proto_rawDesc = nil
	file_ncbi_datasets_v1alpha1_mongo_options_proto_goTypes = nil
	file_ncbi_datasets_v1alpha1_mongo_options_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: zitadel/auth_n_key.proto

package authn

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	object "github.com/isodude/zitadel-go/v2/pkg/client/zitadel/object"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyType int32

const (
	KeyType_KEY_TYPE_UNSPECIFIED KeyType = 0
	KeyType_KEY_TYPE_JSON        KeyType = 1
)

// Enum value maps for KeyType.
var (
	KeyType_name = map[int32]string{
		0: "KEY_TYPE_UNSPECIFIED",
		1: "KEY_TYPE_JSON",
	}
	KeyType_value = map[string]int32{
		"KEY_TYPE_UNSPECIFIED": 0,
		"KEY_TYPE_JSON":        1,
	}
)

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}

func (x KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_zitadel_auth_n_key_proto_enumTypes[0].Descriptor()
}

func (KeyType) Type() protoreflect.EnumType {
	return &file_zitadel_auth_n_key_proto_enumTypes[0]
}

func (x KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyType.Descriptor instead.
func (KeyType) EnumDescriptor() ([]byte, []int) {
	return file_zitadel_auth_n_key_proto_rawDescGZIP(), []int{0}
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Details        *object.ObjectDetails  `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Type           KeyType                `protobuf:"varint,3,opt,name=type,proto3,enum=zitadel.authn.v1.KeyType" json:"type,omitempty"`
	ExpirationDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_auth_n_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_auth_n_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_zitadel_auth_n_key_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Key) GetDetails() *object.ObjectDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Key) GetType() KeyType {
	if x != nil {
		return x.Type
	}
	return KeyType_KEY_TYPE_UNSPECIFIED
}

func (x *Key) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

var File_zitadel_auth_n_key_proto protoreflect.FileDescriptor

var file_zitadel_auth_n_key_proto_rawDesc = []byte{
	0x0a, 0x18, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6e,
	0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x7a, 0x69, 0x74, 0x61,
	0x64, 0x65, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x7a, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x92, 0x41, 0x15, 0x4a, 0x13, 0x22, 0x36,
	0x39, 0x36, 0x32, 0x39, 0x30, 0x32, 0x33, 0x39, 0x30, 0x36, 0x34, 0x38, 0x38, 0x33, 0x33, 0x34,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x1d, 0x92, 0x41, 0x1a, 0x32, 0x18, 0x74, 0x68, 0x65, 0x20, 0x66, 0x69,
	0x6c, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6b,
	0x65, 0x79, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x3e,
	0x92, 0x41, 0x3b, 0x32, 0x1a, 0x74, 0x68, 0x65, 0x20, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20,
	0x6b, 0x65, 0x79, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4a,
	0x1d, 0x22, 0x33, 0x30, 0x31, 0x39, 0x2d, 0x30, 0x34, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x38, 0x3a,
	0x34, 0x35, 0x3a, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x5a, 0x22, 0x52, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x2a, 0x36,
	0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4b, 0x45, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4b, 0x45, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74,
	0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_auth_n_key_proto_rawDescOnce sync.Once
	file_zitadel_auth_n_key_proto_rawDescData = file_zitadel_auth_n_key_proto_rawDesc
)

func file_zitadel_auth_n_key_proto_rawDescGZIP() []byte {
	file_zitadel_auth_n_key_proto_rawDescOnce.Do(func() {
		file_zitadel_auth_n_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_auth_n_key_proto_rawDescData)
	})
	return file_zitadel_auth_n_key_proto_rawDescData
}

var file_zitadel_auth_n_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zitadel_auth_n_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_zitadel_auth_n_key_proto_goTypes = []interface{}{
	(KeyType)(0),                  // 0: zitadel.authn.v1.KeyType
	(*Key)(nil),                   // 1: zitadel.authn.v1.Key
	(*object.ObjectDetails)(nil),  // 2: zitadel.v1.ObjectDetails
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_zitadel_auth_n_key_proto_depIdxs = []int32{
	2, // 0: zitadel.authn.v1.Key.details:type_name -> zitadel.v1.ObjectDetails
	0, // 1: zitadel.authn.v1.Key.type:type_name -> zitadel.authn.v1.KeyType
	3, // 2: zitadel.authn.v1.Key.expiration_date:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_zitadel_auth_n_key_proto_init() }
func file_zitadel_auth_n_key_proto_init() {
	if File_zitadel_auth_n_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_auth_n_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
			RawDescriptor: file_zitadel_auth_n_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_auth_n_key_proto_goTypes,
		DependencyIndexes: file_zitadel_auth_n_key_proto_depIdxs,
		EnumInfos:         file_zitadel_auth_n_key_proto_enumTypes,
		MessageInfos:      file_zitadel_auth_n_key_proto_msgTypes,
	}.Build()
	File_zitadel_auth_n_key_proto = out.File
	file_zitadel_auth_n_key_proto_rawDesc = nil
	file_zitadel_auth_n_key_proto_goTypes = nil
	file_zitadel_auth_n_key_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FHHBJLJCHEA.proto

package proto

import (
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

type FHHBJLJCHEA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num         uint32 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	DJHODDOHMNB uint32 `protobuf:"varint,11,opt,name=DJHODDOHMNB,proto3" json:"DJHODDOHMNB,omitempty"`
}

func (x *FHHBJLJCHEA) Reset() {
	*x = FHHBJLJCHEA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FHHBJLJCHEA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FHHBJLJCHEA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FHHBJLJCHEA) ProtoMessage() {}

func (x *FHHBJLJCHEA) ProtoReflect() protoreflect.Message {
	mi := &file_FHHBJLJCHEA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FHHBJLJCHEA.ProtoReflect.Descriptor instead.
func (*FHHBJLJCHEA) Descriptor() ([]byte, []int) {
	return file_FHHBJLJCHEA_proto_rawDescGZIP(), []int{0}
}

func (x *FHHBJLJCHEA) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *FHHBJLJCHEA) GetDJHODDOHMNB() uint32 {
	if x != nil {
		return x.DJHODDOHMNB
	}
	return 0
}

var File_FHHBJLJCHEA_proto protoreflect.FileDescriptor

var file_FHHBJLJCHEA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x48, 0x48, 0x42, 0x4a, 0x4c, 0x4a, 0x43, 0x48, 0x45, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0b, 0x46, 0x48, 0x48, 0x42, 0x4a, 0x4c, 0x4a, 0x43, 0x48,
	0x45, 0x41, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4a, 0x48, 0x4f, 0x44, 0x44, 0x4f, 0x48,
	0x4d, 0x4e, 0x42, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4a, 0x48, 0x4f, 0x44,
	0x44, 0x4f, 0x48, 0x4d, 0x4e, 0x42, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FHHBJLJCHEA_proto_rawDescOnce sync.Once
	file_FHHBJLJCHEA_proto_rawDescData = file_FHHBJLJCHEA_proto_rawDesc
)

func file_FHHBJLJCHEA_proto_rawDescGZIP() []byte {
	file_FHHBJLJCHEA_proto_rawDescOnce.Do(func() {
		file_FHHBJLJCHEA_proto_rawDescData = protoimpl.X.CompressGZIP(file_FHHBJLJCHEA_proto_rawDescData)
	})
	return file_FHHBJLJCHEA_proto_rawDescData
}

var file_FHHBJLJCHEA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FHHBJLJCHEA_proto_goTypes = []interface{}{
	(*FHHBJLJCHEA)(nil), // 0: FHHBJLJCHEA
}
var file_FHHBJLJCHEA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FHHBJLJCHEA_proto_init() }
func file_FHHBJLJCHEA_proto_init() {
	if File_FHHBJLJCHEA_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FHHBJLJCHEA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FHHBJLJCHEA); i {
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
			RawDescriptor: file_FHHBJLJCHEA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FHHBJLJCHEA_proto_goTypes,
		DependencyIndexes: file_FHHBJLJCHEA_proto_depIdxs,
		MessageInfos:      file_FHHBJLJCHEA_proto_msgTypes,
	}.Build()
	File_FHHBJLJCHEA_proto = out.File
	file_FHHBJLJCHEA_proto_rawDesc = nil
	file_FHHBJLJCHEA_proto_goTypes = nil
	file_FHHBJLJCHEA_proto_depIdxs = nil
}

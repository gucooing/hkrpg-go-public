// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MonopolySocialEventStatus.proto

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

type MonopolySocialEventStatus int32

const (
	MonopolySocialEventStatus_MONOPOLY_SOCIAL_EVENT_STATUS_NONE                  MonopolySocialEventStatus = 0
	MonopolySocialEventStatus_MONOPOLY_SOCIAL_EVENT_STATUS_WAITING_SELECT_FRIEND MonopolySocialEventStatus = 1
)

// Enum value maps for MonopolySocialEventStatus.
var (
	MonopolySocialEventStatus_name = map[int32]string{
		0: "MONOPOLY_SOCIAL_EVENT_STATUS_NONE",
		1: "MONOPOLY_SOCIAL_EVENT_STATUS_WAITING_SELECT_FRIEND",
	}
	MonopolySocialEventStatus_value = map[string]int32{
		"MONOPOLY_SOCIAL_EVENT_STATUS_NONE":                  0,
		"MONOPOLY_SOCIAL_EVENT_STATUS_WAITING_SELECT_FRIEND": 1,
	}
)

func (x MonopolySocialEventStatus) Enum() *MonopolySocialEventStatus {
	p := new(MonopolySocialEventStatus)
	*p = x
	return p
}

func (x MonopolySocialEventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MonopolySocialEventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_MonopolySocialEventStatus_proto_enumTypes[0].Descriptor()
}

func (MonopolySocialEventStatus) Type() protoreflect.EnumType {
	return &file_MonopolySocialEventStatus_proto_enumTypes[0]
}

func (x MonopolySocialEventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MonopolySocialEventStatus.Descriptor instead.
func (MonopolySocialEventStatus) EnumDescriptor() ([]byte, []int) {
	return file_MonopolySocialEventStatus_proto_rawDescGZIP(), []int{0}
}

var File_MonopolySocialEventStatus_proto protoreflect.FileDescriptor

var file_MonopolySocialEventStatus_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x7a, 0x0a, 0x19, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x53, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25,
	0x0a, 0x21, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x53, 0x4f, 0x43, 0x49, 0x41,
	0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x36, 0x0a, 0x32, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c,
	0x59, 0x5f, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45,
	0x4c, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonopolySocialEventStatus_proto_rawDescOnce sync.Once
	file_MonopolySocialEventStatus_proto_rawDescData = file_MonopolySocialEventStatus_proto_rawDesc
)

func file_MonopolySocialEventStatus_proto_rawDescGZIP() []byte {
	file_MonopolySocialEventStatus_proto_rawDescOnce.Do(func() {
		file_MonopolySocialEventStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonopolySocialEventStatus_proto_rawDescData)
	})
	return file_MonopolySocialEventStatus_proto_rawDescData
}

var file_MonopolySocialEventStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MonopolySocialEventStatus_proto_goTypes = []interface{}{
	(MonopolySocialEventStatus)(0), // 0: MonopolySocialEventStatus
}
var file_MonopolySocialEventStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MonopolySocialEventStatus_proto_init() }
func file_MonopolySocialEventStatus_proto_init() {
	if File_MonopolySocialEventStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MonopolySocialEventStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonopolySocialEventStatus_proto_goTypes,
		DependencyIndexes: file_MonopolySocialEventStatus_proto_depIdxs,
		EnumInfos:         file_MonopolySocialEventStatus_proto_enumTypes,
	}.Build()
	File_MonopolySocialEventStatus_proto = out.File
	file_MonopolySocialEventStatus_proto_rawDesc = nil
	file_MonopolySocialEventStatus_proto_goTypes = nil
	file_MonopolySocialEventStatus_proto_depIdxs = nil
}

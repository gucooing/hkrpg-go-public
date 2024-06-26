// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DisplayChallengeInfo.proto

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

type DisplayChallengeInfo int32

const (
	DisplayChallengeInfo_BATTLE_RECORD_NONE      DisplayChallengeInfo = 0
	DisplayChallengeInfo_BATTLE_RECORD_CHALLENGE DisplayChallengeInfo = 1
	DisplayChallengeInfo_BATTLE_RECORD_ROGUE     DisplayChallengeInfo = 2
)

// Enum value maps for DisplayChallengeInfo.
var (
	DisplayChallengeInfo_name = map[int32]string{
		0: "BATTLE_RECORD_NONE",
		1: "BATTLE_RECORD_CHALLENGE",
		2: "BATTLE_RECORD_ROGUE",
	}
	DisplayChallengeInfo_value = map[string]int32{
		"BATTLE_RECORD_NONE":      0,
		"BATTLE_RECORD_CHALLENGE": 1,
		"BATTLE_RECORD_ROGUE":     2,
	}
)

func (x DisplayChallengeInfo) Enum() *DisplayChallengeInfo {
	p := new(DisplayChallengeInfo)
	*p = x
	return p
}

func (x DisplayChallengeInfo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DisplayChallengeInfo) Descriptor() protoreflect.EnumDescriptor {
	return file_DisplayChallengeInfo_proto_enumTypes[0].Descriptor()
}

func (DisplayChallengeInfo) Type() protoreflect.EnumType {
	return &file_DisplayChallengeInfo_proto_enumTypes[0]
}

func (x DisplayChallengeInfo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DisplayChallengeInfo.Descriptor instead.
func (DisplayChallengeInfo) EnumDescriptor() ([]byte, []int) {
	return file_DisplayChallengeInfo_proto_rawDescGZIP(), []int{0}
}

var File_DisplayChallengeInfo_proto protoreflect.FileDescriptor

var file_DisplayChallengeInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x64, 0x0a, 0x14,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52,
	0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17,
	0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x43, 0x48,
	0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x42, 0x41, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x10, 0x02, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DisplayChallengeInfo_proto_rawDescOnce sync.Once
	file_DisplayChallengeInfo_proto_rawDescData = file_DisplayChallengeInfo_proto_rawDesc
)

func file_DisplayChallengeInfo_proto_rawDescGZIP() []byte {
	file_DisplayChallengeInfo_proto_rawDescOnce.Do(func() {
		file_DisplayChallengeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_DisplayChallengeInfo_proto_rawDescData)
	})
	return file_DisplayChallengeInfo_proto_rawDescData
}

var file_DisplayChallengeInfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DisplayChallengeInfo_proto_goTypes = []interface{}{
	(DisplayChallengeInfo)(0), // 0: DisplayChallengeInfo
}
var file_DisplayChallengeInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DisplayChallengeInfo_proto_init() }
func file_DisplayChallengeInfo_proto_init() {
	if File_DisplayChallengeInfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DisplayChallengeInfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DisplayChallengeInfo_proto_goTypes,
		DependencyIndexes: file_DisplayChallengeInfo_proto_depIdxs,
		EnumInfos:         file_DisplayChallengeInfo_proto_enumTypes,
	}.Build()
	File_DisplayChallengeInfo_proto = out.File
	file_DisplayChallengeInfo_proto_rawDesc = nil
	file_DisplayChallengeInfo_proto_goTypes = nil
	file_DisplayChallengeInfo_proto_depIdxs = nil
}

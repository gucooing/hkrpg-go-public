// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AddRogueBuffScNotify.proto

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

type AddRogueBuffScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source       RogueBuffSource `protobuf:"varint,4,opt,name=source,proto3,enum=RogueBuffSource" json:"source,omitempty"`
	MazeBuffInfo *RogueBuff      `protobuf:"bytes,1,opt,name=maze_buff_info,json=mazeBuffInfo,proto3" json:"maze_buff_info,omitempty"`
}

func (x *AddRogueBuffScNotify) Reset() {
	*x = AddRogueBuffScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AddRogueBuffScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRogueBuffScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRogueBuffScNotify) ProtoMessage() {}

func (x *AddRogueBuffScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AddRogueBuffScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRogueBuffScNotify.ProtoReflect.Descriptor instead.
func (*AddRogueBuffScNotify) Descriptor() ([]byte, []int) {
	return file_AddRogueBuffScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AddRogueBuffScNotify) GetSource() RogueBuffSource {
	if x != nil {
		return x.Source
	}
	return RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_NONE
}

func (x *AddRogueBuffScNotify) GetMazeBuffInfo() *RogueBuff {
	if x != nil {
		return x.MazeBuffInfo
	}
	return nil
}

var File_AddRogueBuffScNotify_proto protoreflect.FileDescriptor

var file_AddRogueBuffScNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x63,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0e, 0x6d, 0x61, 0x7a, 0x65, 0x5f, 0x62,
	0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x0c, 0x6d, 0x61, 0x7a, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AddRogueBuffScNotify_proto_rawDescOnce sync.Once
	file_AddRogueBuffScNotify_proto_rawDescData = file_AddRogueBuffScNotify_proto_rawDesc
)

func file_AddRogueBuffScNotify_proto_rawDescGZIP() []byte {
	file_AddRogueBuffScNotify_proto_rawDescOnce.Do(func() {
		file_AddRogueBuffScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AddRogueBuffScNotify_proto_rawDescData)
	})
	return file_AddRogueBuffScNotify_proto_rawDescData
}

var file_AddRogueBuffScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AddRogueBuffScNotify_proto_goTypes = []interface{}{
	(*AddRogueBuffScNotify)(nil), // 0: AddRogueBuffScNotify
	(RogueBuffSource)(0),         // 1: RogueBuffSource
	(*RogueBuff)(nil),            // 2: RogueBuff
}
var file_AddRogueBuffScNotify_proto_depIdxs = []int32{
	1, // 0: AddRogueBuffScNotify.source:type_name -> RogueBuffSource
	2, // 1: AddRogueBuffScNotify.maze_buff_info:type_name -> RogueBuff
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AddRogueBuffScNotify_proto_init() }
func file_AddRogueBuffScNotify_proto_init() {
	if File_AddRogueBuffScNotify_proto != nil {
		return
	}
	file_RogueBuff_proto_init()
	file_RogueBuffSource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AddRogueBuffScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRogueBuffScNotify); i {
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
			RawDescriptor: file_AddRogueBuffScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AddRogueBuffScNotify_proto_goTypes,
		DependencyIndexes: file_AddRogueBuffScNotify_proto_depIdxs,
		MessageInfos:      file_AddRogueBuffScNotify_proto_msgTypes,
	}.Build()
	File_AddRogueBuffScNotify_proto = out.File
	file_AddRogueBuffScNotify_proto_rawDesc = nil
	file_AddRogueBuffScNotify_proto_goTypes = nil
	file_AddRogueBuffScNotify_proto_depIdxs = nil
}

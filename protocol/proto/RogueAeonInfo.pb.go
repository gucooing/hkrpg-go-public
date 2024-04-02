// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueAeonInfo.proto

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

type RogueAeonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsUnlocked    bool     `protobuf:"varint,5,opt,name=is_unlocked,json=isUnlocked,proto3" json:"is_unlocked,omitempty"`
	NFFJHLFKCPE   uint32   `protobuf:"varint,11,opt,name=NFFJHLFKCPE,proto3" json:"NFFJHLFKCPE,omitempty"`
	UnlockAeonNum uint32   `protobuf:"varint,15,opt,name=unlock_aeon_num,json=unlockAeonNum,proto3" json:"unlock_aeon_num,omitempty"`
	AeonIdList    []uint32 `protobuf:"varint,13,rep,packed,name=aeon_id_list,json=aeonIdList,proto3" json:"aeon_id_list,omitempty"`
}

func (x *RogueAeonInfo) Reset() {
	*x = RogueAeonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueAeonInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueAeonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueAeonInfo) ProtoMessage() {}

func (x *RogueAeonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueAeonInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueAeonInfo.ProtoReflect.Descriptor instead.
func (*RogueAeonInfo) Descriptor() ([]byte, []int) {
	return file_RogueAeonInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueAeonInfo) GetIsUnlocked() bool {
	if x != nil {
		return x.IsUnlocked
	}
	return false
}

func (x *RogueAeonInfo) GetNFFJHLFKCPE() uint32 {
	if x != nil {
		return x.NFFJHLFKCPE
	}
	return 0
}

func (x *RogueAeonInfo) GetUnlockAeonNum() uint32 {
	if x != nil {
		return x.UnlockAeonNum
	}
	return 0
}

func (x *RogueAeonInfo) GetAeonIdList() []uint32 {
	if x != nil {
		return x.AeonIdList
	}
	return nil
}

var File_RogueAeonInfo_proto protoreflect.FileDescriptor

var file_RogueAeonInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41,
	0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x46, 0x46, 0x4a,
	0x48, 0x4c, 0x46, 0x4b, 0x43, 0x50, 0x45, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e,
	0x46, 0x46, 0x4a, 0x48, 0x4c, 0x46, 0x4b, 0x43, 0x50, 0x45, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x65, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x65, 0x6f, 0x6e, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueAeonInfo_proto_rawDescOnce sync.Once
	file_RogueAeonInfo_proto_rawDescData = file_RogueAeonInfo_proto_rawDesc
)

func file_RogueAeonInfo_proto_rawDescGZIP() []byte {
	file_RogueAeonInfo_proto_rawDescOnce.Do(func() {
		file_RogueAeonInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueAeonInfo_proto_rawDescData)
	})
	return file_RogueAeonInfo_proto_rawDescData
}

var file_RogueAeonInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueAeonInfo_proto_goTypes = []interface{}{
	(*RogueAeonInfo)(nil), // 0: RogueAeonInfo
}
var file_RogueAeonInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueAeonInfo_proto_init() }
func file_RogueAeonInfo_proto_init() {
	if File_RogueAeonInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueAeonInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueAeonInfo); i {
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
			RawDescriptor: file_RogueAeonInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueAeonInfo_proto_goTypes,
		DependencyIndexes: file_RogueAeonInfo_proto_depIdxs,
		MessageInfos:      file_RogueAeonInfo_proto_msgTypes,
	}.Build()
	File_RogueAeonInfo_proto = out.File
	file_RogueAeonInfo_proto_rawDesc = nil
	file_RogueAeonInfo_proto_goTypes = nil
	file_RogueAeonInfo_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PropRogueInfo.proto

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

type PropRogueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId uint32 `protobuf:"varint,10,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	SiteId uint32 `protobuf:"varint,13,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
}

func (x *PropRogueInfo) Reset() {
	*x = PropRogueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PropRogueInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropRogueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropRogueInfo) ProtoMessage() {}

func (x *PropRogueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PropRogueInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropRogueInfo.ProtoReflect.Descriptor instead.
func (*PropRogueInfo) Descriptor() ([]byte, []int) {
	return file_PropRogueInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PropRogueInfo) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *PropRogueInfo) GetSiteId() uint32 {
	if x != nil {
		return x.SiteId
	}
	return 0
}

var File_PropRogueInfo_proto protoreflect.FileDescriptor

var file_PropRogueInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PropRogueInfo_proto_rawDescOnce sync.Once
	file_PropRogueInfo_proto_rawDescData = file_PropRogueInfo_proto_rawDesc
)

func file_PropRogueInfo_proto_rawDescGZIP() []byte {
	file_PropRogueInfo_proto_rawDescOnce.Do(func() {
		file_PropRogueInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PropRogueInfo_proto_rawDescData)
	})
	return file_PropRogueInfo_proto_rawDescData
}

var file_PropRogueInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PropRogueInfo_proto_goTypes = []interface{}{
	(*PropRogueInfo)(nil), // 0: PropRogueInfo
}
var file_PropRogueInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PropRogueInfo_proto_init() }
func file_PropRogueInfo_proto_init() {
	if File_PropRogueInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PropRogueInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropRogueInfo); i {
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
			RawDescriptor: file_PropRogueInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PropRogueInfo_proto_goTypes,
		DependencyIndexes: file_PropRogueInfo_proto_depIdxs,
		MessageInfos:      file_PropRogueInfo_proto_msgTypes,
	}.Build()
	File_PropRogueInfo_proto = out.File
	file_PropRogueInfo_proto_rawDesc = nil
	file_PropRogueInfo_proto_goTypes = nil
	file_PropRogueInfo_proto_depIdxs = nil
}

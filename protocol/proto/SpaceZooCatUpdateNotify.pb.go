// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SpaceZooCatUpdateNotify.proto

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

type SpaceZooCatUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LFOGBCGEDBM bool           `protobuf:"varint,5,opt,name=LFOGBCGEDBM,proto3" json:"LFOGBCGEDBM,omitempty"`
	OMDDHDFENML []*MJNBGFIHJBN `protobuf:"bytes,15,rep,name=OMDDHDFENML,proto3" json:"OMDDHDFENML,omitempty"`
	FMEGELGMOKJ bool           `protobuf:"varint,10,opt,name=FMEGELGMOKJ,proto3" json:"FMEGELGMOKJ,omitempty"`
}

func (x *SpaceZooCatUpdateNotify) Reset() {
	*x = SpaceZooCatUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SpaceZooCatUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceZooCatUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceZooCatUpdateNotify) ProtoMessage() {}

func (x *SpaceZooCatUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SpaceZooCatUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceZooCatUpdateNotify.ProtoReflect.Descriptor instead.
func (*SpaceZooCatUpdateNotify) Descriptor() ([]byte, []int) {
	return file_SpaceZooCatUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceZooCatUpdateNotify) GetLFOGBCGEDBM() bool {
	if x != nil {
		return x.LFOGBCGEDBM
	}
	return false
}

func (x *SpaceZooCatUpdateNotify) GetOMDDHDFENML() []*MJNBGFIHJBN {
	if x != nil {
		return x.OMDDHDFENML
	}
	return nil
}

func (x *SpaceZooCatUpdateNotify) GetFMEGELGMOKJ() bool {
	if x != nil {
		return x.FMEGELGMOKJ
	}
	return false
}

var File_SpaceZooCatUpdateNotify_proto protoreflect.FileDescriptor

var file_SpaceZooCatUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x70, 0x61, 0x63, 0x65, 0x5a, 0x6f, 0x6f, 0x43, 0x61, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4d, 0x4a, 0x4e, 0x42, 0x47, 0x46, 0x49, 0x48, 0x4a, 0x42, 0x4e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x53, 0x70, 0x61, 0x63, 0x65, 0x5a, 0x6f, 0x6f, 0x43,
	0x61, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x4c, 0x46, 0x4f, 0x47, 0x42, 0x43, 0x47, 0x45, 0x44, 0x42, 0x4d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x46, 0x4f, 0x47, 0x42, 0x43, 0x47, 0x45, 0x44, 0x42, 0x4d,
	0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x4d, 0x44, 0x44, 0x48, 0x44, 0x46, 0x45, 0x4e, 0x4d, 0x4c, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4a, 0x4e, 0x42, 0x47, 0x46, 0x49, 0x48,
	0x4a, 0x42, 0x4e, 0x52, 0x0b, 0x4f, 0x4d, 0x44, 0x44, 0x48, 0x44, 0x46, 0x45, 0x4e, 0x4d, 0x4c,
	0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4d, 0x45, 0x47, 0x45, 0x4c, 0x47, 0x4d, 0x4f, 0x4b, 0x4a, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x4d, 0x45, 0x47, 0x45, 0x4c, 0x47, 0x4d, 0x4f,
	0x4b, 0x4a, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SpaceZooCatUpdateNotify_proto_rawDescOnce sync.Once
	file_SpaceZooCatUpdateNotify_proto_rawDescData = file_SpaceZooCatUpdateNotify_proto_rawDesc
)

func file_SpaceZooCatUpdateNotify_proto_rawDescGZIP() []byte {
	file_SpaceZooCatUpdateNotify_proto_rawDescOnce.Do(func() {
		file_SpaceZooCatUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SpaceZooCatUpdateNotify_proto_rawDescData)
	})
	return file_SpaceZooCatUpdateNotify_proto_rawDescData
}

var file_SpaceZooCatUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SpaceZooCatUpdateNotify_proto_goTypes = []interface{}{
	(*SpaceZooCatUpdateNotify)(nil), // 0: SpaceZooCatUpdateNotify
	(*MJNBGFIHJBN)(nil),             // 1: MJNBGFIHJBN
}
var file_SpaceZooCatUpdateNotify_proto_depIdxs = []int32{
	1, // 0: SpaceZooCatUpdateNotify.OMDDHDFENML:type_name -> MJNBGFIHJBN
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SpaceZooCatUpdateNotify_proto_init() }
func file_SpaceZooCatUpdateNotify_proto_init() {
	if File_SpaceZooCatUpdateNotify_proto != nil {
		return
	}
	file_MJNBGFIHJBN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SpaceZooCatUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceZooCatUpdateNotify); i {
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
			RawDescriptor: file_SpaceZooCatUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SpaceZooCatUpdateNotify_proto_goTypes,
		DependencyIndexes: file_SpaceZooCatUpdateNotify_proto_depIdxs,
		MessageInfos:      file_SpaceZooCatUpdateNotify_proto_msgTypes,
	}.Build()
	File_SpaceZooCatUpdateNotify_proto = out.File
	file_SpaceZooCatUpdateNotify_proto_rawDesc = nil
	file_SpaceZooCatUpdateNotify_proto_goTypes = nil
	file_SpaceZooCatUpdateNotify_proto_depIdxs = nil
}

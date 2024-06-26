// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueHandbook.proto

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

type RogueHandbook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuffList    []*CMIEHNGEDEH `protobuf:"bytes,15,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	CGIFALDPBFE []*CAMHEKJONFO `protobuf:"bytes,9,rep,name=CGIFALDPBFE,proto3" json:"CGIFALDPBFE,omitempty"`
	AGGEDHEBPAH []*ABKHJLJEJBJ `protobuf:"bytes,8,rep,name=AGGEDHEBPAH,proto3" json:"AGGEDHEBPAH,omitempty"`
	MiracleList []*CCKEJCDOLCK `protobuf:"bytes,3,rep,name=miracle_list,json=miracleList,proto3" json:"miracle_list,omitempty"`
}

func (x *RogueHandbook) Reset() {
	*x = RogueHandbook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueHandbook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueHandbook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueHandbook) ProtoMessage() {}

func (x *RogueHandbook) ProtoReflect() protoreflect.Message {
	mi := &file_RogueHandbook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueHandbook.ProtoReflect.Descriptor instead.
func (*RogueHandbook) Descriptor() ([]byte, []int) {
	return file_RogueHandbook_proto_rawDescGZIP(), []int{0}
}

func (x *RogueHandbook) GetBuffList() []*CMIEHNGEDEH {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *RogueHandbook) GetCGIFALDPBFE() []*CAMHEKJONFO {
	if x != nil {
		return x.CGIFALDPBFE
	}
	return nil
}

func (x *RogueHandbook) GetAGGEDHEBPAH() []*ABKHJLJEJBJ {
	if x != nil {
		return x.AGGEDHEBPAH
	}
	return nil
}

func (x *RogueHandbook) GetMiracleList() []*CCKEJCDOLCK {
	if x != nil {
		return x.MiracleList
	}
	return nil
}

var File_RogueHandbook_proto protoreflect.FileDescriptor

var file_RogueHandbook_proto_rawDesc = []byte{
	0x0a, 0x13, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x41, 0x4d, 0x48, 0x45, 0x4b, 0x4a, 0x4f, 0x4e,
	0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x43, 0x4b, 0x45, 0x4a, 0x43,
	0x44, 0x4f, 0x4c, 0x43, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x42, 0x4b,
	0x48, 0x4a, 0x4c, 0x4a, 0x45, 0x4a, 0x42, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x43, 0x4d, 0x49, 0x45, 0x48, 0x4e, 0x47, 0x45, 0x44, 0x45, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x6f, 0x6b, 0x12, 0x29, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4d, 0x49, 0x45, 0x48, 0x4e, 0x47,
	0x45, 0x44, 0x45, 0x48, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x0b, 0x43, 0x47, 0x49, 0x46, 0x41, 0x4c, 0x44, 0x50, 0x42, 0x46, 0x45, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x41, 0x4d, 0x48, 0x45, 0x4b, 0x4a, 0x4f, 0x4e, 0x46,
	0x4f, 0x52, 0x0b, 0x43, 0x47, 0x49, 0x46, 0x41, 0x4c, 0x44, 0x50, 0x42, 0x46, 0x45, 0x12, 0x2e,
	0x0a, 0x0b, 0x41, 0x47, 0x47, 0x45, 0x44, 0x48, 0x45, 0x42, 0x50, 0x41, 0x48, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x42, 0x4b, 0x48, 0x4a, 0x4c, 0x4a, 0x45, 0x4a, 0x42,
	0x4a, 0x52, 0x0b, 0x41, 0x47, 0x47, 0x45, 0x44, 0x48, 0x45, 0x42, 0x50, 0x41, 0x48, 0x12, 0x2f,
	0x0a, 0x0c, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x43, 0x4b, 0x45, 0x4a, 0x43, 0x44, 0x4f, 0x4c,
	0x43, 0x4b, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RogueHandbook_proto_rawDescOnce sync.Once
	file_RogueHandbook_proto_rawDescData = file_RogueHandbook_proto_rawDesc
)

func file_RogueHandbook_proto_rawDescGZIP() []byte {
	file_RogueHandbook_proto_rawDescOnce.Do(func() {
		file_RogueHandbook_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueHandbook_proto_rawDescData)
	})
	return file_RogueHandbook_proto_rawDescData
}

var file_RogueHandbook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueHandbook_proto_goTypes = []interface{}{
	(*RogueHandbook)(nil), // 0: RogueHandbook
	(*CMIEHNGEDEH)(nil),   // 1: CMIEHNGEDEH
	(*CAMHEKJONFO)(nil),   // 2: CAMHEKJONFO
	(*ABKHJLJEJBJ)(nil),   // 3: ABKHJLJEJBJ
	(*CCKEJCDOLCK)(nil),   // 4: CCKEJCDOLCK
}
var file_RogueHandbook_proto_depIdxs = []int32{
	1, // 0: RogueHandbook.buff_list:type_name -> CMIEHNGEDEH
	2, // 1: RogueHandbook.CGIFALDPBFE:type_name -> CAMHEKJONFO
	3, // 2: RogueHandbook.AGGEDHEBPAH:type_name -> ABKHJLJEJBJ
	4, // 3: RogueHandbook.miracle_list:type_name -> CCKEJCDOLCK
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_RogueHandbook_proto_init() }
func file_RogueHandbook_proto_init() {
	if File_RogueHandbook_proto != nil {
		return
	}
	file_CAMHEKJONFO_proto_init()
	file_CCKEJCDOLCK_proto_init()
	file_ABKHJLJEJBJ_proto_init()
	file_CMIEHNGEDEH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueHandbook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueHandbook); i {
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
			RawDescriptor: file_RogueHandbook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueHandbook_proto_goTypes,
		DependencyIndexes: file_RogueHandbook_proto_depIdxs,
		MessageInfos:      file_RogueHandbook_proto_msgTypes,
	}.Build()
	File_RogueHandbook_proto = out.File
	file_RogueHandbook_proto_rawDesc = nil
	file_RogueHandbook_proto_goTypes = nil
	file_RogueHandbook_proto_depIdxs = nil
}

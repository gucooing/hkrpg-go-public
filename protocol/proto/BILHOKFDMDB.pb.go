// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BILHOKFDMDB.proto

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

type BILHOKFDMDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HIANPODFGON      uint32             `protobuf:"varint,13,opt,name=HIANPODFGON,proto3" json:"HIANPODFGON,omitempty"`
	PPMPGOGOLIF      []uint32           `protobuf:"varint,11,rep,packed,name=PPMPGOGOLIF,proto3" json:"PPMPGOGOLIF,omitempty"`
	BaseAvatarIdList []uint32           `protobuf:"varint,2,rep,packed,name=base_avatar_id_list,json=baseAvatarIdList,proto3" json:"base_avatar_id_list,omitempty"`
	Status           RogueStatus        `protobuf:"varint,12,opt,name=status,proto3,enum=RogueStatus" json:"status,omitempty"`
	CurReachRoomNum  uint32             `protobuf:"varint,7,opt,name=cur_reach_room_num,json=curReachRoomNum,proto3" json:"cur_reach_room_num,omitempty"`
	BuffInfo         *ChessRogueBuff    `protobuf:"bytes,4,opt,name=buff_info,json=buffInfo,proto3" json:"buff_info,omitempty"`
	MiracleInfo      *ChessRogueMiracle `protobuf:"bytes,9,opt,name=miracle_info,json=miracleInfo,proto3" json:"miracle_info,omitempty"`
	MapId            uint32             `protobuf:"varint,6,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	IADMGHFBJMI      uint32             `protobuf:"varint,8,opt,name=IADMGHFBJMI,proto3" json:"IADMGHFBJMI,omitempty"`
}

func (x *BILHOKFDMDB) Reset() {
	*x = BILHOKFDMDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BILHOKFDMDB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BILHOKFDMDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BILHOKFDMDB) ProtoMessage() {}

func (x *BILHOKFDMDB) ProtoReflect() protoreflect.Message {
	mi := &file_BILHOKFDMDB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BILHOKFDMDB.ProtoReflect.Descriptor instead.
func (*BILHOKFDMDB) Descriptor() ([]byte, []int) {
	return file_BILHOKFDMDB_proto_rawDescGZIP(), []int{0}
}

func (x *BILHOKFDMDB) GetHIANPODFGON() uint32 {
	if x != nil {
		return x.HIANPODFGON
	}
	return 0
}

func (x *BILHOKFDMDB) GetPPMPGOGOLIF() []uint32 {
	if x != nil {
		return x.PPMPGOGOLIF
	}
	return nil
}

func (x *BILHOKFDMDB) GetBaseAvatarIdList() []uint32 {
	if x != nil {
		return x.BaseAvatarIdList
	}
	return nil
}

func (x *BILHOKFDMDB) GetStatus() RogueStatus {
	if x != nil {
		return x.Status
	}
	return RogueStatus_ROGUE_STATUS_NONE
}

func (x *BILHOKFDMDB) GetCurReachRoomNum() uint32 {
	if x != nil {
		return x.CurReachRoomNum
	}
	return 0
}

func (x *BILHOKFDMDB) GetBuffInfo() *ChessRogueBuff {
	if x != nil {
		return x.BuffInfo
	}
	return nil
}

func (x *BILHOKFDMDB) GetMiracleInfo() *ChessRogueMiracle {
	if x != nil {
		return x.MiracleInfo
	}
	return nil
}

func (x *BILHOKFDMDB) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *BILHOKFDMDB) GetIADMGHFBJMI() uint32 {
	if x != nil {
		return x.IADMGHFBJMI
	}
	return 0
}

var File_BILHOKFDMDB_proto protoreflect.FileDescriptor

var file_BILHOKFDMDB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x49, 0x4c, 0x48, 0x4f, 0x4b, 0x46, 0x44, 0x4d, 0x44, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x02, 0x0a, 0x0b, 0x42, 0x49, 0x4c, 0x48, 0x4f, 0x4b,
	0x46, 0x44, 0x4d, 0x44, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x49, 0x41, 0x4e, 0x50, 0x4f, 0x44,
	0x46, 0x47, 0x4f, 0x4e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x49, 0x41, 0x4e,
	0x50, 0x4f, 0x44, 0x46, 0x47, 0x4f, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x50, 0x4d, 0x50, 0x47,
	0x4f, 0x47, 0x4f, 0x4c, 0x49, 0x46, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x50,
	0x4d, 0x50, 0x47, 0x4f, 0x47, 0x4f, 0x4c, 0x49, 0x46, 0x12, 0x2d, 0x0a, 0x13, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b,
	0x0a, 0x12, 0x63, 0x75, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x52,
	0x65, 0x61, 0x63, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x09, 0x62,
	0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52,
	0x08, 0x62, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x0c, 0x6d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x41, 0x44, 0x4d, 0x47,
	0x48, 0x46, 0x42, 0x4a, 0x4d, 0x49, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x41,
	0x44, 0x4d, 0x47, 0x48, 0x46, 0x42, 0x4a, 0x4d, 0x49, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BILHOKFDMDB_proto_rawDescOnce sync.Once
	file_BILHOKFDMDB_proto_rawDescData = file_BILHOKFDMDB_proto_rawDesc
)

func file_BILHOKFDMDB_proto_rawDescGZIP() []byte {
	file_BILHOKFDMDB_proto_rawDescOnce.Do(func() {
		file_BILHOKFDMDB_proto_rawDescData = protoimpl.X.CompressGZIP(file_BILHOKFDMDB_proto_rawDescData)
	})
	return file_BILHOKFDMDB_proto_rawDescData
}

var file_BILHOKFDMDB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BILHOKFDMDB_proto_goTypes = []interface{}{
	(*BILHOKFDMDB)(nil),       // 0: BILHOKFDMDB
	(RogueStatus)(0),          // 1: RogueStatus
	(*ChessRogueBuff)(nil),    // 2: ChessRogueBuff
	(*ChessRogueMiracle)(nil), // 3: ChessRogueMiracle
}
var file_BILHOKFDMDB_proto_depIdxs = []int32{
	1, // 0: BILHOKFDMDB.status:type_name -> RogueStatus
	2, // 1: BILHOKFDMDB.buff_info:type_name -> ChessRogueBuff
	3, // 2: BILHOKFDMDB.miracle_info:type_name -> ChessRogueMiracle
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_BILHOKFDMDB_proto_init() }
func file_BILHOKFDMDB_proto_init() {
	if File_BILHOKFDMDB_proto != nil {
		return
	}
	file_RogueStatus_proto_init()
	file_ChessRogueBuff_proto_init()
	file_ChessRogueMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BILHOKFDMDB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BILHOKFDMDB); i {
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
			RawDescriptor: file_BILHOKFDMDB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BILHOKFDMDB_proto_goTypes,
		DependencyIndexes: file_BILHOKFDMDB_proto_depIdxs,
		MessageInfos:      file_BILHOKFDMDB_proto_msgTypes,
	}.Build()
	File_BILHOKFDMDB_proto = out.File
	file_BILHOKFDMDB_proto_rawDesc = nil
	file_BILHOKFDMDB_proto_goTypes = nil
	file_BILHOKFDMDB_proto_depIdxs = nil
}

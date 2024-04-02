// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneInfo.proto

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

type SceneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtraData          map[string]int32        `protobuf:"bytes,81,rep,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	EntityGroupList    []*SceneEntityGroupInfo `protobuf:"bytes,1649,rep,name=entity_group_list,json=entityGroupList,proto3" json:"entity_group_list,omitempty"`
	ClientPosVersion   uint32                  `protobuf:"varint,7,opt,name=client_pos_version,json=clientPosVersion,proto3" json:"client_pos_version,omitempty"`
	GroupIdList        []uint32                `protobuf:"varint,13,rep,packed,name=group_id_list,json=groupIdList,proto3" json:"group_id_list,omitempty"`
	PlaneId            uint32                  `protobuf:"varint,8,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
	WorldId            uint32                  `protobuf:"varint,11,opt,name=world_id,json=worldId,proto3" json:"world_id,omitempty"` // ?
	LightenSectionList []uint32                `protobuf:"varint,3,rep,packed,name=lighten_section_list,json=lightenSectionList,proto3" json:"lighten_section_list,omitempty"`
	LeaderEntityId     uint32                  `protobuf:"varint,14,opt,name=leader_entity_id,json=leaderEntityId,proto3" json:"leader_entity_id,omitempty"` // ?
	FloorId            uint32                  `protobuf:"varint,10,opt,name=floor_id,json=floorId,proto3" json:"floor_id,omitempty"`
	EntryId            uint32                  `protobuf:"varint,5,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	GameModeType       uint32                  `protobuf:"varint,2,opt,name=game_mode_type,json=gameModeType,proto3" json:"game_mode_type,omitempty"`
	EntityList         []*SceneEntityInfo      `protobuf:"bytes,6,rep,name=entity_list,json=entityList,proto3" json:"entity_list,omitempty"`
	GroupStateList     []*SceneGroupState      `protobuf:"bytes,1244,rep,name=group_state_list,json=groupStateList,proto3" json:"group_state_list,omitempty"`
}

func (x *SceneInfo) Reset() {
	*x = SceneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneInfo) ProtoMessage() {}

func (x *SceneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneInfo.ProtoReflect.Descriptor instead.
func (*SceneInfo) Descriptor() ([]byte, []int) {
	return file_SceneInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneInfo) GetExtraData() map[string]int32 {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

func (x *SceneInfo) GetEntityGroupList() []*SceneEntityGroupInfo {
	if x != nil {
		return x.EntityGroupList
	}
	return nil
}

func (x *SceneInfo) GetClientPosVersion() uint32 {
	if x != nil {
		return x.ClientPosVersion
	}
	return 0
}

func (x *SceneInfo) GetGroupIdList() []uint32 {
	if x != nil {
		return x.GroupIdList
	}
	return nil
}

func (x *SceneInfo) GetPlaneId() uint32 {
	if x != nil {
		return x.PlaneId
	}
	return 0
}

func (x *SceneInfo) GetWorldId() uint32 {
	if x != nil {
		return x.WorldId
	}
	return 0
}

func (x *SceneInfo) GetLightenSectionList() []uint32 {
	if x != nil {
		return x.LightenSectionList
	}
	return nil
}

func (x *SceneInfo) GetLeaderEntityId() uint32 {
	if x != nil {
		return x.LeaderEntityId
	}
	return 0
}

func (x *SceneInfo) GetFloorId() uint32 {
	if x != nil {
		return x.FloorId
	}
	return 0
}

func (x *SceneInfo) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *SceneInfo) GetGameModeType() uint32 {
	if x != nil {
		return x.GameModeType
	}
	return 0
}

func (x *SceneInfo) GetEntityList() []*SceneEntityInfo {
	if x != nil {
		return x.EntityList
	}
	return nil
}

func (x *SceneInfo) GetGroupStateList() []*SceneGroupState {
	if x != nil {
		return x.GroupStateList
	}
	return nil
}

var File_SceneInfo_proto protoreflect.FileDescriptor

var file_SceneInfo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x04, 0x0a, 0x09,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x51, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xf1, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x14, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66,
	0x6c, 0x6f, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x10, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xdc,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneInfo_proto_rawDescOnce sync.Once
	file_SceneInfo_proto_rawDescData = file_SceneInfo_proto_rawDesc
)

func file_SceneInfo_proto_rawDescGZIP() []byte {
	file_SceneInfo_proto_rawDescOnce.Do(func() {
		file_SceneInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneInfo_proto_rawDescData)
	})
	return file_SceneInfo_proto_rawDescData
}

var file_SceneInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SceneInfo_proto_goTypes = []interface{}{
	(*SceneInfo)(nil),            // 0: SceneInfo
	nil,                          // 1: SceneInfo.ExtraDataEntry
	(*SceneEntityGroupInfo)(nil), // 2: SceneEntityGroupInfo
	(*SceneEntityInfo)(nil),      // 3: SceneEntityInfo
	(*SceneGroupState)(nil),      // 4: SceneGroupState
}
var file_SceneInfo_proto_depIdxs = []int32{
	1, // 0: SceneInfo.extra_data:type_name -> SceneInfo.ExtraDataEntry
	2, // 1: SceneInfo.entity_group_list:type_name -> SceneEntityGroupInfo
	3, // 2: SceneInfo.entity_list:type_name -> SceneEntityInfo
	4, // 3: SceneInfo.group_state_list:type_name -> SceneGroupState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_SceneInfo_proto_init() }
func file_SceneInfo_proto_init() {
	if File_SceneInfo_proto != nil {
		return
	}
	file_SceneEntityInfo_proto_init()
	file_SceneGroupState_proto_init()
	file_SceneEntityGroupInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneInfo); i {
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
			RawDescriptor: file_SceneInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneInfo_proto_goTypes,
		DependencyIndexes: file_SceneInfo_proto_depIdxs,
		MessageInfos:      file_SceneInfo_proto_msgTypes,
	}.Build()
	File_SceneInfo_proto = out.File
	file_SceneInfo_proto_rawDesc = nil
	file_SceneInfo_proto_goTypes = nil
	file_SceneInfo_proto_depIdxs = nil
}

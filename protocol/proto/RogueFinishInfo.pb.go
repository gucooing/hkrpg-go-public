// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueFinishInfo.proto

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

type RogueFinishInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishedRoomCount uint32                 `protobuf:"varint,6,opt,name=finished_room_count,json=finishedRoomCount,proto3" json:"finished_room_count,omitempty"`
	TotalScore        uint32                 `protobuf:"varint,3,opt,name=total_score,json=totalScore,proto3" json:"total_score,omitempty"`
	PIENECGLNPM       bool                   `protobuf:"varint,8,opt,name=PIENECGLNPM,proto3" json:"PIENECGLNPM,omitempty"`
	IMDACMGEFCL       *ILKNLOHCBJA           `protobuf:"bytes,1,opt,name=IMDACMGEFCL,proto3" json:"IMDACMGEFCL,omitempty"`
	ScoreInfo         *RogueExploreScoreInfo `protobuf:"bytes,12,opt,name=score_info,json=scoreInfo,proto3" json:"score_info,omitempty"`
	Record            *RogueRecordInfo       `protobuf:"bytes,5,opt,name=record,proto3" json:"record,omitempty"`
	BDIBIICAHAC       uint32                 `protobuf:"varint,11,opt,name=BDIBIICAHAC,proto3" json:"BDIBIICAHAC,omitempty"`
	ScoreId           uint32                 `protobuf:"varint,9,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
	BIBMJCEMCFA       *ItemList              `protobuf:"bytes,2,opt,name=BIBMJCEMCFA,proto3" json:"BIBMJCEMCFA,omitempty"`
	AreaId            uint32                 `protobuf:"varint,1131,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	KPLGPNDGEAE       uint32                 `protobuf:"varint,13,opt,name=KPLGPNDGEAE,proto3" json:"KPLGPNDGEAE,omitempty"`
	ReachedRoomCount  uint32                 `protobuf:"varint,1911,opt,name=reached_room_count,json=reachedRoomCount,proto3" json:"reached_room_count,omitempty"`
	PrevRewardInfo    *RogueScoreRewardInfo  `protobuf:"bytes,14,opt,name=prev_reward_info,json=prevRewardInfo,proto3" json:"prev_reward_info,omitempty"`
	NextRewardInfo    *RogueScoreRewardInfo  `protobuf:"bytes,7,opt,name=next_reward_info,json=nextRewardInfo,proto3" json:"next_reward_info,omitempty"`
	PCMFBEEEMBM       *ItemList              `protobuf:"bytes,10,opt,name=PCMFBEEEMBM,proto3" json:"PCMFBEEEMBM,omitempty"`
	IsWin             bool                   `protobuf:"varint,4,opt,name=is_win,json=isWin,proto3" json:"is_win,omitempty"`
	BHLOCDABFLL       *ItemList              `protobuf:"bytes,15,opt,name=BHLOCDABFLL,proto3" json:"BHLOCDABFLL,omitempty"`
}

func (x *RogueFinishInfo) Reset() {
	*x = RogueFinishInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueFinishInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueFinishInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueFinishInfo) ProtoMessage() {}

func (x *RogueFinishInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueFinishInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueFinishInfo.ProtoReflect.Descriptor instead.
func (*RogueFinishInfo) Descriptor() ([]byte, []int) {
	return file_RogueFinishInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueFinishInfo) GetFinishedRoomCount() uint32 {
	if x != nil {
		return x.FinishedRoomCount
	}
	return 0
}

func (x *RogueFinishInfo) GetTotalScore() uint32 {
	if x != nil {
		return x.TotalScore
	}
	return 0
}

func (x *RogueFinishInfo) GetPIENECGLNPM() bool {
	if x != nil {
		return x.PIENECGLNPM
	}
	return false
}

func (x *RogueFinishInfo) GetIMDACMGEFCL() *ILKNLOHCBJA {
	if x != nil {
		return x.IMDACMGEFCL
	}
	return nil
}

func (x *RogueFinishInfo) GetScoreInfo() *RogueExploreScoreInfo {
	if x != nil {
		return x.ScoreInfo
	}
	return nil
}

func (x *RogueFinishInfo) GetRecord() *RogueRecordInfo {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *RogueFinishInfo) GetBDIBIICAHAC() uint32 {
	if x != nil {
		return x.BDIBIICAHAC
	}
	return 0
}

func (x *RogueFinishInfo) GetScoreId() uint32 {
	if x != nil {
		return x.ScoreId
	}
	return 0
}

func (x *RogueFinishInfo) GetBIBMJCEMCFA() *ItemList {
	if x != nil {
		return x.BIBMJCEMCFA
	}
	return nil
}

func (x *RogueFinishInfo) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *RogueFinishInfo) GetKPLGPNDGEAE() uint32 {
	if x != nil {
		return x.KPLGPNDGEAE
	}
	return 0
}

func (x *RogueFinishInfo) GetReachedRoomCount() uint32 {
	if x != nil {
		return x.ReachedRoomCount
	}
	return 0
}

func (x *RogueFinishInfo) GetPrevRewardInfo() *RogueScoreRewardInfo {
	if x != nil {
		return x.PrevRewardInfo
	}
	return nil
}

func (x *RogueFinishInfo) GetNextRewardInfo() *RogueScoreRewardInfo {
	if x != nil {
		return x.NextRewardInfo
	}
	return nil
}

func (x *RogueFinishInfo) GetPCMFBEEEMBM() *ItemList {
	if x != nil {
		return x.PCMFBEEEMBM
	}
	return nil
}

func (x *RogueFinishInfo) GetIsWin() bool {
	if x != nil {
		return x.IsWin
	}
	return false
}

func (x *RogueFinishInfo) GetBHLOCDABFLL() *ItemList {
	if x != nil {
		return x.BHLOCDABFLL
	}
	return nil
}

var File_RogueFinishInfo_proto protoreflect.FileDescriptor

var file_RogueFinishInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4c, 0x4b, 0x4e, 0x4c, 0x4f, 0x48, 0x43, 0x42, 0x4a, 0x41,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x05, 0x0a, 0x0f, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e,
	0x0a, 0x13, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x49, 0x45, 0x4e, 0x45, 0x43, 0x47, 0x4c, 0x4e, 0x50, 0x4d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x49, 0x45, 0x4e, 0x45, 0x43, 0x47, 0x4c, 0x4e, 0x50,
	0x4d, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x4d, 0x44, 0x41, 0x43, 0x4d, 0x47, 0x45, 0x46, 0x43, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x4c, 0x4b, 0x4e, 0x4c, 0x4f, 0x48,
	0x43, 0x42, 0x4a, 0x41, 0x52, 0x0b, 0x49, 0x4d, 0x44, 0x41, 0x43, 0x4d, 0x47, 0x45, 0x46, 0x43,
	0x4c, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x44, 0x49, 0x42, 0x49, 0x49, 0x43, 0x41, 0x48, 0x41,
	0x43, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x44, 0x49, 0x42, 0x49, 0x49, 0x43,
	0x41, 0x48, 0x41, 0x43, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x0b, 0x42, 0x49, 0x42, 0x4d, 0x4a, 0x43, 0x45, 0x4d, 0x43, 0x46, 0x41, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x0b, 0x42, 0x49, 0x42, 0x4d, 0x4a, 0x43, 0x45, 0x4d, 0x43, 0x46, 0x41, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0xeb, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x50, 0x4c, 0x47, 0x50, 0x4e,
	0x44, 0x47, 0x45, 0x41, 0x45, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x50, 0x4c,
	0x47, 0x50, 0x4e, 0x44, 0x47, 0x45, 0x41, 0x45, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xf7,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x6f,
	0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x5f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x0b, 0x50, 0x43, 0x4d,
	0x46, 0x42, 0x45, 0x45, 0x45, 0x4d, 0x42, 0x4d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x50, 0x43, 0x4d, 0x46, 0x42,
	0x45, 0x45, 0x45, 0x4d, 0x42, 0x4d, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x57, 0x69, 0x6e, 0x12, 0x2b, 0x0a,
	0x0b, 0x42, 0x48, 0x4c, 0x4f, 0x43, 0x44, 0x41, 0x42, 0x46, 0x4c, 0x4c, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x42,
	0x48, 0x4c, 0x4f, 0x43, 0x44, 0x41, 0x42, 0x46, 0x4c, 0x4c, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueFinishInfo_proto_rawDescOnce sync.Once
	file_RogueFinishInfo_proto_rawDescData = file_RogueFinishInfo_proto_rawDesc
)

func file_RogueFinishInfo_proto_rawDescGZIP() []byte {
	file_RogueFinishInfo_proto_rawDescOnce.Do(func() {
		file_RogueFinishInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueFinishInfo_proto_rawDescData)
	})
	return file_RogueFinishInfo_proto_rawDescData
}

var file_RogueFinishInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueFinishInfo_proto_goTypes = []interface{}{
	(*RogueFinishInfo)(nil),       // 0: RogueFinishInfo
	(*ILKNLOHCBJA)(nil),           // 1: ILKNLOHCBJA
	(*RogueExploreScoreInfo)(nil), // 2: RogueExploreScoreInfo
	(*RogueRecordInfo)(nil),       // 3: RogueRecordInfo
	(*ItemList)(nil),              // 4: ItemList
	(*RogueScoreRewardInfo)(nil),  // 5: RogueScoreRewardInfo
}
var file_RogueFinishInfo_proto_depIdxs = []int32{
	1, // 0: RogueFinishInfo.IMDACMGEFCL:type_name -> ILKNLOHCBJA
	2, // 1: RogueFinishInfo.score_info:type_name -> RogueExploreScoreInfo
	3, // 2: RogueFinishInfo.record:type_name -> RogueRecordInfo
	4, // 3: RogueFinishInfo.BIBMJCEMCFA:type_name -> ItemList
	5, // 4: RogueFinishInfo.prev_reward_info:type_name -> RogueScoreRewardInfo
	5, // 5: RogueFinishInfo.next_reward_info:type_name -> RogueScoreRewardInfo
	4, // 6: RogueFinishInfo.PCMFBEEEMBM:type_name -> ItemList
	4, // 7: RogueFinishInfo.BHLOCDABFLL:type_name -> ItemList
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_RogueFinishInfo_proto_init() }
func file_RogueFinishInfo_proto_init() {
	if File_RogueFinishInfo_proto != nil {
		return
	}
	file_RogueScoreRewardInfo_proto_init()
	file_ILKNLOHCBJA_proto_init()
	file_RogueRecordInfo_proto_init()
	file_RogueExploreScoreInfo_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueFinishInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueFinishInfo); i {
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
			RawDescriptor: file_RogueFinishInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueFinishInfo_proto_goTypes,
		DependencyIndexes: file_RogueFinishInfo_proto_depIdxs,
		MessageInfos:      file_RogueFinishInfo_proto_msgTypes,
	}.Build()
	File_RogueFinishInfo_proto = out.File
	file_RogueFinishInfo_proto_rawDesc = nil
	file_RogueFinishInfo_proto_goTypes = nil
	file_RogueFinishInfo_proto_depIdxs = nil
}

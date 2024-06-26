// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HandleRogueCommonPendingActionCsReq.proto

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

type HandleRogueCommonPendingActionCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueLocation uint32 `protobuf:"varint,6,opt,name=queue_location,json=queueLocation,proto3" json:"queue_location,omitempty"`
	// Types that are assignable to Action:
	//
	//	*HandleRogueCommonPendingActionCsReq_BuffSelectResult
	//	*HandleRogueCommonPendingActionCsReq_EKEBAOIEGPC
	//	*HandleRogueCommonPendingActionCsReq_FKFPCIODHNE
	//	*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult
	//	*HandleRogueCommonPendingActionCsReq_MiracleSelectResult
	//	*HandleRogueCommonPendingActionCsReq_MNNFFEBLNBH
	//	*HandleRogueCommonPendingActionCsReq_DJJFNHCFLJN
	//	*HandleRogueCommonPendingActionCsReq_GEHPONOPIMG
	//	*HandleRogueCommonPendingActionCsReq_GOBLCCNDJNO
	//	*HandleRogueCommonPendingActionCsReq_BonusSelectResult
	Action isHandleRogueCommonPendingActionCsReq_Action `protobuf_oneof:"action"`
}

func (x *HandleRogueCommonPendingActionCsReq) Reset() {
	*x = HandleRogueCommonPendingActionCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleRogueCommonPendingActionCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleRogueCommonPendingActionCsReq) ProtoMessage() {}

func (x *HandleRogueCommonPendingActionCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleRogueCommonPendingActionCsReq.ProtoReflect.Descriptor instead.
func (*HandleRogueCommonPendingActionCsReq) Descriptor() ([]byte, []int) {
	return file_HandleRogueCommonPendingActionCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *HandleRogueCommonPendingActionCsReq) GetQueueLocation() uint32 {
	if x != nil {
		return x.QueueLocation
	}
	return 0
}

func (m *HandleRogueCommonPendingActionCsReq) GetAction() isHandleRogueCommonPendingActionCsReq_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBuffSelectResult() *RogueBuffSelectResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_BuffSelectResult); ok {
		return x.BuffSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetEKEBAOIEGPC() *IIDPIILOGJI {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_EKEBAOIEGPC); ok {
		return x.EKEBAOIEGPC
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetFKFPCIODHNE() *KMPOJOBOKOP {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_FKFPCIODHNE); ok {
		return x.FKFPCIODHNE
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBuffRerollSelectResult() *RogueBuffRerollResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult); ok {
		return x.BuffRerollSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetMiracleSelectResult() *RogueMiracleSelectResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult); ok {
		return x.MiracleSelectResult
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetMNNFFEBLNBH() *DKCPGELJMGA {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_MNNFFEBLNBH); ok {
		return x.MNNFFEBLNBH
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetDJJFNHCFLJN() *KMLFFKILCHI {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_DJJFNHCFLJN); ok {
		return x.DJJFNHCFLJN
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetGEHPONOPIMG() *MCAEMBMJGMJ {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_GEHPONOPIMG); ok {
		return x.GEHPONOPIMG
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetGOBLCCNDJNO() *GFLCNDDIOLM {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_GOBLCCNDJNO); ok {
		return x.GOBLCCNDJNO
	}
	return nil
}

func (x *HandleRogueCommonPendingActionCsReq) GetBonusSelectResult() *RogueBonusSelectResult {
	if x, ok := x.GetAction().(*HandleRogueCommonPendingActionCsReq_BonusSelectResult); ok {
		return x.BonusSelectResult
	}
	return nil
}

type isHandleRogueCommonPendingActionCsReq_Action interface {
	isHandleRogueCommonPendingActionCsReq_Action()
}

type HandleRogueCommonPendingActionCsReq_BuffSelectResult struct {
	BuffSelectResult *RogueBuffSelectResult `protobuf:"bytes,1586,opt,name=buff_select_result,json=buffSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_EKEBAOIEGPC struct {
	EKEBAOIEGPC *IIDPIILOGJI `protobuf:"bytes,537,opt,name=EKEBAOIEGPC,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_FKFPCIODHNE struct {
	FKFPCIODHNE *KMPOJOBOKOP `protobuf:"bytes,295,opt,name=FKFPCIODHNE,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult struct {
	BuffRerollSelectResult *RogueBuffRerollResult `protobuf:"bytes,308,opt,name=buff_reroll_select_result,json=buffRerollSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_MiracleSelectResult struct {
	MiracleSelectResult *RogueMiracleSelectResult `protobuf:"bytes,1209,opt,name=miracle_select_result,json=miracleSelectResult,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_MNNFFEBLNBH struct {
	MNNFFEBLNBH *DKCPGELJMGA `protobuf:"bytes,1911,opt,name=MNNFFEBLNBH,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_DJJFNHCFLJN struct {
	DJJFNHCFLJN *KMLFFKILCHI `protobuf:"bytes,1338,opt,name=DJJFNHCFLJN,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_GEHPONOPIMG struct {
	GEHPONOPIMG *MCAEMBMJGMJ `protobuf:"bytes,481,opt,name=GEHPONOPIMG,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_GOBLCCNDJNO struct {
	GOBLCCNDJNO *GFLCNDDIOLM `protobuf:"bytes,1632,opt,name=GOBLCCNDJNO,proto3,oneof"`
}

type HandleRogueCommonPendingActionCsReq_BonusSelectResult struct {
	BonusSelectResult *RogueBonusSelectResult `protobuf:"bytes,1156,opt,name=bonus_select_result,json=bonusSelectResult,proto3,oneof"`
}

func (*HandleRogueCommonPendingActionCsReq_BuffSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_EKEBAOIEGPC) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_FKFPCIODHNE) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_MiracleSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_MNNFFEBLNBH) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_DJJFNHCFLJN) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_GEHPONOPIMG) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_GOBLCCNDJNO) isHandleRogueCommonPendingActionCsReq_Action() {
}

func (*HandleRogueCommonPendingActionCsReq_BonusSelectResult) isHandleRogueCommonPendingActionCsReq_Action() {
}

var File_HandleRogueCommonPendingActionCsReq_proto protoreflect.FileDescriptor

var file_HandleRogueCommonPendingActionCsReq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4d, 0x50, 0x4f, 0x4a, 0x4f, 0x42, 0x4f,
	0x4b, 0x4f, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4b, 0x43, 0x50, 0x47,
	0x45, 0x4c, 0x4a, 0x4d, 0x47, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x43, 0x41, 0x45, 0x4d,
	0x42, 0x4d, 0x4a, 0x47, 0x4d, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4d,
	0x4c, 0x46, 0x46, 0x4b, 0x49, 0x4c, 0x43, 0x48, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x49, 0x49, 0x44, 0x50, 0x49, 0x49, 0x4c, 0x4f, 0x47, 0x4a, 0x49, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x47, 0x46, 0x4c, 0x43, 0x4e, 0x44, 0x44, 0x49, 0x4f, 0x4c, 0x4d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x05, 0x0a, 0x23, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a,
	0x0e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x12, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0xb2, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x10, 0x62, 0x75, 0x66,
	0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a,
	0x0b, 0x45, 0x4b, 0x45, 0x42, 0x41, 0x4f, 0x49, 0x45, 0x47, 0x50, 0x43, 0x18, 0x99, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x49, 0x44, 0x50, 0x49, 0x49, 0x4c, 0x4f, 0x47, 0x4a,
	0x49, 0x48, 0x00, 0x52, 0x0b, 0x45, 0x4b, 0x45, 0x42, 0x41, 0x4f, 0x49, 0x45, 0x47, 0x50, 0x43,
	0x12, 0x31, 0x0a, 0x0b, 0x46, 0x4b, 0x46, 0x50, 0x43, 0x49, 0x4f, 0x44, 0x48, 0x4e, 0x45, 0x18,
	0xa7, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4d, 0x50, 0x4f, 0x4a, 0x4f, 0x42,
	0x4f, 0x4b, 0x4f, 0x50, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x4b, 0x46, 0x50, 0x43, 0x49, 0x4f, 0x44,
	0x48, 0x4e, 0x45, 0x12, 0x54, 0x0a, 0x19, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x72, 0x65, 0x72, 0x6f,
	0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0xb4, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42,
	0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48,
	0x00, 0x52, 0x16, 0x62, 0x75, 0x66, 0x66, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x15, 0x6d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0xb9, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x13, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4d,
	0x4e, 0x4e, 0x46, 0x46, 0x45, 0x42, 0x4c, 0x4e, 0x42, 0x48, 0x18, 0xf7, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4b, 0x43, 0x50, 0x47, 0x45, 0x4c, 0x4a, 0x4d, 0x47, 0x41, 0x48,
	0x00, 0x52, 0x0b, 0x4d, 0x4e, 0x4e, 0x46, 0x46, 0x45, 0x42, 0x4c, 0x4e, 0x42, 0x48, 0x12, 0x31,
	0x0a, 0x0b, 0x44, 0x4a, 0x4a, 0x46, 0x4e, 0x48, 0x43, 0x46, 0x4c, 0x4a, 0x4e, 0x18, 0xba, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4d, 0x4c, 0x46, 0x46, 0x4b, 0x49, 0x4c, 0x43,
	0x48, 0x49, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x4a, 0x4a, 0x46, 0x4e, 0x48, 0x43, 0x46, 0x4c, 0x4a,
	0x4e, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x45, 0x48, 0x50, 0x4f, 0x4e, 0x4f, 0x50, 0x49, 0x4d, 0x47,
	0x18, 0xe1, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x43, 0x41, 0x45, 0x4d, 0x42,
	0x4d, 0x4a, 0x47, 0x4d, 0x4a, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x45, 0x48, 0x50, 0x4f, 0x4e, 0x4f,
	0x50, 0x49, 0x4d, 0x47, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x4f, 0x42, 0x4c, 0x43, 0x43, 0x4e, 0x44,
	0x4a, 0x4e, 0x4f, 0x18, 0xe0, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x46, 0x4c,
	0x43, 0x4e, 0x44, 0x44, 0x49, 0x4f, 0x4c, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x4f, 0x42, 0x4c,
	0x43, 0x43, 0x4e, 0x44, 0x4a, 0x4e, 0x4f, 0x12, 0x4a, 0x0a, 0x13, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x84,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x11, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescOnce sync.Once
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescData = file_HandleRogueCommonPendingActionCsReq_proto_rawDesc
)

func file_HandleRogueCommonPendingActionCsReq_proto_rawDescGZIP() []byte {
	file_HandleRogueCommonPendingActionCsReq_proto_rawDescOnce.Do(func() {
		file_HandleRogueCommonPendingActionCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HandleRogueCommonPendingActionCsReq_proto_rawDescData)
	})
	return file_HandleRogueCommonPendingActionCsReq_proto_rawDescData
}

var file_HandleRogueCommonPendingActionCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HandleRogueCommonPendingActionCsReq_proto_goTypes = []interface{}{
	(*HandleRogueCommonPendingActionCsReq)(nil), // 0: HandleRogueCommonPendingActionCsReq
	(*RogueBuffSelectResult)(nil),               // 1: RogueBuffSelectResult
	(*IIDPIILOGJI)(nil),                         // 2: IIDPIILOGJI
	(*KMPOJOBOKOP)(nil),                         // 3: KMPOJOBOKOP
	(*RogueBuffRerollResult)(nil),               // 4: RogueBuffRerollResult
	(*RogueMiracleSelectResult)(nil),            // 5: RogueMiracleSelectResult
	(*DKCPGELJMGA)(nil),                         // 6: DKCPGELJMGA
	(*KMLFFKILCHI)(nil),                         // 7: KMLFFKILCHI
	(*MCAEMBMJGMJ)(nil),                         // 8: MCAEMBMJGMJ
	(*GFLCNDDIOLM)(nil),                         // 9: GFLCNDDIOLM
	(*RogueBonusSelectResult)(nil),              // 10: RogueBonusSelectResult
}
var file_HandleRogueCommonPendingActionCsReq_proto_depIdxs = []int32{
	1,  // 0: HandleRogueCommonPendingActionCsReq.buff_select_result:type_name -> RogueBuffSelectResult
	2,  // 1: HandleRogueCommonPendingActionCsReq.EKEBAOIEGPC:type_name -> IIDPIILOGJI
	3,  // 2: HandleRogueCommonPendingActionCsReq.FKFPCIODHNE:type_name -> KMPOJOBOKOP
	4,  // 3: HandleRogueCommonPendingActionCsReq.buff_reroll_select_result:type_name -> RogueBuffRerollResult
	5,  // 4: HandleRogueCommonPendingActionCsReq.miracle_select_result:type_name -> RogueMiracleSelectResult
	6,  // 5: HandleRogueCommonPendingActionCsReq.MNNFFEBLNBH:type_name -> DKCPGELJMGA
	7,  // 6: HandleRogueCommonPendingActionCsReq.DJJFNHCFLJN:type_name -> KMLFFKILCHI
	8,  // 7: HandleRogueCommonPendingActionCsReq.GEHPONOPIMG:type_name -> MCAEMBMJGMJ
	9,  // 8: HandleRogueCommonPendingActionCsReq.GOBLCCNDJNO:type_name -> GFLCNDDIOLM
	10, // 9: HandleRogueCommonPendingActionCsReq.bonus_select_result:type_name -> RogueBonusSelectResult
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_HandleRogueCommonPendingActionCsReq_proto_init() }
func file_HandleRogueCommonPendingActionCsReq_proto_init() {
	if File_HandleRogueCommonPendingActionCsReq_proto != nil {
		return
	}
	file_RogueBuffRerollResult_proto_init()
	file_RogueBonusSelectResult_proto_init()
	file_KMPOJOBOKOP_proto_init()
	file_DKCPGELJMGA_proto_init()
	file_RogueBuffSelectResult_proto_init()
	file_MCAEMBMJGMJ_proto_init()
	file_KMLFFKILCHI_proto_init()
	file_RogueMiracleSelectResult_proto_init()
	file_IIDPIILOGJI_proto_init()
	file_GFLCNDDIOLM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleRogueCommonPendingActionCsReq); i {
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
	file_HandleRogueCommonPendingActionCsReq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HandleRogueCommonPendingActionCsReq_BuffSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_EKEBAOIEGPC)(nil),
		(*HandleRogueCommonPendingActionCsReq_FKFPCIODHNE)(nil),
		(*HandleRogueCommonPendingActionCsReq_BuffRerollSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_MiracleSelectResult)(nil),
		(*HandleRogueCommonPendingActionCsReq_MNNFFEBLNBH)(nil),
		(*HandleRogueCommonPendingActionCsReq_DJJFNHCFLJN)(nil),
		(*HandleRogueCommonPendingActionCsReq_GEHPONOPIMG)(nil),
		(*HandleRogueCommonPendingActionCsReq_GOBLCCNDJNO)(nil),
		(*HandleRogueCommonPendingActionCsReq_BonusSelectResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HandleRogueCommonPendingActionCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HandleRogueCommonPendingActionCsReq_proto_goTypes,
		DependencyIndexes: file_HandleRogueCommonPendingActionCsReq_proto_depIdxs,
		MessageInfos:      file_HandleRogueCommonPendingActionCsReq_proto_msgTypes,
	}.Build()
	File_HandleRogueCommonPendingActionCsReq_proto = out.File
	file_HandleRogueCommonPendingActionCsReq_proto_rawDesc = nil
	file_HandleRogueCommonPendingActionCsReq_proto_goTypes = nil
	file_HandleRogueCommonPendingActionCsReq_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: COBGPBGFJPP.proto

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

type COBGPBGFJPP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndTime     int64  `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	NIKFGGJACEC uint32 `protobuf:"varint,5,opt,name=NIKFGGJACEC,proto3" json:"NIKFGGJACEC,omitempty"`
	BeginTime   int64  `protobuf:"varint,7,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
}

func (x *COBGPBGFJPP) Reset() {
	*x = COBGPBGFJPP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_COBGPBGFJPP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COBGPBGFJPP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COBGPBGFJPP) ProtoMessage() {}

func (x *COBGPBGFJPP) ProtoReflect() protoreflect.Message {
	mi := &file_COBGPBGFJPP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COBGPBGFJPP.ProtoReflect.Descriptor instead.
func (*COBGPBGFJPP) Descriptor() ([]byte, []int) {
	return file_COBGPBGFJPP_proto_rawDescGZIP(), []int{0}
}

func (x *COBGPBGFJPP) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *COBGPBGFJPP) GetNIKFGGJACEC() uint32 {
	if x != nil {
		return x.NIKFGGJACEC
	}
	return 0
}

func (x *COBGPBGFJPP) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

var File_COBGPBGFJPP_proto protoreflect.FileDescriptor

var file_COBGPBGFJPP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x4f, 0x42, 0x47, 0x50, 0x42, 0x47, 0x46, 0x4a, 0x50, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0b, 0x43, 0x4f, 0x42, 0x47, 0x50, 0x42, 0x47, 0x46, 0x4a,
	0x50, 0x50, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x4e, 0x49, 0x4b, 0x46, 0x47, 0x47, 0x4a, 0x41, 0x43, 0x45, 0x43, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x49, 0x4b, 0x46, 0x47, 0x47, 0x4a, 0x41, 0x43, 0x45, 0x43, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_COBGPBGFJPP_proto_rawDescOnce sync.Once
	file_COBGPBGFJPP_proto_rawDescData = file_COBGPBGFJPP_proto_rawDesc
)

func file_COBGPBGFJPP_proto_rawDescGZIP() []byte {
	file_COBGPBGFJPP_proto_rawDescOnce.Do(func() {
		file_COBGPBGFJPP_proto_rawDescData = protoimpl.X.CompressGZIP(file_COBGPBGFJPP_proto_rawDescData)
	})
	return file_COBGPBGFJPP_proto_rawDescData
}

var file_COBGPBGFJPP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_COBGPBGFJPP_proto_goTypes = []interface{}{
	(*COBGPBGFJPP)(nil), // 0: COBGPBGFJPP
}
var file_COBGPBGFJPP_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_COBGPBGFJPP_proto_init() }
func file_COBGPBGFJPP_proto_init() {
	if File_COBGPBGFJPP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_COBGPBGFJPP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COBGPBGFJPP); i {
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
			RawDescriptor: file_COBGPBGFJPP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_COBGPBGFJPP_proto_goTypes,
		DependencyIndexes: file_COBGPBGFJPP_proto_depIdxs,
		MessageInfos:      file_COBGPBGFJPP_proto_msgTypes,
	}.Build()
	File_COBGPBGFJPP_proto = out.File
	file_COBGPBGFJPP_proto_rawDesc = nil
	file_COBGPBGFJPP_proto_goTypes = nil
	file_COBGPBGFJPP_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/base/v1/base.proto

package v1

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

type BaseListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastScore int64  `protobuf:"varint,1,opt,name=last_score,json=lastScore,proto3" json:"last_score,omitempty"`
	Limit     uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *BaseListRequest) Reset() {
	*x = BaseListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_v1_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseListRequest) ProtoMessage() {}

func (x *BaseListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_v1_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseListRequest.ProtoReflect.Descriptor instead.
func (*BaseListRequest) Descriptor() ([]byte, []int) {
	return file_api_base_v1_base_proto_rawDescGZIP(), []int{0}
}

func (x *BaseListRequest) GetLastScore() int64 {
	if x != nil {
		return x.LastScore
	}
	return 0
}

func (x *BaseListRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type BaseListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastScore int64 `protobuf:"varint,1,opt,name=last_score,json=lastScore,proto3" json:"last_score,omitempty"`
	HasMore   bool  `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *BaseListResponse) Reset() {
	*x = BaseListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_v1_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseListResponse) ProtoMessage() {}

func (x *BaseListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_v1_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseListResponse.ProtoReflect.Descriptor instead.
func (*BaseListResponse) Descriptor() ([]byte, []int) {
	return file_api_base_v1_base_proto_rawDescGZIP(), []int{1}
}

func (x *BaseListResponse) GetLastScore() int64 {
	if x != nil {
		return x.LastScore
	}
	return 0
}

func (x *BaseListResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

var File_api_base_v1_base_proto protoreflect.FileDescriptor

var file_api_base_v1_base_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x46, 0x0a, 0x0f, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x4c, 0x0a,
	0x10, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x34,
	0x63, 0x6c, 0x61, 0x62, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_base_v1_base_proto_rawDescOnce sync.Once
	file_api_base_v1_base_proto_rawDescData = file_api_base_v1_base_proto_rawDesc
)

func file_api_base_v1_base_proto_rawDescGZIP() []byte {
	file_api_base_v1_base_proto_rawDescOnce.Do(func() {
		file_api_base_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_base_v1_base_proto_rawDescData)
	})
	return file_api_base_v1_base_proto_rawDescData
}

var file_api_base_v1_base_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_base_v1_base_proto_goTypes = []interface{}{
	(*BaseListRequest)(nil),  // 0: api.base.v1.BaseListRequest
	(*BaseListResponse)(nil), // 1: api.base.v1.BaseListResponse
}
var file_api_base_v1_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_base_v1_base_proto_init() }
func file_api_base_v1_base_proto_init() {
	if File_api_base_v1_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_base_v1_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseListRequest); i {
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
		file_api_base_v1_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseListResponse); i {
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
			RawDescriptor: file_api_base_v1_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_base_v1_base_proto_goTypes,
		DependencyIndexes: file_api_base_v1_base_proto_depIdxs,
		MessageInfos:      file_api_base_v1_base_proto_msgTypes,
	}.Build()
	File_api_base_v1_base_proto = out.File
	file_api_base_v1_base_proto_rawDesc = nil
	file_api_base_v1_base_proto_goTypes = nil
	file_api_base_v1_base_proto_depIdxs = nil
}

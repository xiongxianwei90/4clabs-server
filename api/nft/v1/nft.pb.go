// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/nft/v1/nft.proto

package v1

import (
	v1 "4clabs-server/api/base/v1"
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

type GetNftDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	TokenId         string `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (x *GetNftDetailRequest) Reset() {
	*x = GetNftDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_nft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNftDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNftDetailRequest) ProtoMessage() {}

func (x *GetNftDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_nft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNftDetailRequest.ProtoReflect.Descriptor instead.
func (*GetNftDetailRequest) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_nft_proto_rawDescGZIP(), []int{0}
}

func (x *GetNftDetailRequest) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *GetNftDetailRequest) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

type GetNftDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail *Detail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *GetNftDetailResponse) Reset() {
	*x = GetNftDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_nft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNftDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNftDetailResponse) ProtoMessage() {}

func (x *GetNftDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_nft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNftDetailResponse.ProtoReflect.Descriptor instead.
func (*GetNftDetailResponse) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_nft_proto_rawDescGZIP(), []int{1}
}

func (x *GetNftDetailResponse) GetDetail() *Detail {
	if x != nil {
		return x.Detail
	}
	return nil
}

type GetAddressNftsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseListRequest *v1.BaseListRequest `protobuf:"bytes,1,opt,name=base_list_request,json=baseListRequest,proto3" json:"base_list_request,omitempty"`
	Address         string              `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetAddressNftsRequest) Reset() {
	*x = GetAddressNftsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_nft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAddressNftsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressNftsRequest) ProtoMessage() {}

func (x *GetAddressNftsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_nft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressNftsRequest.ProtoReflect.Descriptor instead.
func (*GetAddressNftsRequest) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_nft_proto_rawDescGZIP(), []int{2}
}

func (x *GetAddressNftsRequest) GetBaseListRequest() *v1.BaseListRequest {
	if x != nil {
		return x.BaseListRequest
	}
	return nil
}

func (x *GetAddressNftsRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetAddressNftResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseListResponse *v1.BaseListResponse `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3" json:"base_list_response,omitempty"`
	Summaries        []*Summary           `protobuf:"bytes,2,rep,name=summaries,proto3" json:"summaries,omitempty"`
}

func (x *GetAddressNftResponse) Reset() {
	*x = GetAddressNftResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_nft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAddressNftResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressNftResponse) ProtoMessage() {}

func (x *GetAddressNftResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_nft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressNftResponse.ProtoReflect.Descriptor instead.
func (*GetAddressNftResponse) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_nft_proto_rawDescGZIP(), []int{3}
}

func (x *GetAddressNftResponse) GetBaseListResponse() *v1.BaseListResponse {
	if x != nil {
		return x.BaseListResponse
	}
	return nil
}

func (x *GetAddressNftResponse) GetSummaries() []*Summary {
	if x != nil {
		return x.Summaries
	}
	return nil
}

var File_api_nft_v1_nft_proto protoreflect.FileDescriptor

var file_api_nft_v1_nft_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f,
	0x6e, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x42,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x22, 0x7b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4e, 0x66, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x11, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x97, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4e, 0x66,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x12, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x09,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x34, 0x63, 0x6c,
	0x61, 0x62, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e,
	0x66, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_nft_v1_nft_proto_rawDescOnce sync.Once
	file_api_nft_v1_nft_proto_rawDescData = file_api_nft_v1_nft_proto_rawDesc
)

func file_api_nft_v1_nft_proto_rawDescGZIP() []byte {
	file_api_nft_v1_nft_proto_rawDescOnce.Do(func() {
		file_api_nft_v1_nft_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_nft_v1_nft_proto_rawDescData)
	})
	return file_api_nft_v1_nft_proto_rawDescData
}

var file_api_nft_v1_nft_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_nft_v1_nft_proto_goTypes = []interface{}{
	(*GetNftDetailRequest)(nil),   // 0: api.nft.v1.GetNftDetailRequest
	(*GetNftDetailResponse)(nil),  // 1: api.nft.v1.GetNftDetailResponse
	(*GetAddressNftsRequest)(nil), // 2: api.nft.v1.GetAddressNftsRequest
	(*GetAddressNftResponse)(nil), // 3: api.nft.v1.GetAddressNftResponse
	(*Detail)(nil),                // 4: api.nft.v1.Detail
	(*v1.BaseListRequest)(nil),    // 5: api.base.v1.BaseListRequest
	(*v1.BaseListResponse)(nil),   // 6: api.base.v1.BaseListResponse
	(*Summary)(nil),               // 7: api.nft.v1.Summary
}
var file_api_nft_v1_nft_proto_depIdxs = []int32{
	4, // 0: api.nft.v1.GetNftDetailResponse.detail:type_name -> api.nft.v1.Detail
	5, // 1: api.nft.v1.GetAddressNftsRequest.base_list_request:type_name -> api.base.v1.BaseListRequest
	6, // 2: api.nft.v1.GetAddressNftResponse.base_list_response:type_name -> api.base.v1.BaseListResponse
	7, // 3: api.nft.v1.GetAddressNftResponse.summaries:type_name -> api.nft.v1.Summary
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_nft_v1_nft_proto_init() }
func file_api_nft_v1_nft_proto_init() {
	if File_api_nft_v1_nft_proto != nil {
		return
	}
	file_api_nft_v1_typs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_nft_v1_nft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNftDetailRequest); i {
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
		file_api_nft_v1_nft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNftDetailResponse); i {
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
		file_api_nft_v1_nft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAddressNftsRequest); i {
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
		file_api_nft_v1_nft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAddressNftResponse); i {
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
			RawDescriptor: file_api_nft_v1_nft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_nft_v1_nft_proto_goTypes,
		DependencyIndexes: file_api_nft_v1_nft_proto_depIdxs,
		MessageInfos:      file_api_nft_v1_nft_proto_msgTypes,
	}.Build()
	File_api_nft_v1_nft_proto = out.File
	file_api_nft_v1_nft_proto_rawDesc = nil
	file_api_nft_v1_nft_proto_goTypes = nil
	file_api_nft_v1_nft_proto_depIdxs = nil
}

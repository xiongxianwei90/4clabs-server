// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/nft/v1/comic_nft.proto

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

type ListComicNftRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseListRequest *v1.BaseListRequest `protobuf:"bytes,1,opt,name=base_list_request,json=baseListRequest,proto3" json:"base_list_request,omitempty"`
}

func (x *ListComicNftRequest) Reset() {
	*x = ListComicNftRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_comic_nft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComicNftRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComicNftRequest) ProtoMessage() {}

func (x *ListComicNftRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_comic_nft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComicNftRequest.ProtoReflect.Descriptor instead.
func (*ListComicNftRequest) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_comic_nft_proto_rawDescGZIP(), []int{0}
}

func (x *ListComicNftRequest) GetBaseListRequest() *v1.BaseListRequest {
	if x != nil {
		return x.BaseListRequest
	}
	return nil
}

type ListComicNftResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseListResponse *v1.BaseListResponse `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3" json:"base_list_response,omitempty"`
	ComicNft         []*ComicNft          `protobuf:"bytes,2,rep,name=comic_nft,json=comicNft,proto3" json:"comic_nft,omitempty"`
}

func (x *ListComicNftResponse) Reset() {
	*x = ListComicNftResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_comic_nft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComicNftResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComicNftResponse) ProtoMessage() {}

func (x *ListComicNftResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_comic_nft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComicNftResponse.ProtoReflect.Descriptor instead.
func (*ListComicNftResponse) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_comic_nft_proto_rawDescGZIP(), []int{1}
}

func (x *ListComicNftResponse) GetBaseListResponse() *v1.BaseListResponse {
	if x != nil {
		return x.BaseListResponse
	}
	return nil
}

func (x *ListComicNftResponse) GetComicNft() []*ComicNft {
	if x != nil {
		return x.ComicNft
	}
	return nil
}

type PurchaseComicNftRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenId      string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	BuyerAddress string `protobuf:"bytes,2,opt,name=buyer_address,json=buyerAddress,proto3" json:"buyer_address,omitempty"`
}

func (x *PurchaseComicNftRequest) Reset() {
	*x = PurchaseComicNftRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_comic_nft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseComicNftRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseComicNftRequest) ProtoMessage() {}

func (x *PurchaseComicNftRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_comic_nft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseComicNftRequest.ProtoReflect.Descriptor instead.
func (*PurchaseComicNftRequest) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_comic_nft_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseComicNftRequest) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *PurchaseComicNftRequest) GetBuyerAddress() string {
	if x != nil {
		return x.BuyerAddress
	}
	return ""
}

type PurchaseComicNftResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PurchaseComicNftResponse) Reset() {
	*x = PurchaseComicNftResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_comic_nft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseComicNftResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseComicNftResponse) ProtoMessage() {}

func (x *PurchaseComicNftResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_comic_nft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseComicNftResponse.ProtoReflect.Descriptor instead.
func (*PurchaseComicNftResponse) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_comic_nft_proto_rawDescGZIP(), []int{3}
}

type ListComicNftByComicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComicId         string              `protobuf:"bytes,1,opt,name=comic_id,json=comicId,proto3" json:"comic_id,omitempty"`
	BaseListRequest *v1.BaseListRequest `protobuf:"bytes,2,opt,name=base_list_request,json=baseListRequest,proto3" json:"base_list_request,omitempty"`
}

func (x *ListComicNftByComicRequest) Reset() {
	*x = ListComicNftByComicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_comic_nft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComicNftByComicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComicNftByComicRequest) ProtoMessage() {}

func (x *ListComicNftByComicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_comic_nft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComicNftByComicRequest.ProtoReflect.Descriptor instead.
func (*ListComicNftByComicRequest) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_comic_nft_proto_rawDescGZIP(), []int{4}
}

func (x *ListComicNftByComicRequest) GetComicId() string {
	if x != nil {
		return x.ComicId
	}
	return ""
}

func (x *ListComicNftByComicRequest) GetBaseListRequest() *v1.BaseListRequest {
	if x != nil {
		return x.BaseListRequest
	}
	return nil
}

type ListComicNftByComicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseListResponse *v1.BaseListResponse `protobuf:"bytes,1,opt,name=base_list_response,json=baseListResponse,proto3" json:"base_list_response,omitempty"`
	ComicNft         []*ComicNft          `protobuf:"bytes,2,rep,name=comic_nft,json=comicNft,proto3" json:"comic_nft,omitempty"`
}

func (x *ListComicNftByComicResponse) Reset() {
	*x = ListComicNftByComicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nft_v1_comic_nft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComicNftByComicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComicNftByComicResponse) ProtoMessage() {}

func (x *ListComicNftByComicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_nft_v1_comic_nft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComicNftByComicResponse.ProtoReflect.Descriptor instead.
func (*ListComicNftByComicResponse) Descriptor() ([]byte, []int) {
	return file_api_nft_v1_comic_nft_proto_rawDescGZIP(), []int{5}
}

func (x *ListComicNftByComicResponse) GetBaseListResponse() *v1.BaseListResponse {
	if x != nil {
		return x.BaseListResponse
	}
	return nil
}

func (x *ListComicNftByComicResponse) GetComicNft() []*ComicNft {
	if x != nil {
		return x.ComicNft
	}
	return nil
}

var File_api_nft_v1_comic_nft_proto protoreflect.FileDescriptor

var file_api_nft_v1_comic_nft_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x69, 0x63, 0x5f, 0x6e, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48,
	0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4b, 0x0a, 0x12, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x62, 0x61,
	0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x6e, 0x66, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66,
	0x74, 0x22, 0x59, 0x0a, 0x17, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d,
	0x69, 0x63, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x1a, 0x0a, 0x18,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x69, 0x63,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x69, 0x63,
	0x49, 0x64, 0x12, 0x48, 0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x62, 0x61, 0x73,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9d, 0x01, 0x0a,
	0x1b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66, 0x74, 0x42, 0x79, 0x43,
	0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x12,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x69, 0x63, 0x5f, 0x6e, 0x66, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x4e,
	0x66, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x4e, 0x66, 0x74, 0x42, 0x2b, 0x0a, 0x0a,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x34, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6e, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_nft_v1_comic_nft_proto_rawDescOnce sync.Once
	file_api_nft_v1_comic_nft_proto_rawDescData = file_api_nft_v1_comic_nft_proto_rawDesc
)

func file_api_nft_v1_comic_nft_proto_rawDescGZIP() []byte {
	file_api_nft_v1_comic_nft_proto_rawDescOnce.Do(func() {
		file_api_nft_v1_comic_nft_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_nft_v1_comic_nft_proto_rawDescData)
	})
	return file_api_nft_v1_comic_nft_proto_rawDescData
}

var file_api_nft_v1_comic_nft_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_nft_v1_comic_nft_proto_goTypes = []interface{}{
	(*ListComicNftRequest)(nil),         // 0: api.nft.v1.ListComicNftRequest
	(*ListComicNftResponse)(nil),        // 1: api.nft.v1.ListComicNftResponse
	(*PurchaseComicNftRequest)(nil),     // 2: api.nft.v1.PurchaseComicNftRequest
	(*PurchaseComicNftResponse)(nil),    // 3: api.nft.v1.PurchaseComicNftResponse
	(*ListComicNftByComicRequest)(nil),  // 4: api.nft.v1.ListComicNftByComicRequest
	(*ListComicNftByComicResponse)(nil), // 5: api.nft.v1.ListComicNftByComicResponse
	(*v1.BaseListRequest)(nil),          // 6: api.base.v1.BaseListRequest
	(*v1.BaseListResponse)(nil),         // 7: api.base.v1.BaseListResponse
	(*ComicNft)(nil),                    // 8: api.nft.v1.ComicNft
}
var file_api_nft_v1_comic_nft_proto_depIdxs = []int32{
	6, // 0: api.nft.v1.ListComicNftRequest.base_list_request:type_name -> api.base.v1.BaseListRequest
	7, // 1: api.nft.v1.ListComicNftResponse.base_list_response:type_name -> api.base.v1.BaseListResponse
	8, // 2: api.nft.v1.ListComicNftResponse.comic_nft:type_name -> api.nft.v1.ComicNft
	6, // 3: api.nft.v1.ListComicNftByComicRequest.base_list_request:type_name -> api.base.v1.BaseListRequest
	7, // 4: api.nft.v1.ListComicNftByComicResponse.base_list_response:type_name -> api.base.v1.BaseListResponse
	8, // 5: api.nft.v1.ListComicNftByComicResponse.comic_nft:type_name -> api.nft.v1.ComicNft
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_nft_v1_comic_nft_proto_init() }
func file_api_nft_v1_comic_nft_proto_init() {
	if File_api_nft_v1_comic_nft_proto != nil {
		return
	}
	file_api_nft_v1_typs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_nft_v1_comic_nft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComicNftRequest); i {
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
		file_api_nft_v1_comic_nft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComicNftResponse); i {
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
		file_api_nft_v1_comic_nft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseComicNftRequest); i {
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
		file_api_nft_v1_comic_nft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseComicNftResponse); i {
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
		file_api_nft_v1_comic_nft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComicNftByComicRequest); i {
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
		file_api_nft_v1_comic_nft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComicNftByComicResponse); i {
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
			RawDescriptor: file_api_nft_v1_comic_nft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_nft_v1_comic_nft_proto_goTypes,
		DependencyIndexes: file_api_nft_v1_comic_nft_proto_depIdxs,
		MessageInfos:      file_api_nft_v1_comic_nft_proto_msgTypes,
	}.Build()
	File_api_nft_v1_comic_nft_proto = out.File
	file_api_nft_v1_comic_nft_proto_rawDesc = nil
	file_api_nft_v1_comic_nft_proto_goTypes = nil
	file_api_nft_v1_comic_nft_proto_depIdxs = nil
}

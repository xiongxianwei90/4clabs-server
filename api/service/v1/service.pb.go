// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: api/service/v1/service.proto

package v1

import (
	v12 "4clabs-server/api/auth/v1"
	v1 "4clabs-server/api/nft/v1"
	v11 "4clabs-server/api/tickets/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_service_v1_service_proto protoreflect.FileDescriptor

var file_api_service_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x16,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x66, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfc, 0x08, 0x0a, 0x03, 0x4e, 0x66, 0x74, 0x12,
	0x7c, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x6f, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x12, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x63,
	0x6f, 0x6d, 0x69, 0x63, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6b,
	0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x66, 0x74, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x73, 0x69, 0x74, 0x65, 0x72, 0x4e, 0x66, 0x74, 0x73, 0x12,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x66, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x7e, 0x0a, 0x12, 0x49, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x57, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x4d, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x74, 0x4d,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x63, 0x61, 0x6e, 0x5f, 0x6d,
	0x69, 0x6e, 0x74, 0x12, 0x88, 0x01, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x6f, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x6f, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x69,
	0x67, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x85,
	0x01, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f,
	0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x7a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4e, 0x66, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e,
	0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4e, 0x66, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4e, 0x66, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x6e, 0x66,
	0x74, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x66, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x12, 0x36,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x21, 0x5a, 0x1f, 0x34, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_api_service_v1_service_proto_goTypes = []interface{}{
	(*v1.ListComicWorkRequest)(nil),           // 0: api.nft.v1.ListComicWorkRequest
	(*v1.ComicCreateRequest)(nil),             // 1: api.nft.v1.ComicCreateRequest
	(*v1.RegisterNftRequest)(nil),             // 2: api.nft.v1.RegisterNftRequest
	(*v1.ListRegisterNftRequest)(nil),         // 3: api.nft.v1.ListRegisterNftRequest
	(*v11.CanMintRequest)(nil),                // 4: api.tickets.v1.CanMintRequest
	(*v12.VerifySignToLoginSignRequest)(nil),  // 5: api.auth.v1.VerifySignToLoginSignRequest
	(*v12.FetchSignMessageRequest)(nil),       // 6: api.auth.v1.FetchSignMessageRequest
	(*v1.GetAddressNftsRequest)(nil),          // 7: api.nft.v1.GetAddressNftsRequest
	(*v1.GetNftDetailRequest)(nil),            // 8: api.nft.v1.GetNftDetailRequest
	(*v1.ListComicWorkResponse)(nil),          // 9: api.nft.v1.ListComicWorkResponse
	(*v1.ComicCreateResponse)(nil),            // 10: api.nft.v1.ComicCreateResponse
	(*v1.RegisterNftResponse)(nil),            // 11: api.nft.v1.RegisterNftResponse
	(*v1.ListRegisterNftResponse)(nil),        // 12: api.nft.v1.ListRegisterNftResponse
	(*v11.CantMintResponse)(nil),              // 13: api.tickets.v1.CantMintResponse
	(*v12.VerifySignToLoginSighResponse)(nil), // 14: api.auth.v1.VerifySignToLoginSighResponse
	(*v12.FetchSignMessageResponse)(nil),      // 15: api.auth.v1.FetchSignMessageResponse
	(*v1.GetAddressNftResponse)(nil),          // 16: api.nft.v1.GetAddressNftResponse
	(*v1.GetNftDetailResponse)(nil),           // 17: api.nft.v1.GetNftDetailResponse
}
var file_api_service_v1_service_proto_depIdxs = []int32{
	0,  // 0: api.service.v1.Nft.ListComicWorks:input_type -> api.nft.v1.ListComicWorkRequest
	1,  // 1: api.service.v1.Nft.CreateComic:input_type -> api.nft.v1.ComicCreateRequest
	2,  // 2: api.service.v1.Nft.RegisterNft:input_type -> api.nft.v1.RegisterNftRequest
	3,  // 3: api.service.v1.Nft.ListRegsiterNfts:input_type -> api.nft.v1.ListRegisterNftRequest
	4,  // 4: api.service.v1.Nft.InTicketsWLRequest:input_type -> api.tickets.v1.CanMintRequest
	5,  // 5: api.service.v1.Nft.SignToLogin:input_type -> api.auth.v1.VerifySignToLoginSignRequest
	6,  // 6: api.service.v1.Nft.FetchNonce:input_type -> api.auth.v1.FetchSignMessageRequest
	7,  // 7: api.service.v1.Nft.GetAddressNfts:input_type -> api.nft.v1.GetAddressNftsRequest
	8,  // 8: api.service.v1.Nft.GetNftDetail:input_type -> api.nft.v1.GetNftDetailRequest
	9,  // 9: api.service.v1.Nft.ListComicWorks:output_type -> api.nft.v1.ListComicWorkResponse
	10, // 10: api.service.v1.Nft.CreateComic:output_type -> api.nft.v1.ComicCreateResponse
	11, // 11: api.service.v1.Nft.RegisterNft:output_type -> api.nft.v1.RegisterNftResponse
	12, // 12: api.service.v1.Nft.ListRegsiterNfts:output_type -> api.nft.v1.ListRegisterNftResponse
	13, // 13: api.service.v1.Nft.InTicketsWLRequest:output_type -> api.tickets.v1.CantMintResponse
	14, // 14: api.service.v1.Nft.SignToLogin:output_type -> api.auth.v1.VerifySignToLoginSighResponse
	15, // 15: api.service.v1.Nft.FetchNonce:output_type -> api.auth.v1.FetchSignMessageResponse
	16, // 16: api.service.v1.Nft.GetAddressNfts:output_type -> api.nft.v1.GetAddressNftResponse
	17, // 17: api.service.v1.Nft.GetNftDetail:output_type -> api.nft.v1.GetNftDetailResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_service_v1_service_proto_init() }
func file_api_service_v1_service_proto_init() {
	if File_api_service_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_service_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_v1_service_proto_goTypes,
		DependencyIndexes: file_api_service_v1_service_proto_depIdxs,
	}.Build()
	File_api_service_v1_service_proto = out.File
	file_api_service_v1_service_proto_rawDesc = nil
	file_api_service_v1_service_proto_goTypes = nil
	file_api_service_v1_service_proto_depIdxs = nil
}

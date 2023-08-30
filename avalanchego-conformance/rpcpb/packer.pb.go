// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: rpcpb/packer.proto

package rpcpb

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

type BuildVertexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodecVersion uint32   `protobuf:"varint,1,opt,name=codec_version,json=codecVersion,proto3" json:"codec_version,omitempty"`
	ChainId      []byte   `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Height       uint64   `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Epoch        uint32   `protobuf:"varint,4,opt,name=epoch,proto3" json:"epoch,omitempty"`
	ParentIds    [][]byte `protobuf:"bytes,5,rep,name=parent_ids,json=parentIds,proto3" json:"parent_ids,omitempty"`
	Txs          [][]byte `protobuf:"bytes,6,rep,name=txs,proto3" json:"txs,omitempty"`
	VtxBytes     []byte   `protobuf:"bytes,7,opt,name=vtx_bytes,json=vtxBytes,proto3" json:"vtx_bytes,omitempty"`
}

func (x *BuildVertexRequest) Reset() {
	*x = BuildVertexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_packer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildVertexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildVertexRequest) ProtoMessage() {}

func (x *BuildVertexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_packer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildVertexRequest.ProtoReflect.Descriptor instead.
func (*BuildVertexRequest) Descriptor() ([]byte, []int) {
	return file_rpcpb_packer_proto_rawDescGZIP(), []int{0}
}

func (x *BuildVertexRequest) GetCodecVersion() uint32 {
	if x != nil {
		return x.CodecVersion
	}
	return 0
}

func (x *BuildVertexRequest) GetChainId() []byte {
	if x != nil {
		return x.ChainId
	}
	return nil
}

func (x *BuildVertexRequest) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BuildVertexRequest) GetEpoch() uint32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *BuildVertexRequest) GetParentIds() [][]byte {
	if x != nil {
		return x.ParentIds
	}
	return nil
}

func (x *BuildVertexRequest) GetTxs() [][]byte {
	if x != nil {
		return x.Txs
	}
	return nil
}

func (x *BuildVertexRequest) GetVtxBytes() []byte {
	if x != nil {
		return x.VtxBytes
	}
	return nil
}

type BuildVertexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpectedBytes []byte `protobuf:"bytes,1,opt,name=expected_bytes,json=expectedBytes,proto3" json:"expected_bytes,omitempty"`
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Success       bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BuildVertexResponse) Reset() {
	*x = BuildVertexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcpb_packer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildVertexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildVertexResponse) ProtoMessage() {}

func (x *BuildVertexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcpb_packer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildVertexResponse.ProtoReflect.Descriptor instead.
func (*BuildVertexResponse) Descriptor() ([]byte, []int) {
	return file_rpcpb_packer_proto_rawDescGZIP(), []int{1}
}

func (x *BuildVertexResponse) GetExpectedBytes() []byte {
	if x != nil {
		return x.ExpectedBytes
	}
	return nil
}

func (x *BuildVertexResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BuildVertexResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_rpcpb_packer_proto protoreflect.FileDescriptor

var file_rpcpb_packer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x70, 0x63, 0x70, 0x62, 0x22, 0xd0, 0x01, 0x0a, 0x12,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x74, 0x78,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x76, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x70,
	0x0a, 0x13, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0x57, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78,
	0x12, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x56, 0x65,
	0x72, 0x74, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x70,
	0x63, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x2d, 0x72, 0x75, 0x73, 0x74, 0x2f,
	0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3b, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcpb_packer_proto_rawDescOnce sync.Once
	file_rpcpb_packer_proto_rawDescData = file_rpcpb_packer_proto_rawDesc
)

func file_rpcpb_packer_proto_rawDescGZIP() []byte {
	file_rpcpb_packer_proto_rawDescOnce.Do(func() {
		file_rpcpb_packer_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcpb_packer_proto_rawDescData)
	})
	return file_rpcpb_packer_proto_rawDescData
}

var file_rpcpb_packer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpcpb_packer_proto_goTypes = []interface{}{
	(*BuildVertexRequest)(nil),  // 0: rpcpb.BuildVertexRequest
	(*BuildVertexResponse)(nil), // 1: rpcpb.BuildVertexResponse
}
var file_rpcpb_packer_proto_depIdxs = []int32{
	0, // 0: rpcpb.PackerService.BuildVertex:input_type -> rpcpb.BuildVertexRequest
	1, // 1: rpcpb.PackerService.BuildVertex:output_type -> rpcpb.BuildVertexResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpcpb_packer_proto_init() }
func file_rpcpb_packer_proto_init() {
	if File_rpcpb_packer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcpb_packer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildVertexRequest); i {
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
		file_rpcpb_packer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildVertexResponse); i {
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
			RawDescriptor: file_rpcpb_packer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpcpb_packer_proto_goTypes,
		DependencyIndexes: file_rpcpb_packer_proto_depIdxs,
		MessageInfos:      file_rpcpb_packer_proto_msgTypes,
	}.Build()
	File_rpcpb_packer_proto = out.File
	file_rpcpb_packer_proto_rawDesc = nil
	file_rpcpb_packer_proto_goTypes = nil
	file_rpcpb_packer_proto_depIdxs = nil
}

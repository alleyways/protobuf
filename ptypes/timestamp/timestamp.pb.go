// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/golang/protobuf/ptypes/timestamp/timestamp.proto

package timestamp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

// Symbols defined in public import of google/protobuf/timestamp.proto

type Timestamp = timestamppb.Timestamp

var File_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto protoreflect.FileDescriptor

var file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3b, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDescOnce sync.Once
	file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDescData = file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDesc
)

func file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDescGZIP() []byte {
	file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDescOnce.Do(func() {
		file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDescData)
	})
	return file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDescData
}

var file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_goTypes = []interface{}{}
var file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_depIdxs = []int32{}

func init() { file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_init() }
func file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_init() {
	if File_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto != nil {
		return
	}
	File_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto = protoimpl.FileBuilder{
		RawDescriptor:     file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDesc,
		GoTypes:           file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_goTypes,
		DependencyIndexes: file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_depIdxs,
		FilesRegistry:     protoregistry.GlobalFiles,
		TypesRegistry:     protoregistry.GlobalTypes,
	}.Init()
	file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_rawDesc = nil
	file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_goTypes = nil
	file_github_com_golang_protobuf_ptypes_timestamp_timestamp_proto_depIdxs = nil
}

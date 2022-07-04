// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: creation/service/v1/creation_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type CreationErrorReason int32

const (
	CreationErrorReason_UNKNOWN_ERROR                   CreationErrorReason = 0
	CreationErrorReason_GET_ARTICLE_DRAFT_FAILED        CreationErrorReason = 1
	CreationErrorReason_GET_ARTICLE_LIST_FAILED         CreationErrorReason = 2
	CreationErrorReason_GET_DRAFT_LIST_FAILED           CreationErrorReason = 3
	CreationErrorReason_CREATE_ARTICLE_FAILED           CreationErrorReason = 4
	CreationErrorReason_CREATE_ARTICLE_DRAFT_FAILED     CreationErrorReason = 5
	CreationErrorReason_CREATE_ARTICLE_FOLDER_FAILED    CreationErrorReason = 6
	CreationErrorReason_CREATE_ARTICLE_STATISTIC_FAILED CreationErrorReason = 7
	CreationErrorReason_CREATE_ARTICLE_CACHE_FAILED     CreationErrorReason = 8
	CreationErrorReason_CREATE_ARTICLE_SEARCH_FAILED    CreationErrorReason = 9
	CreationErrorReason_DELETE_ARTICLE_DRAFT_FAILED     CreationErrorReason = 10
	CreationErrorReason_RECORD_NOT_FOUND                CreationErrorReason = 11
	CreationErrorReason_DRAFT_MARK_FAILED               CreationErrorReason = 12
	CreationErrorReason_SEND_ARTICLE_FAILED             CreationErrorReason = 13
	CreationErrorReason_SEND_TO_MQ_FAILED               CreationErrorReason = 14
)

// Enum value maps for CreationErrorReason.
var (
	CreationErrorReason_name = map[int32]string{
		0:  "UNKNOWN_ERROR",
		1:  "GET_ARTICLE_DRAFT_FAILED",
		2:  "GET_ARTICLE_LIST_FAILED",
		3:  "GET_DRAFT_LIST_FAILED",
		4:  "CREATE_ARTICLE_FAILED",
		5:  "CREATE_ARTICLE_DRAFT_FAILED",
		6:  "CREATE_ARTICLE_FOLDER_FAILED",
		7:  "CREATE_ARTICLE_STATISTIC_FAILED",
		8:  "CREATE_ARTICLE_CACHE_FAILED",
		9:  "CREATE_ARTICLE_SEARCH_FAILED",
		10: "DELETE_ARTICLE_DRAFT_FAILED",
		11: "RECORD_NOT_FOUND",
		12: "DRAFT_MARK_FAILED",
		13: "SEND_ARTICLE_FAILED",
		14: "SEND_TO_MQ_FAILED",
	}
	CreationErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR":                   0,
		"GET_ARTICLE_DRAFT_FAILED":        1,
		"GET_ARTICLE_LIST_FAILED":         2,
		"GET_DRAFT_LIST_FAILED":           3,
		"CREATE_ARTICLE_FAILED":           4,
		"CREATE_ARTICLE_DRAFT_FAILED":     5,
		"CREATE_ARTICLE_FOLDER_FAILED":    6,
		"CREATE_ARTICLE_STATISTIC_FAILED": 7,
		"CREATE_ARTICLE_CACHE_FAILED":     8,
		"CREATE_ARTICLE_SEARCH_FAILED":    9,
		"DELETE_ARTICLE_DRAFT_FAILED":     10,
		"RECORD_NOT_FOUND":                11,
		"DRAFT_MARK_FAILED":               12,
		"SEND_ARTICLE_FAILED":             13,
		"SEND_TO_MQ_FAILED":               14,
	}
)

func (x CreationErrorReason) Enum() *CreationErrorReason {
	p := new(CreationErrorReason)
	*p = x
	return p
}

func (x CreationErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreationErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_creation_service_v1_creation_error_proto_enumTypes[0].Descriptor()
}

func (CreationErrorReason) Type() protoreflect.EnumType {
	return &file_creation_service_v1_creation_error_proto_enumTypes[0]
}

func (x CreationErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreationErrorReason.Descriptor instead.
func (CreationErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_error_proto_rawDescGZIP(), []int{0}
}

var File_creation_service_v1_creation_error_proto protoreflect.FileDescriptor

var file_creation_service_v1_creation_error_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc8, 0x03, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x45, 0x54, 0x5f, 0x41,
	0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x52, 0x54,
	0x49, 0x43, 0x4c, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x45, 0x54, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f,
	0x4c, 0x49, 0x53, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x19, 0x0a,
	0x15, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x4c, 0x44,
	0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x23, 0x0a, 0x1f, 0x43,
	0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x49, 0x53, 0x54, 0x49, 0x43, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07,
	0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43,
	0x4c, 0x45, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x08, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x52, 0x54, 0x49,
	0x43, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x09, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x41, 0x52,
	0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x52,
	0x41, 0x46, 0x54, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x0c, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45,
	0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x4d, 0x51, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x0e, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x62, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creation_service_v1_creation_error_proto_rawDescOnce sync.Once
	file_creation_service_v1_creation_error_proto_rawDescData = file_creation_service_v1_creation_error_proto_rawDesc
)

func file_creation_service_v1_creation_error_proto_rawDescGZIP() []byte {
	file_creation_service_v1_creation_error_proto_rawDescOnce.Do(func() {
		file_creation_service_v1_creation_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_creation_service_v1_creation_error_proto_rawDescData)
	})
	return file_creation_service_v1_creation_error_proto_rawDescData
}

var file_creation_service_v1_creation_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_creation_service_v1_creation_error_proto_goTypes = []interface{}{
	(CreationErrorReason)(0), // 0: creation.v1.CreationErrorReason
}
var file_creation_service_v1_creation_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_creation_service_v1_creation_error_proto_init() }
func file_creation_service_v1_creation_error_proto_init() {
	if File_creation_service_v1_creation_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_creation_service_v1_creation_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creation_service_v1_creation_error_proto_goTypes,
		DependencyIndexes: file_creation_service_v1_creation_error_proto_depIdxs,
		EnumInfos:         file_creation_service_v1_creation_error_proto_enumTypes,
	}.Build()
	File_creation_service_v1_creation_error_proto = out.File
	file_creation_service_v1_creation_error_proto_rawDesc = nil
	file_creation_service_v1_creation_error_proto_goTypes = nil
	file_creation_service_v1_creation_error_proto_depIdxs = nil
}

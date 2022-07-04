// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: creation/service/v1/creation.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetArticleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetArticleListReq) Reset() {
	*x = GetArticleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleListReq) ProtoMessage() {}

func (x *GetArticleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleListReq.ProtoReflect.Descriptor instead.
func (*GetArticleListReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{0}
}

func (x *GetArticleListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetArticleListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article []*GetArticleListReply_Article `protobuf:"bytes,1,rep,name=article,proto3" json:"article,omitempty"`
}

func (x *GetArticleListReply) Reset() {
	*x = GetArticleListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleListReply) ProtoMessage() {}

func (x *GetArticleListReply) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleListReply.ProtoReflect.Descriptor instead.
func (*GetArticleListReply) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{1}
}

func (x *GetArticleListReply) GetArticle() []*GetArticleListReply_Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type GetLastArticleDraftReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetLastArticleDraftReq) Reset() {
	*x = GetLastArticleDraftReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastArticleDraftReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastArticleDraftReq) ProtoMessage() {}

func (x *GetLastArticleDraftReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastArticleDraftReq.ProtoReflect.Descriptor instead.
func (*GetLastArticleDraftReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{2}
}

func (x *GetLastArticleDraftReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetLastArticleDraftReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetLastArticleDraftReply) Reset() {
	*x = GetLastArticleDraftReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastArticleDraftReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastArticleDraftReply) ProtoMessage() {}

func (x *GetLastArticleDraftReply) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastArticleDraftReply.ProtoReflect.Descriptor instead.
func (*GetLastArticleDraftReply) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{3}
}

func (x *GetLastArticleDraftReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetLastArticleDraftReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CreateArticleReq) Reset() {
	*x = CreateArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleReq) ProtoMessage() {}

func (x *CreateArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleReq.ProtoReflect.Descriptor instead.
func (*CreateArticleReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{4}
}

func (x *CreateArticleReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateArticleReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateArticleCacheAndSearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CreateArticleCacheAndSearchReq) Reset() {
	*x = CreateArticleCacheAndSearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleCacheAndSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleCacheAndSearchReq) ProtoMessage() {}

func (x *CreateArticleCacheAndSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleCacheAndSearchReq.ProtoReflect.Descriptor instead.
func (*CreateArticleCacheAndSearchReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{5}
}

func (x *CreateArticleCacheAndSearchReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateArticleCacheAndSearchReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateArticleDraftReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CreateArticleDraftReq) Reset() {
	*x = CreateArticleDraftReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleDraftReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleDraftReq) ProtoMessage() {}

func (x *CreateArticleDraftReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleDraftReq.ProtoReflect.Descriptor instead.
func (*CreateArticleDraftReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{6}
}

func (x *CreateArticleDraftReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateArticleDraftReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateArticleDraftReply) Reset() {
	*x = CreateArticleDraftReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleDraftReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleDraftReply) ProtoMessage() {}

func (x *CreateArticleDraftReply) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleDraftReply.ProtoReflect.Descriptor instead.
func (*CreateArticleDraftReply) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{7}
}

func (x *CreateArticleDraftReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ArticleDraftMarkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ArticleDraftMarkReq) Reset() {
	*x = ArticleDraftMarkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleDraftMarkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleDraftMarkReq) ProtoMessage() {}

func (x *ArticleDraftMarkReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleDraftMarkReq.ProtoReflect.Descriptor instead.
func (*ArticleDraftMarkReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{8}
}

func (x *ArticleDraftMarkReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleDraftMarkReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetArticleDraftListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetArticleDraftListReq) Reset() {
	*x = GetArticleDraftListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleDraftListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleDraftListReq) ProtoMessage() {}

func (x *GetArticleDraftListReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleDraftListReq.ProtoReflect.Descriptor instead.
func (*GetArticleDraftListReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{9}
}

func (x *GetArticleDraftListReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetArticleDraftListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Draft []*GetArticleDraftListReply_Draft `protobuf:"bytes,1,rep,name=draft,proto3" json:"draft,omitempty"`
}

func (x *GetArticleDraftListReply) Reset() {
	*x = GetArticleDraftListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleDraftListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleDraftListReply) ProtoMessage() {}

func (x *GetArticleDraftListReply) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleDraftListReply.ProtoReflect.Descriptor instead.
func (*GetArticleDraftListReply) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{10}
}

func (x *GetArticleDraftListReply) GetDraft() []*GetArticleDraftListReply_Draft {
	if x != nil {
		return x.Draft
	}
	return nil
}

type SendArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *SendArticleReq) Reset() {
	*x = SendArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendArticleReq) ProtoMessage() {}

func (x *SendArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendArticleReq.ProtoReflect.Descriptor instead.
func (*SendArticleReq) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{11}
}

func (x *SendArticleReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SendArticleReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetArticleListReply_Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetArticleListReply_Article) Reset() {
	*x = GetArticleListReply_Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleListReply_Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleListReply_Article) ProtoMessage() {}

func (x *GetArticleListReply_Article) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleListReply_Article.ProtoReflect.Descriptor instead.
func (*GetArticleListReply_Article) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetArticleListReply_Article) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetArticleListReply_Article) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetArticleDraftListReply_Draft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetArticleDraftListReply_Draft) Reset() {
	*x = GetArticleDraftListReply_Draft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creation_service_v1_creation_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleDraftListReply_Draft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleDraftListReply_Draft) ProtoMessage() {}

func (x *GetArticleDraftListReply_Draft) ProtoReflect() protoreflect.Message {
	mi := &file_creation_service_v1_creation_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleDraftListReply_Draft.ProtoReflect.Descriptor instead.
func (*GetArticleDraftListReply_Draft) Descriptor() ([]byte, []int) {
	return file_creation_service_v1_creation_proto_rawDescGZIP(), []int{10, 0}
}

func (x *GetArticleDraftListReply_Draft) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_creation_service_v1_creation_proto protoreflect.FileDescriptor

var file_creation_service_v1_creation_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x88,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x2d, 0x0a, 0x07, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x42, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x40, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x41, 0x6e, 0x64,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x29,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44,
	0x72, 0x61, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x13, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x36,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x76, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x05,
	0x64, 0x72, 0x61, 0x66, 0x74, 0x1a, 0x17, 0x0a, 0x05, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e,
	0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0xee,
	0x05, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x70, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x63, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44,
	0x72, 0x61, 0x66, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2b, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x41, 0x6e, 0x64, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x60, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x10, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44,
	0x72, 0x61, 0x66, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x20, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72,
	0x61, 0x66, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x25, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x1c, 0x5a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creation_service_v1_creation_proto_rawDescOnce sync.Once
	file_creation_service_v1_creation_proto_rawDescData = file_creation_service_v1_creation_proto_rawDesc
)

func file_creation_service_v1_creation_proto_rawDescGZIP() []byte {
	file_creation_service_v1_creation_proto_rawDescOnce.Do(func() {
		file_creation_service_v1_creation_proto_rawDescData = protoimpl.X.CompressGZIP(file_creation_service_v1_creation_proto_rawDescData)
	})
	return file_creation_service_v1_creation_proto_rawDescData
}

var file_creation_service_v1_creation_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_creation_service_v1_creation_proto_goTypes = []interface{}{
	(*GetArticleListReq)(nil),              // 0: creation.v1.GetArticleListReq
	(*GetArticleListReply)(nil),            // 1: creation.v1.GetArticleListReply
	(*GetLastArticleDraftReq)(nil),         // 2: creation.v1.GetLastArticleDraftReq
	(*GetLastArticleDraftReply)(nil),       // 3: creation.v1.GetLastArticleDraftReply
	(*CreateArticleReq)(nil),               // 4: creation.v1.CreateArticleReq
	(*CreateArticleCacheAndSearchReq)(nil), // 5: creation.v1.CreateArticleCacheAndSearchReq
	(*CreateArticleDraftReq)(nil),          // 6: creation.v1.CreateArticleDraftReq
	(*CreateArticleDraftReply)(nil),        // 7: creation.v1.CreateArticleDraftReply
	(*ArticleDraftMarkReq)(nil),            // 8: creation.v1.ArticleDraftMarkReq
	(*GetArticleDraftListReq)(nil),         // 9: creation.v1.GetArticleDraftListReq
	(*GetArticleDraftListReply)(nil),       // 10: creation.v1.GetArticleDraftListReply
	(*SendArticleReq)(nil),                 // 11: creation.v1.SendArticleReq
	(*GetArticleListReply_Article)(nil),    // 12: creation.v1.GetArticleListReply.Article
	(*GetArticleDraftListReply_Draft)(nil), // 13: creation.v1.GetArticleDraftListReply.Draft
	(*emptypb.Empty)(nil),                  // 14: google.protobuf.Empty
}
var file_creation_service_v1_creation_proto_depIdxs = []int32{
	12, // 0: creation.v1.GetArticleListReply.article:type_name -> creation.v1.GetArticleListReply.Article
	13, // 1: creation.v1.GetArticleDraftListReply.draft:type_name -> creation.v1.GetArticleDraftListReply.Draft
	0,  // 2: creation.v1.Creation.GetArticleList:input_type -> creation.v1.GetArticleListReq
	2,  // 3: creation.v1.Creation.GetLastArticleDraft:input_type -> creation.v1.GetLastArticleDraftReq
	4,  // 4: creation.v1.Creation.CreateArticle:input_type -> creation.v1.CreateArticleReq
	5,  // 5: creation.v1.Creation.CreateArticleCacheAndSearch:input_type -> creation.v1.CreateArticleCacheAndSearchReq
	6,  // 6: creation.v1.Creation.CreateArticleDraft:input_type -> creation.v1.CreateArticleDraftReq
	8,  // 7: creation.v1.Creation.ArticleDraftMark:input_type -> creation.v1.ArticleDraftMarkReq
	9,  // 8: creation.v1.Creation.GetArticleDraftList:input_type -> creation.v1.GetArticleDraftListReq
	11, // 9: creation.v1.Creation.SendArticle:input_type -> creation.v1.SendArticleReq
	1,  // 10: creation.v1.Creation.GetArticleList:output_type -> creation.v1.GetArticleListReply
	3,  // 11: creation.v1.Creation.GetLastArticleDraft:output_type -> creation.v1.GetLastArticleDraftReply
	14, // 12: creation.v1.Creation.CreateArticle:output_type -> google.protobuf.Empty
	14, // 13: creation.v1.Creation.CreateArticleCacheAndSearch:output_type -> google.protobuf.Empty
	7,  // 14: creation.v1.Creation.CreateArticleDraft:output_type -> creation.v1.CreateArticleDraftReply
	14, // 15: creation.v1.Creation.ArticleDraftMark:output_type -> google.protobuf.Empty
	10, // 16: creation.v1.Creation.GetArticleDraftList:output_type -> creation.v1.GetArticleDraftListReply
	14, // 17: creation.v1.Creation.SendArticle:output_type -> google.protobuf.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_creation_service_v1_creation_proto_init() }
func file_creation_service_v1_creation_proto_init() {
	if File_creation_service_v1_creation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_creation_service_v1_creation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleListReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleListReply); i {
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
		file_creation_service_v1_creation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastArticleDraftReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastArticleDraftReply); i {
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
		file_creation_service_v1_creation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleCacheAndSearchReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleDraftReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleDraftReply); i {
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
		file_creation_service_v1_creation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleDraftMarkReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleDraftListReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleDraftListReply); i {
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
		file_creation_service_v1_creation_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendArticleReq); i {
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
		file_creation_service_v1_creation_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleListReply_Article); i {
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
		file_creation_service_v1_creation_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleDraftListReply_Draft); i {
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
			RawDescriptor: file_creation_service_v1_creation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_creation_service_v1_creation_proto_goTypes,
		DependencyIndexes: file_creation_service_v1_creation_proto_depIdxs,
		MessageInfos:      file_creation_service_v1_creation_proto_msgTypes,
	}.Build()
	File_creation_service_v1_creation_proto = out.File
	file_creation_service_v1_creation_proto_rawDesc = nil
	file_creation_service_v1_creation_proto_goTypes = nil
	file_creation_service_v1_creation_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: borrow.protobuf

package borrow_server

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy protobuf package is being used.
const _ = proto.ProtoPackageIsVersion4

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author       string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Publish      string `protobuf:"bytes,4,opt,name=publish,proto3" json:"publish,omitempty"`
	Introduction string `protobuf:"bytes,5,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Number       int64  `protobuf:"varint,6,opt,name=number,proto3" json:"number,omitempty"`
	ISBN         string `protobuf:"bytes,7,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_borrow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_borrow_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPublish() string {
	if x != nil {
		return x.Publish
	}
	return ""
}

func (x *Book) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *Book) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Book) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

type GetBookListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBookListReq) Reset() {
	*x = GetBookListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListReq) ProtoMessage() {}

func (x *GetBookListReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListReq.ProtoReflect.Descriptor instead.
func (*GetBookListReq) Descriptor() ([]byte, []int) {
	return file_borrow_proto_rawDescGZIP(), []int{1}
}

type GetBookListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBookListRsp) Reset() {
	*x = GetBookListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListRsp) ProtoMessage() {}

func (x *GetBookListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_borrow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListRsp.ProtoReflect.Descriptor instead.
func (*GetBookListRsp) Descriptor() ([]byte, []int) {
	return file_borrow_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookListRsp) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type GetBookDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookDetailReq) Reset() {
	*x = GetBookDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookDetailReq) ProtoMessage() {}

func (x *GetBookDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookDetailReq.ProtoReflect.Descriptor instead.
func (*GetBookDetailReq) Descriptor() ([]byte, []int) {
	return file_borrow_proto_rawDescGZIP(), []int{3}
}

func (x *GetBookDetailReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBookDetailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookDetailRsp) Reset() {
	*x = GetBookDetailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookDetailRsp) ProtoMessage() {}

func (x *GetBookDetailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_borrow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookDetailRsp.ProtoReflect.Descriptor instead.
func (*GetBookDetailRsp) Descriptor() ([]byte, []int) {
	return file_borrow_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookDetailRsp) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type BorrowBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId     int64  `protobuf:"varint,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BorrowTime string `protobuf:"bytes,3,opt,name=borrow_time,json=borrowTime,proto3" json:"borrow_time,omitempty"`
}

func (x *BorrowBookReq) Reset() {
	*x = BorrowBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookReq) ProtoMessage() {}

func (x *BorrowBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookReq.ProtoReflect.Descriptor instead.
func (*BorrowBookReq) Descriptor() ([]byte, []int) {
	return file_borrow_proto_rawDescGZIP(), []int{5}
}

func (x *BorrowBookReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BorrowBookReq) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *BorrowBookReq) GetBorrowTime() string {
	if x != nil {
		return x.BorrowTime
	}
	return ""
}

type BorrowBookRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BorrowBookRsp) Reset() {
	*x = BorrowBookRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowBookRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookRsp) ProtoMessage() {}

func (x *BorrowBookRsp) ProtoReflect() protoreflect.Message {
	mi := &file_borrow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookRsp.ProtoReflect.Descriptor instead.
func (*BorrowBookRsp) Descriptor() ([]byte, []int) {
	return file_borrow_proto_rawDescGZIP(), []int{6}
}

var File_borrow_proto protoreflect.FileDescriptor

var file_borrow_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xac, 0x01,
	0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x53, 0x42, 0x4e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x22, 0x10, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x3b,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x12, 0x29, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x62, 0x0a, 0x0d,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x0f, 0x0a, 0x0d, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73,
	0x70, 0x32, 0xf8, 0x01, 0x0a, 0x06, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x12, 0x4d, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c,
	0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x3b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_borrow_proto_rawDescOnce sync.Once
	file_borrow_proto_rawDescData = file_borrow_proto_rawDesc
)

func file_borrow_proto_rawDescGZIP() []byte {
	file_borrow_proto_rawDescOnce.Do(func() {
		file_borrow_proto_rawDescData = protoimpl.X.CompressGZIP(file_borrow_proto_rawDescData)
	})
	return file_borrow_proto_rawDescData
}

var file_borrow_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_borrow_proto_goTypes = []interface{}{
	(*Book)(nil),             // 0: borrow_server.Book
	(*GetBookListReq)(nil),   // 1: borrow_server.GetBookListReq
	(*GetBookListRsp)(nil),   // 2: borrow_server.GetBookListRsp
	(*GetBookDetailReq)(nil), // 3: borrow_server.GetBookDetailReq
	(*GetBookDetailRsp)(nil), // 4: borrow_server.GetBookDetailRsp
	(*BorrowBookReq)(nil),    // 5: borrow_server.BorrowBookReq
	(*BorrowBookRsp)(nil),    // 6: borrow_server.BorrowBookRsp
}
var file_borrow_proto_depIdxs = []int32{
	0, // 0: borrow_server.GetBookListRsp.books:type_name -> borrow_server.Book
	0, // 1: borrow_server.GetBookDetailRsp.book:type_name -> borrow_server.Book
	1, // 2: borrow_server.Borrow.GetBookList:input_type -> borrow_server.GetBookListReq
	3, // 3: borrow_server.Borrow.GetBookDetail:input_type -> borrow_server.GetBookDetailReq
	5, // 4: borrow_server.Borrow.BorrowBook:input_type -> borrow_server.BorrowBookReq
	2, // 5: borrow_server.Borrow.GetBookList:output_type -> borrow_server.GetBookListRsp
	4, // 6: borrow_server.Borrow.GetBookDetail:output_type -> borrow_server.GetBookDetailRsp
	6, // 7: borrow_server.Borrow.BorrowBook:output_type -> borrow_server.BorrowBookRsp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_borrow_proto_init() }
func file_borrow_proto_init() {
	if File_borrow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_borrow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_borrow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookListReq); i {
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
		file_borrow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookListRsp); i {
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
		file_borrow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookDetailReq); i {
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
		file_borrow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookDetailRsp); i {
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
		file_borrow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowBookReq); i {
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
		file_borrow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowBookRsp); i {
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
			RawDescriptor: file_borrow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_borrow_proto_goTypes,
		DependencyIndexes: file_borrow_proto_depIdxs,
		MessageInfos:      file_borrow_proto_msgTypes,
	}.Build()
	File_borrow_proto = out.File
	file_borrow_proto_rawDesc = nil
	file_borrow_proto_goTypes = nil
	file_borrow_proto_depIdxs = nil
}
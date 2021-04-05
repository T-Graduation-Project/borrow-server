// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: protocol/borrow.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Borrow service

func NewBorrowEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Borrow service

type BorrowService interface {
	BorrowBook(ctx context.Context, in *BorrowBookReq, opts ...client.CallOption) (*BorrowBookRsp, error)
	ReturnBook(ctx context.Context, in *ReturnBookReq, opts ...client.CallOption) (*ReturnBookRsp, error)
	DeleteRecord(ctx context.Context, in *DeleteRecordReq, opts ...client.CallOption) (*DeleteRecordRsp, error)
}

type borrowService struct {
	c    client.Client
	name string
}

func NewBorrowService(name string, c client.Client) BorrowService {
	return &borrowService{
		c:    c,
		name: name,
	}
}

func (c *borrowService) BorrowBook(ctx context.Context, in *BorrowBookReq, opts ...client.CallOption) (*BorrowBookRsp, error) {
	req := c.c.NewRequest(c.name, "Borrow.BorrowBook", in)
	out := new(BorrowBookRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowService) ReturnBook(ctx context.Context, in *ReturnBookReq, opts ...client.CallOption) (*ReturnBookRsp, error) {
	req := c.c.NewRequest(c.name, "Borrow.ReturnBook", in)
	out := new(ReturnBookRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowService) DeleteRecord(ctx context.Context, in *DeleteRecordReq, opts ...client.CallOption) (*DeleteRecordRsp, error) {
	req := c.c.NewRequest(c.name, "Borrow.DeleteRecord", in)
	out := new(DeleteRecordRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Borrow service

type BorrowHandler interface {
	BorrowBook(context.Context, *BorrowBookReq, *BorrowBookRsp) error
	ReturnBook(context.Context, *ReturnBookReq, *ReturnBookRsp) error
	DeleteRecord(context.Context, *DeleteRecordReq, *DeleteRecordRsp) error
}

func RegisterBorrowHandler(s server.Server, hdlr BorrowHandler, opts ...server.HandlerOption) error {
	type borrow interface {
		BorrowBook(ctx context.Context, in *BorrowBookReq, out *BorrowBookRsp) error
		ReturnBook(ctx context.Context, in *ReturnBookReq, out *ReturnBookRsp) error
		DeleteRecord(ctx context.Context, in *DeleteRecordReq, out *DeleteRecordRsp) error
	}
	type Borrow struct {
		borrow
	}
	h := &borrowHandler{hdlr}
	return s.Handle(s.NewHandler(&Borrow{h}, opts...))
}

type borrowHandler struct {
	BorrowHandler
}

func (h *borrowHandler) BorrowBook(ctx context.Context, in *BorrowBookReq, out *BorrowBookRsp) error {
	return h.BorrowHandler.BorrowBook(ctx, in, out)
}

func (h *borrowHandler) ReturnBook(ctx context.Context, in *ReturnBookReq, out *ReturnBookRsp) error {
	return h.BorrowHandler.ReturnBook(ctx, in, out)
}

func (h *borrowHandler) DeleteRecord(ctx context.Context, in *DeleteRecordReq, out *DeleteRecordRsp) error {
	return h.BorrowHandler.DeleteRecord(ctx, in, out)
}

package api

import (
	"context"
	"github.com/T-Graduation-Project/borrow-server/app/service"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
)

type borrowService struct{}

var Borrow = &borrowService{}

// 借书
func (s *borrowService) BorrowBook(
	ctx context.Context, r *protobuf.BorrowBookReq) (*protobuf.BorrowBookRsp, error) {
	rsp, err := service.BorrowBook(r)
	return rsp, err
}

// 还书
func (s *borrowService) ReturnBook(
	ctx context.Context, r *protobuf.ReturnBookReq) (*protobuf.ReturnBookRsp, error) {
	rsp, err := service.ReturnBook(r)
	return rsp, err
}

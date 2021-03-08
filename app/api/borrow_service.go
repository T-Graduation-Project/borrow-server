package api

import (
	"github.com/T-Graduation-Project/borrow-server/app/service"
	"github.com/T-Graduation-Project/borrow-server/lib/response"
	proto "github.com/T-Graduation-Project/protobuf/borrow_server"
	"github.com/gogf/gf/net/ghttp"
)

type borrowService struct{}

var Borrow = new(borrowService)

// 获取图书列表
func (s *borrowService) GetBookList(r *ghttp.Request) {
	rsp, err := service.GetBookList()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, rsp)
}

func (s *borrowService) BorrowBook(r *ghttp.Request) {
	var req proto.BorrowBookReq
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	rsp, err := service.BorrowBook(req)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, rsp)
}

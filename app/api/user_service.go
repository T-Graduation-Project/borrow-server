package api

import (
	"github.com/T-Graduation-Project/borrow-server/app/service"
	"github.com/T-Graduation-Project/borrow-server/lib/response"
	"github.com/gogf/gf/net/ghttp"
)

type userService struct{}

var User = new(userService)


func (s *userService) GetUserInfo(r *ghttp.Request) {
	//rsp, err := service.GetBookList()
	//if err != nil {
	//	response.JsonExit(r, 1, err.Error())
	//}
	//response.JsonExit(r, 0, rsp)
}


func (s *userService) GetUserBorrowInfo(r *ghttp.Request) {
	rsp, err := service.GetBookList()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, rsp)
}
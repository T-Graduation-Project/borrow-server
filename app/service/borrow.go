package service

import (
	"github.com/T-Graduation-Project/borrow-server/app/dao"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

func GetBookList() (*borrow_server.GetBookListRsp, error) {
	rsp := new(borrow_server.GetBookListRsp)
	data := dao.SelectBook("default", "")
	dataMap := g.Map{
		"Data": data,
	}
	if err := gconv.Struct(dataMap, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func BorrowBook(req borrow_server.BorrowBookReq) (*borrow_server.BorrowBookRsp, error) {
	rsp := new(borrow_server.BorrowBookRsp)
	borrowInfo := g.Map{
		"book_id":   req.BookId,
		"user_id":   req.UserId,
		"lend_date": time.Now(),
	}
	err := dao.SaveBorrowList("default", borrowInfo)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

package service

import (
	"github.com/T-Graduation-Project/borrow-server/app/dao"
	proto "github.com/T-Graduation-Project/proto/borrow_server"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

func GetBookList() (*proto.GetBookListRsp, error) {
	rsp := new(proto.GetBookListRsp)
	data := dao.SelectBook("default", "")
	dataMap := g.Map{
		"Data": data,
	}
	if err := gconv.Struct(dataMap, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

func BorrowBook(req proto.BorrowBookReq) (*proto.BorrowBookRsp, error) {
	rsp := new(proto.BorrowBookRsp)
	borrowInfo := g.Map{
		"book_id":   req.BookId,
		"user_id":   req.UserId,
		"lend_date": time.Now(),
	}
	err := dao.SaveBorroList("default", borrowInfo)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

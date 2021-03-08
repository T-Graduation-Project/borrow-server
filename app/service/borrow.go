package service

import (
	"fmt"
	"github.com/T-Graduation-Project/borrow-server/app/dao"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"time"
)

func GetBookList() (*protobuf.GetBookListRsp, error) {
	rsp := new(protobuf.GetBookListRsp)
	db := g.DB("default")
	err := db.Table("book_info").Scan(&rsp.Books)
	return rsp, err
}

func BorrowBook(r *protobuf.BorrowBookReq) (*protobuf.BorrowBookRsp, error) {
	rsp := new(protobuf.BorrowBookRsp)
	borrowInfo := g.Map{
		"book_id":   r.BookId,
		"user_id":   r.UserId,
		"lend_date": time.Now(),
	}
	err := dao.SaveBorrowList("default", borrowInfo)
	rsp.Msg = "success"
	return rsp, err
}

func ReturnBook(r *protobuf.ReturnBookReq) (*protobuf.ReturnBookRsp, error) {
	rsp := new(protobuf.ReturnBookRsp)
	db := g.DB("default")
	res, err := db.Table("borrow_list").Data(
		g.Map{"back_date": gtime.Datetime()}).Where(
		"book_id=? AND user_id=?", r.BookId, r.UserId).Update()
	fmt.Println(res)
	rsp.Msg = "success"
	return rsp, err
}

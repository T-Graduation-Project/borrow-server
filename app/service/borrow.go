package service

import (
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
	// 检查是否有重复的借阅记录
	hasRepeat, err := hasRepeatBorrowRecord(r.UserId, r.BookId)
	if err != nil || hasRepeat == true {
		return rsp, err
	}
	// 写入借阅记录
	borrowInfo := g.Map{
		"book_id":   r.BookId,
		"user_id":   r.UserId,
		"lend_date": time.Now(),
	}
	db := g.DB("default")
	_, err = db.Table("borrow_list").Save(borrowInfo)
	if err != nil {
		g.Log().Println("SaveBorrowList failed")
	}
	// 修改 book_info 中的书本库存数量
	book, err := db.Table("book_info").Where("id=?", r.BookId).One()
	num := book.Map()["number"].(int)
	if num == 0 {
		rsp.Msg = "The number of this book is 0"
		return rsp, err
	}
	_, err = db.Table("book_info").Update(
		g.Map{"number": num - 1}, "id", r.BookId)
	rsp.Msg = "borrow success"
	return rsp, err
}

func ReturnBook(r *protobuf.ReturnBookReq) (*protobuf.ReturnBookRsp, error) {
	rsp := new(protobuf.ReturnBookRsp)
	db := g.DB("default")
	res, err := db.Table("borrow_list").Data(
		g.Map{"back_date": gtime.Datetime()}).Where(
		"book_id=? AND user_id=?", r.BookId, r.UserId).Update()
	g.Log().Println(res)

	// 修改 book_info 中的书本库存数量
	book, err := db.Table("book_info").Where("id=?", r.BookId).One()
	_, err = db.Table("book_info").Update(
		g.Map{"number": book.Map()["number"].(int) + 1}, "id", r.BookId)

	rsp.Msg = "success"
	return rsp, err
}

// 检查借阅记录，不允许同一个人重复借相同的书
func hasRepeatBorrowRecord(userId int64, bookId int64) (bool, error) {
	res, err := g.DB("default").Table("borrow_list").Where(
		"book_id=? AND user_id=?", bookId, userId).All()
	g.Log().Println("repeat record:", res)
	if err != nil || res != nil {
		return true, err
	}
	return false, err
}

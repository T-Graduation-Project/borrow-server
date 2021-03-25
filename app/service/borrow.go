package service

import (
	"context"
	bookPb "github.com/T-Graduation-Project/book-server/protobuf"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"google.golang.org/grpc"
)

var (
	db  = g.DB("default")
	log = glog.New()
)

func BorrowBook(r *protobuf.BorrowBookReq) (*protobuf.BorrowBookRsp, error) {
	rsp := &protobuf.BorrowBookRsp{
		Code: 1,
		Msg:  "borrow book error",
	}
	// 检查是否有重复的借阅记录
	hasRepeat, err := hasRepeatBorrowRecord(r.UserId, r.BookId)
	if err != nil || hasRepeat == true {
		log.Info("has repeat record:", err)
		return rsp, err
	}
	// 写入借阅记录
	borrowInfo := g.Map{
		"book_id":   r.BookId,
		"user_id":   r.UserId,
		"lend_date": gtime.Date(),
	}
	_, err = db.Table("borrow_list").Save(borrowInfo)
	if err != nil {
		log.Info("SaveBorrowList failed")
	}

	// book_info 中的书本库存数量 -1
	book, err := db.Table("book_info").Where("id=?", r.BookId).One()
	num := book.Map()["number"].(int)
	isbnCode := book.Map()["ISBN"].(string)
	if num == 0 {
		rsp.Msg = "The number of this book is 0"
		return rsp, err
	}
	updateBookNumber(isbnCode, num-1)

	rsp.Code = 0
	rsp.Msg = "borrow success"
	return rsp, err
}

func ReturnBook(r *protobuf.ReturnBookReq) (*protobuf.ReturnBookRsp, error) {
	rsp := &protobuf.ReturnBookRsp{
		Code: 1,
		Msg:  "return book error",
	}
	// 检查是否存在借阅记录，存在且状态为未归还才可以还书
	can, err := canBeBack(r.UserId, r.BookId, r.BorrowDate)
	if err != nil {
		log.Fatal("err:", err)
		return rsp, err
	}
	if can == false {
		rsp.Msg = "this book has been back"
		return rsp, err
	}
	// 更新 borrow_list
	res, err := db.Table("borrow_list").Data(
		g.Map{"back_date": gtime.Date()}).Where(
		"book_id=? AND user_id=?", r.BookId, r.UserId).Update()
	if err != nil {
		log.Info("update borrow_list error:", err)
		return rsp, err
	}
	log.Info(res)

	// 修改 book_info 中的书本库存数量
	book, err := db.Table("book_info").Where("id=?", r.BookId).One()
	num := book.Map()["number"].(int)
	isbnCode := book.Map()["ISBN"].(string)
	updateBookNumber(isbnCode, num+1)

	rsp.Code = 0
	rsp.Msg = "success"
	return rsp, err
}

func DeleteRecord(req *protobuf.DeleteRecordReq) (*protobuf.DeleteRecordRsp, error) {
	rsp := &protobuf.DeleteRecordRsp{
		Code: 1,
		Msg:  "Delete record err",
	}
	r, err := db.Table("book_info").Delete("id=?", req.Id)
	if err != nil {
		return rsp, err
	}
	log.Info("db info:", r)
	rsp.Code = 0
	rsp.Msg = "Delete record success"
	return rsp, nil
}

// 检查是否可以借阅，
// 检查借阅记录，不允许同一个人重复借相同的书
func hasRepeatBorrowRecord(userId int64, bookId int64) (bool, error) {
	res, err := db.Table("borrow_list").Where(
		"book_id=? AND user_id=?", bookId, userId).All()
	log.Info("repeat record:", res)
	if err != nil || res != nil {
		return true, err
	}
	return false, err
}

// 检查是否可归还：存在记录且未归还 back_date 为空
func canBeBack(userId int64, bookId int64, lendDate string) (bool, error) {
	record, err := db.Table("borrow_list").Where(
		"book_id=? AND user_id=? AND lend_date=?", bookId, userId, lendDate).One()
	log.Info("has record:", record)
	if err != nil {
		return false, err
	} else {
		if record.Map()["back_date"] != nil {
			log.Info("this book has been back:", record)
			return false, nil
		}
	}
	return true, err
}

func updateBookNumber(isbnCode string, bookNum int) {
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	bookClient := bookPb.NewBooksClient(conn)
	updateBooksReq := &bookPb.UpdateBooksReq{
		Books: []*bookPb.Book{
			{
				ISBN:   isbnCode,
				Number: gconv.Int64(bookNum),
			},
		},
	}
	bookClient.UpdateBooks(context.Background(), updateBooksReq)
}

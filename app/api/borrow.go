package api

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

type BorrowApi struct{}

var (
	db  = g.DB("default")
	log = glog.New()
)

func (s *BorrowApi) BorrowBook(
	ctx context.Context, req *protobuf.BorrowBookReq, rsp *protobuf.BorrowBookRsp) error {
	rsp.Code = 1
	rsp.Msg = "borrow book error"

	// 检查是否有重复的借阅记录
	can, err := canBeBorrowed(req.UserId, req.BookId)
	if err != nil {
		log.Info("can not be borrowed", err)
		return err
	}
	if can != true {
		log.Info("can not be borrowed")
		return err
	}
	// 写入借阅记录
	borrowInfo := g.Map{
		"book_id":   req.BookId,
		"user_id":   req.UserId,
		"lend_date": gtime.Date(),
	}
	_, err = db.Table("borrow_list").Save(borrowInfo)
	if err != nil {
		log.Info("SaveBorrowList failed")
	}

	// book_info 中的书本库存数量 -1
	book, err := db.Table("book_info").Where("id=?", req.BookId).One()
	num := book.Map()["number"].(int)
	isbnCode := book.Map()["ISBN"].(string)
	if num == 0 {
		rsp.Msg = "The number of this book is 0"
		return err
	}
	updateBookNumber(isbnCode, num-1)

	rsp.Code = 0
	rsp.Msg = "borrow book success"
	return err
}

func (s *BorrowApi) ReturnBook(
	ctx context.Context, req *protobuf.ReturnBookReq, rsp *protobuf.ReturnBookRsp) error {
	rsp.Code = 1
	rsp.Msg = "return book error"

	// 检查是否存在借阅记录，存在且状态为未归还才可以还书
	can, err := canBeBack(req.UserId, req.BookId, req.BorrowDate)
	if err != nil {
		log.Fatal("err:", err)
		return err
	}
	if can == false {
		rsp.Msg = "this book has been back"
		return err
	}
	// 更新 borrow_list
	res, err := db.Table("borrow_list").Data(
		g.Map{"back_date": gtime.Date()}).Where(
		"book_id=? AND user_id=?", req.BookId, req.UserId).Update()
	if err != nil {
		log.Info("update borrow_list error:", err)
		return err
	}
	log.Info(res)

	// 修改 book_info 中的书本库存数量
	book, err := db.Table("book_info").Where("id=?", req.BookId).One()
	num := book.Map()["number"].(int)
	isbnCode := book.Map()["ISBN"].(string)
	updateBookNumber(isbnCode, num+1)

	rsp.Code = 0
	rsp.Msg = "return book success"
	return err
}

func (s *BorrowApi) DeleteRecord(
	ctx context.Context, req *protobuf.DeleteRecordReq, rsp *protobuf.DeleteRecordRsp) error {
	rsp.Code = 1
	rsp.Msg = "Delete record err"

	r, err := db.Table("book_info").Delete("id=?", req.Id)
	if err != nil {
		return err
	}
	log.Info("db info:", r)
	rsp.Code = 0
	rsp.Msg = "Delete record success"
	return nil
}

// 检查是否可以借阅：
// 假设借书数量无限，同一本书最多借一本
// 借阅表中不存在 book_id，user_id 相同的借阅记录
// 若存在则应当 back_date 不为空（已还）
func canBeBorrowed(userId int64, bookId int64) (bool, error) {
	res, err := db.Table("borrow_list").Where(
		"book_id=? AND user_id=?", bookId, userId).All()
	if err != nil {
		return false, err
	}
	for _, record := range res {
		if record.Map()["back_date"] == nil {
			return false, nil
		}
	}
	return true, nil
}

// 检查是否可归还：存在记录且未归还，即 back_date 为空
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
	// book-server 的地址写死，之后优化
	conn, err := grpc.Dial("121.5.238.116:8001", grpc.WithInsecure())
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

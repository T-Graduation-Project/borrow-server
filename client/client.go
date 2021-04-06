package main

import (
	"context"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/micro/go-micro/v2"
)

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QiLCJQYXNzd29yZCI6InRlc3QiLCJleHAiOjE2MTc1MjgyODMsImlzcyI6InRva2VuX3NlcnZpY2UifQ.O3iYnyFfLWbpKLyO74LZbsQoQg-56zSlxwcSsm2BT5E"
)

var (
	log = g.Log()
)

func main() {
	service := micro.NewService(micro.Name("borrow.client"))
	service.Init()
	client := protobuf.NewBorrowService("borrow", service.Client())

	//tokenContext := metadata.NewOutgoingContext(
	//	context.Background(), metadata.New(map[string]string{"token": token}))
	// 借书
	r, err := client.BorrowBook(
		context.Background(), &protobuf.BorrowBookReq{UserId: 1, BookId: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("####### get server Greeting response: %s", r)

	// 还书
	r2, err := client.ReturnBook(
		context.Background(), &protobuf.ReturnBookReq{UserId: 1, BookId: 2, BorrowDate: gtime.Date()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("####### get server Greeting response: %s", r2)

	//r3, err := c.DeleteRecord(context.Background(), &protobuf.DeleteRecordReq{Id: 13})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r3)
}

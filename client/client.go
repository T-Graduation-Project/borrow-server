package main

import (
	"context"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

var (
	log = g.Log()
)

func main() {
	service := micro.NewService(
		micro.Name("borrow.client"),
		micro.Registry(etcdv3.NewRegistry(
			registry.Addrs(g.Cfg().GetString("registry_addr")),
		)),
	)
	client := protobuf.NewBorrowService("borrow", service.Client())

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

	// 查询借阅记录
	r3, err := client.GetBorrowHistory(
		context.Background(), &protobuf.GetBorrowHistoryReq{
			UserId: 1,
		})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info(r3.Records)

	//r3, err := c.DeleteRecord(context.Background(), &protobuf.DeleteRecordReq{Id: 13})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r3)
}

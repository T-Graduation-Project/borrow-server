package main

import (
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:8000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewBorrowClient(conn)

	// 借书
	//r, err := c.BorrowBook(context.Background(), &protobuf.BorrowBookReq{UserId: 1, BookId: 2})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r)

	// 还书
	r2, err := c.ReturnBook(
		context.Background(), &protobuf.ReturnBookReq{UserId: 1, BookId: 2, BorrowDate: "2021-03-25"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("####### get server Greeting response: %s", r2)

	//r3, err := c.DeleteRecord(context.Background(), &protobuf.DeleteRecordReq{Id: 13})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("####### get server Greeting response: %s", r3)
}

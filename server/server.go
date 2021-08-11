package main

import (
	"context"
	"fmt"
	"grpc-course/contact/contactpb"
	"log"
	"net"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type server struct{}

func (c *server) Insert(ctx context.Context, req *contactpb.InsertRequest) (*contactpb.InsertResponse, error) {

	log.Printf("calling insert %+v\n", req.Contact)
	ci := ConvertPbContact2ContactInfo(req.Contact)
	err := ci.Insert()

	if err != nil {
		resp := &contactpb.InsertResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("insert err %v", err),
		}

		return resp, nil
	}
	resp := &contactpb.InsertResponse{
		StatusCode: 1,
		Message:    "Ok",
	}
	return resp, nil
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	connectStr := "root:123456@tcp(127.0.0.1:3306)/contact?charset=utf8"
	maxIdle := 30
	maxConn := 30
	err := orm.RegisterDataBase("default", "mysql", connectStr, maxIdle, maxConn)
	if err != nil {
		log.Panicf("register db err: %v", err)
	}
	orm.RegisterModel(new(ContactInfo))

	err = orm.RunSyncdb("default", true, false)

	if err != nil {
		log.Panicf("run syncdb err: %v", err)
	}

	fmt.Println("connect db success")
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50070")

	if err != nil {
		log.Fatal("err while listening %v", err)
	}

	s := grpc.NewServer()

	contactpb.RegisterContactServiceServer(s, &server{})

	fmt.Println("Contact Service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Error while connecting to server", err)
	}
}

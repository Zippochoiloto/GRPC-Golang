package main

import (
	"context"
	"grpc-course/contact/contactpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50070", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Error when dialing: ", err)
	}

	defer cc.Close()

	client := contactpb.NewContactServiceClient(cc)

	insertContact(client, "09876", "Contact 1", "Address 1") // ko bi time out

}

func insertContact(cli contactpb.ContactServiceClient, phone, name, addr string) {
	req := &contactpb.InsertRequest{
		Contact: &contactpb.Contact{
			PhoneNumber: phone,
			Name:        name,
			Address:     addr,
		},
	}
	resp, err := cli.Insert(context.Background(), req)
	if err != nil {
		log.Printf("call insert err: %v", err)
		return
	}

	log.Printf("insert response %v", resp)
}

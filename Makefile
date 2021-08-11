gen-contact:
	protoc contactpb/contact.proto --go_out=plugins=grpc:.
	echo $GOPATH
run-server:
	go run server/server.go server/models.go
run-client:
	go run client/client.go
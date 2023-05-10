package main

import (
	"log"
	"net"
	pb "oa-review/proto/services"
	server "oa-review/user/rpc_server"

	"google.golang.org/grpc"
)

func main() {
	userServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
		return
	}

	pb.RegisterUserServiceServer(userServer, &server.UserService{})

	log.Println("User server successfully")
	userServer.Serve(lis)
}

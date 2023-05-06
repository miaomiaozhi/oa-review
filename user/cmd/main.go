package main

import (
	"log"
	"net"
	server "oa-review/user/rpc_server"
	pb "oa-review/user/services"

	"google.golang.org/grpc"
)

func main() {
	userServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	pb.RegisterUserServiceServer(userServer, &server.UserService{})

	log.Println("Success on user server")
	userServer.Serve(lis)
}

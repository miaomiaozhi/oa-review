package user

import (
	"log"
	"net"
	pb "oa-review/user/services"
	server "oa-review/user/rpc_server"
	"google.golang.org/grpc"
)

func Run() {
	userServer := grpc.NewServer()
	pb.RegisterUserServiceServer(userServer, &server.UserService{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

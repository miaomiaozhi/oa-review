package main

import (
	"log"
	"net"
	"oa-review/user/model"
	server "oa-review/user/rpc_server"
	pb "oa-review/user/services"
	_ "oa-review/user/conf"
	"google.golang.org/grpc"
)

func main() {
	// 初始化数据库
	if err := model.InitDataBase(); err != nil {
		panic("Error on init user DAO")
	}

	userServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	pb.RegisterUserServiceServer(userServer, &server.UserService{})

	log.Println("Success on user server")
	userServer.Serve(lis)
}

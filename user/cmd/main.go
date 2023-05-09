package main

import (
	"log"
	"net"
	pb "oa-review/proto/services"
	_ "oa-review/user/conf"
	"oa-review/user/model"
	server "oa-review/user/rpc_server"

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
		return
	}

	pb.RegisterUserServiceServer(userServer, &server.UserService{})

	log.Println("User server successfully")
	userServer.Serve(lis)
}

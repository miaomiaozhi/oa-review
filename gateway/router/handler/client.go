package handler

import (
	"log"

	pb "oa-review/gateway/services"

	"google.golang.org/grpc"
)

type grpcClient struct {
	conn       *grpc.ClientConn
	userClient pb.UserServiceClient
	// reviewClient pb.reviewClient
}

var C grpcClient

func init() {
	conn, err := newGrpcClient()
	if err != nil {
		return
	}
	C.conn = conn
	C.userClient = pb.NewUserServiceClient(conn)
	log.Println("Succees created grpc client ")
}

func (c *grpcClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func newGrpcClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Error on create grpc client")
		return nil, err
	}
	return conn, nil
}

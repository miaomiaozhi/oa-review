package rpcserver

import (
	"context"
	"fmt"
	services "oa-review/user/services"
)

type UserService struct {
}

func (userService *UserService) Login(ctx context.Context, req *services.UserLoginRequest) (*services.UserLoginResponse, error) {
	userId, userPsw := req.UserId, req.UserPassword

	return &services.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, user psw :%v", userId, userPsw),
		Token:      "login",
	}, nil
}

func (userService *UserService) Register(ctx context.Context, req *services.UserRegisterRequest) (*services.UserRegisterResponse, error) {
	return &services.UserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, user psw :%v", req.UserId, req.UserPassword),
		Token:      "register",
	}, nil
}

func (userService *UserService) GetInfo(ctx context.Context, req *services.UserGetInfoRequest) (*services.UserGetInfoResponse, error) {
	return &services.UserGetInfoResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, user psw :%v", req.UserId, req.UserPassword),
		User:       &services.User{},
	}, nil
}

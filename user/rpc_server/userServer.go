package rpcserver

import (
	"context"
	"fmt"
	services "oa-review/user/services"
)

type UserService struct {
	services.UnimplementedUserServiceServer
}

/*
req:
UserId string
UserPassword string
rep
*/
func (userService *UserService) Login(ctx context.Context, req *services.UserLoginRequest) (*services.UserLoginResponse, error) {
	userId, userPsw := req.UserId, req.UserPassword

	if userId == "" || userPsw == "" {
		return &services.UserLoginResponse{
			StatusCode: 400,
			StatusMsg:  "user id and user password can not be empty",
			Token:      "jwt token",
		}, nil
	}

	// DAO finding user
	return &services.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, user psw :%v", userId, userPsw),
		Token:      "jwt token",
	}, nil
}

/*
req:
userid string
userpassword string
priority int32
*/
func (userService *UserService) Register(ctx context.Context, req *services.UserRegisterRequest) (*services.UserRegisterResponse, error) {
	userId, userPsw := req.UserId, req.UserPassword

	if userId == "" || userPsw == "" {
		return &services.UserRegisterResponse{
			StatusCode: 400,
			StatusMsg:  "user id and user password can not be empty",
			Token:      "register",
		}, nil
	}

	// register store
	return &services.UserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, user psw :%v", req.UserId, req.UserPassword),
		Token:      "register",
	}, nil
}

/*
req:
userid int64
userpassword string
*/
func (userService *UserService) GetInfo(ctx context.Context, req *services.UserGetInfoRequest) (*services.UserGetInfoResponse, error) {
	userId, userPsw := req.UserId, req.UserPassword
	if userId < 0 || userPsw == "" {
		return &services.UserGetInfoResponse{
			StatusCode: 400,
			StatusMsg:  "user id/password illegal",
			User:       &services.User{},
		}, nil
	}

	// DAO

	return &services.UserGetInfoResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, user psw :%v", req.UserId, req.UserPassword),
		User:       &services.User{},
	}, nil
}

/*
req:
userid int64
app context string
*/
func (userService *UserService) SubmitApplication(ctx context.Context, req *services.UserSubmitApplicationRequest) (*services.UserSubmitApplicationResponse, error) {
	userId, appCtx := req.UserId, req.ApplicationContext
	if userId < 0 || appCtx == "" {
		return &services.UserSubmitApplicationResponse{
			StatusCode: 400,
			StatusMsg:  "user_id or app_context illegal",
		}, nil
	}

	// DAO

	return &services.UserSubmitApplicationResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user_id :%v, app_context :%v", userId, appCtx),
	}, nil
}

/*
req:
userid int64
appid int64
*/
func (userService *UserService) RetrievalApplication(ctx context.Context, req *services.UserRetrievalApplicationRequest) (*services.UserRetrievalApplicationResponse, error) {
	userId, appId := req.UserId, req.ApplicationId
	// Dao find user & application list
	if userId < 0 || appId < 0 {
		return &services.UserRetrievalApplicationResponse{
			StatusCode: 400,
			StatusMsg:  "user_id or app_id illegal",
		}, nil
	}

	// DAO

	return &services.UserRetrievalApplicationResponse{
		StatusCode:   200,
		StatusMsg:    fmt.Sprintf("user_id :%v, app_id :%v", userId, appId),
		ReviewStatue: false,
	}, nil
}

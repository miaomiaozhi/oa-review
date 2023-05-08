package rpcserver

import (
	"context"
	"fmt"
	"log"
	model "oa-review/user/model"
	services "oa-review/user/services"
	"strconv"
	"time"
)

type UserService struct {
	services.UnimplementedUserServiceServer
}

// tmp cache
var Users []*model.User
var Reviewers []*model.User
var AppList []*model.Application

func init() {
	log.Println("init user server tmp cache")
	Users = make([]*model.User, 0)
	Reviewers = make([]*model.User, 0)
	AppList = make([]*model.Application, 0)
}

/*
req:
UserId string
UserPassword string
rep
*/
func (userService *UserService) Login(ctx context.Context, req *services.UserLoginRequest) (*services.UserLoginResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserLoginResponse, error) {
		return &services.UserLoginResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
			Token:      "",
		}, nil
	}
	userIdStr, userPsw := req.UserId, req.UserPassword

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if userIdStr == "" || userPsw == "" || err != nil {
		return ErrResponse("error login user fmt")
	}

	userIdx, exist := CheckUserExist(&services.User{
		UserId: userId,
	})
	if !exist || Users[userIdx].Password != userPsw {
		return ErrResponse(fmt.Sprintf("can not find user"))
	}

	// DAO finding user
	return &services.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, app size :%v", Users[userIdx].UserId, len(Users[userIdx].Applications)),
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
	ErrResponse := func(errorMsg string) (*services.UserRegisterResponse, error) {
		return &services.UserRegisterResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
			Token:      "",
		}, nil
	}
	user := model.User{
		Password:     req.UserPassword,
		Name:         "empty", // delete
		Applications: make([]*model.Application, 0),
		Priority:     req.Priority,
		CreatedAt:    time.Now().UTC(),
	}

	if req.UserId == "" || req.UserPassword == "" {
		return ErrResponse("user id and user password can not be empty")
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return ErrResponse("error register user fmt")
	}

	_, exist := CheckUserExist(&services.User{
		UserId: userId,
	})
	if exist {
		return ErrResponse("already exist")
	}
	user.UserId = userId
	Users = append(Users, &user)
	return &services.UserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "register successfully",
		Token:      "",
	}, nil
}

/*
GET
req:
userid int64
userpassword string
*/
func (userService *UserService) GetInfo(ctx context.Context, req *services.UserGetInfoRequest) (*services.UserGetInfoResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserGetInfoResponse, error) {
		return &services.UserGetInfoResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
			User:       nil,
		}, nil
	}
	if req.UserId < 0 || req.UserPassword == "" {
		return ErrResponse("user id and user password can not be empty")
	}
	userIdx := 0
	if t, cor := CheckUserPswCorrect(req.UserId, req.UserPassword); !cor {
		userIdx = t
		return ErrResponse("password wrong")
	}

	user := ModelUserToServicesUser(*Users[userIdx])
	// user = *model.ModelUserToServicesUser(Users[userIdx])
	return &services.UserGetInfoResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, user psw :%v", req.UserId, req.UserPassword),
		User:       &user,
	}, nil
}

/*
req:
userid int64
app context string
*/
func (userService *UserService) SubmitApplication(ctx context.Context, req *services.UserSubmitApplicationRequest) (*services.UserSubmitApplicationResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserSubmitApplicationResponse, error) {
		return &services.UserSubmitApplicationResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
		}, nil
	}
	if req.UserId < 0 {
		return ErrResponse("user id and user password can not be empty")
	}
	userIdx, exist := CheckUserExist(&services.User{UserId: req.UserId})
	if !exist {
		return ErrResponse("can not find user")
	}
	app := &model.Application{
		ApplicationId: int64(len(AppList)),
		Context:       req.ApplicationContext,
		ReviewStatus:  false,
		UserId:        req.UserId,
	}
	Users[userIdx].Applications = append(Users[userIdx].Applications, app)
	AppList = append(AppList, app)
	return &services.UserSubmitApplicationResponse{
		StatusCode: 200,
		StatusMsg:  "submit application successfully",
	}, nil
}

/*
req:
userid int64
appid int64
*/
func (userService *UserService) RetrievalApplication(ctx context.Context, req *services.UserRetrievalApplicationRequest) (*services.UserRetrievalApplicationResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserRetrievalApplicationResponse, error) {
		return &services.UserRetrievalApplicationResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
		}, nil
	}
	if req.UserId < 0 || req.ApplicationId < 0 {
		return ErrResponse("user id or app id illegal")
	}
	if _, CheckUserExist := CheckUserExist(&services.User{UserId: req.UserId}); !CheckUserExist {
		return ErrResponse("user not find")
	}
	appIdx, exist := CheckAppExist(&services.Application{ApplicationId: req.ApplicationId})
	if !exist {
		return ErrResponse("application not find")
	}
	return &services.UserRetrievalApplicationResponse{
		StatusCode:   200,
		StatusMsg:    "ok",
		ReviewStatue: AppList[appIdx].ReviewStatus,
	}, nil
}

// userful app
func CheckUserExist(user *services.User) (int, bool) {
	userIdx := int(-1)
	for i, v := range Users {
		if user.UserId == v.UserId {
			userIdx = i
			break
		}
	}
	if userIdx == -1 {
		return -1, false
	} else {
		return userIdx, true
	}
}

func CheckUserPswCorrect(userId int64, userPsw string) (int, bool) {
	userIdx, exist := CheckUserExist(&services.User{
		UserId: userId,
	})
	if !exist {
		return -1, false
	}
	return userIdx, Users[userIdx].Password == userPsw
}

func ModelAppToServicesApp(apps []*model.Application) []*services.Application {
	res := make([]*services.Application, 0)
	for _, v := range apps {
		res = append(res, &services.Application{
			ApplicationId: v.ApplicationId,
			Context:       v.Context,
			ReviewStatus:  v.ReviewStatus,
		})
	}
	return res
}

func ModelUserToServicesUser(user model.User) services.User {
	return services.User{
		UserId:       user.UserId,
		Applications: ModelAppToServicesApp(user.Applications),
		Priority:     user.Priority,
	}
}

func CheckAppExist(app *services.Application) (int, bool) {
	appIdx := int(-1)
	for i, v := range AppList {
		if app.ApplicationId == v.ApplicationId {
			appIdx = i
			break
		}
	}
	if appIdx == -1 {
		return -1, false
	} else {
		return appIdx, true
	}
}

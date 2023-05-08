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
var Users map[int64]*model.User
var Reviewers map[int64]*model.User
var AppList map[int64]*model.Application

func init() {
	log.Println("init user server tmp cache")
	Users = make(map[int64]*model.User)
	Reviewers = make(map[int64]*model.User)
	AppList = make(map[int64]*model.Application)
}

/*
req:
UserId string
UserPassword string
rep

FINISH
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

	if _, exist := Users[userId]; !exist || Users[userId].Password != userPsw {
		return ErrResponse(fmt.Sprintf("can not find user"))
	}

	// DAO finding user
	return &services.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, app size :%v", Users[userId].UserId, len(Users[userId].Applications)),
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

	if req.UserId == "" || req.UserPassword == "" || req.UserName == "" {
		return ErrResponse("user id and user password can not be empty")
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return ErrResponse("illegal user id")
	}

	// 普通用户注册
	user := model.User{
		UserId:       userId,
		Password:     req.UserPassword,
		Name:         req.UserName,
		Applications: make([]int64, 0),
		Priority:     req.Priority,
		CreatedAt:    time.Now().UTC(),
	}

	// 审核人注册
	if user.Priority > 0 {
		Reviewers[user.UserId] = &user
	}

	if _, exist := Users[userId]; exist {
		return ErrResponse("already exist")
	}

	user.UserId = userId
	Users[userId] = &user
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
	if _, exist := Users[req.UserId]; !exist || Users[req.UserId].Password != req.UserPassword {
		return ErrResponse("password wrong")
	}

	user := ModelUserToServicesUser(*Users[req.UserId])
	// user = *model.ModelUserToServicesUser(Users[userIdx])
	return &services.UserGetInfoResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, welcome login", req.UserId),
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

	// 判断用户是否存在
	if _, exist := Users[req.UserId]; !exist {
		return ErrResponse("can not find user")
	}

	// 新建申请
	app := &model.Application{
		ApplicationId:    int64(len(AppList)),
		Context:          req.ApplicationContext,
		ReviewStatus:     false,
		UserId:           req.UserId,
		ApprovedReviewer: make(map[int64]bool), // 已通过的审核人
	}

	// 将 请求加入用户的请求列表
	userId := req.UserId
	Users[userId].Applications = append(Users[userId].Applications, app.ApplicationId)

	// 将请求加入 数据库
	AppList[app.ApplicationId] = app
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
	if _, ok := Users[req.UserId]; !ok {
		return ErrResponse("user not find")
	}

	res := AppIdListToApp(Users[req.UserId].Applications)

	return &services.UserRetrievalApplicationResponse{
		StatusCode:   200,
		StatusMsg:    "ok",
		Applications: res,
	}, nil
}

/*
提交申请
*/
func (userService *UserService) SubmitReview(ctx context.Context, req *services.UserSubmitReviewRequest) (*services.UserSubmitReviewResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserSubmitReviewResponse, error) {
		return &services.UserSubmitReviewResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
		}, nil
	}
	if req.UserId < 0 {
		return ErrResponse("user id illegal")
	}
	if _, userExist := Users[req.UserId]; !userExist {
		return ErrResponse("user not find")
	}
	if _, appExist := AppList[req.ApplicationId]; !appExist {
		return ErrResponse("app not find")
	}

	// updata sql data app
	if req.ReviewStatus {
		AppList[req.ApplicationId].ApprovedReviewer[req.UserId] = true
		AppList[req.ApplicationId].ReviewStatus = (len(AppList[req.ApplicationId].ApprovedReviewer) == len(Reviewers))
	}

	// update user data applist

	return &services.UserSubmitReviewResponse{
		StatusCode: 200,
		StatusMsg:  "ok",
	}, nil
}

// userful api

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

// 数据库中的 user 里只有 app ID List 因此要转换一下
func ModelUserToServicesUser(user model.User) services.User {
	res := services.User{
		UserId:       user.UserId,
		Applications: AppIdListToApp(user.Applications),
		Priority:     user.Priority,
	}
	return res
}

func AppIdListToApp(appIdList []int64) []*services.Application {
	res := make([]*services.Application, 0)
	for _, v := range appIdList {
		// TODO: DAO

		modelApp := AppList[int64(v)]
		res = append(res, &services.Application{
			ApplicationId: int64(modelApp.ApplicationId),
			Context:       string(modelApp.Context),
			ReviewStatus:  bool(modelApp.ReviewStatus),
		})
	}
	return res
}

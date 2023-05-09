package rpcserver

import (
	"context"
	"fmt"
	model "oa-review/user/model"
	services "oa-review/user/services"
	"strconv"
	"time"
)

type UserService struct {
	services.UnimplementedUserServiceServer
}

// TODO jwt

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

	// DAO finding user
	if _, exist := Users[userId]; !exist || Users[userId].Password != userPsw {
		return ErrResponse(fmt.Sprintf("can not find user"))
	}

	return &services.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, app size :%v", Users[userId].UserId, len(Users[userId].Applications)),
		Token:      "jwt token",
	}, nil
}

/*
用户注册
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

	// DAO find user by id
	if _, exist := Users[userId]; exist {
		return ErrResponse("already exist")
	}

	// 审核人注册 DAO create reviewer
	if user.Priority > 0 {
		Reviewers[user.UserId] = &model.Reviewer{
			UserId:       user.UserId,
			Name:         user.Name,
			Applications: user.Applications,
			Options:      make([]*model.ReviewOption, 0),
			Priority:     user.Priority,
			CreatedAt:    user.CreatedAt,
		}
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
	// DAO find user by user id
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
提交申请
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

	// DAO find user by user id 判断用户是否存在
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
	// DAO find user/ updata info
	Users[userId].Applications = append(Users[userId].Applications, app.ApplicationId)

	// 将请求加入 数据库
	// DAO create app
	AppList[app.ApplicationId] = app
	return &services.UserSubmitApplicationResponse{
		StatusCode: 200,
		StatusMsg:  "submit application successfully",
	}, nil
}

/*
检索请求列表
// TODO 无法处理 JSON 无法显示 0 值的问题
*/
func (userService *UserService) RetrievalApplication(ctx context.Context, req *services.UserRetrievalApplicationRequest) (*services.UserRetrievalApplicationResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserRetrievalApplicationResponse, error) {
		return &services.UserRetrievalApplicationResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
		}, nil
	}

	// DAO find usr by id
	// 查询实体之后 直接操作
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
		// find app by app id
		modelApp := AppList[int64(v)]
		res = append(res, &services.Application{
			ApplicationId: int64(modelApp.ApplicationId),
			Context:       string(modelApp.Context),
			ReviewStatus:  bool(modelApp.ReviewStatus),
		})
	}
	return res
}

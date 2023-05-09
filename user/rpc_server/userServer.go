package rpcserver

import (
	"context"
	"fmt"
	services "oa-review/proto/services"
	model "oa-review/user/model"
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

	user, err := model.NewUserDaoInstance().FindUserByUserId(userId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if user.Password != userPsw {
		return ErrResponse("password incorect")
	}

	return &services.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, app size :%v", user.UserId, len(user.Applications)),
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
	exist, err := model.NewUserDaoInstance().CheckUserExist(userId)
	if exist {
		return ErrResponse("user already exist")
	}
	if err != nil && err.Error() != "record not found" {
		return ErrResponse(err.Error())
	}

	// 审核人注册 DAO create reviewer
	if user.Priority > 0 {
		reviewer := model.Reviewer{
			ReviewerId:   user.UserId,
			Name:         user.Name,
			Applications: user.Applications,
			Options:      make([]*model.ReviewOption, 0),
			Priority:     user.Priority,
			CreatedAt:    user.CreatedAt,
		}
		model.NewReviewerDaoInstance().CreateReviewer(&reviewer)
	}

	// // DAO create user
	// user.UserId = userId
	// Users[userId] = &user
	user.UserId = userId
	_, err = model.NewUserDaoInstance().CreateUser(&user)
	if err != nil {
		return ErrResponse(err.Error())
	}
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
	user, err := model.NewUserDaoInstance().FindUserByUserId(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if user.Password != req.UserPassword {
		return ErrResponse("password incorect")
	}

	servicesUser, err := ModelUserToServicesUser(user)
	if err != nil {
		return ErrResponse(err.Error())
	}
	// user = *model.ModelUserToServicesUser(Users[userIdx])
	return &services.UserGetInfoResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, welcome login", req.UserId),
		User:       servicesUser,
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

	exist, err := model.NewUserDaoInstance().CheckUserExist(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if !exist {
		return ErrResponse("user not found")
	}

	// 新建申请
	appTableSize, err := model.NewApplicationDaoInstance().TableSize()
	if err != nil {
		return ErrResponse(err.Error())
	}
	app := &model.Application{
		ApplicationId:    appTableSize + 1,
		Context:          req.ApplicationContext,
		ReviewStatus:     false,
		UserId:           req.UserId,
		ApprovedReviewer: make(map[int64]bool), // 已通过的审核人
	}

	// 将 请求加入用户的请求列表
	userId := req.UserId
	// // DAO find user/ updata info
	// Users[userId].Applications = append(Users[userId].Applications, app.ApplicationId)
	if err := model.NewUserDaoInstance().AddApplicationForUser(userId, app.ApplicationId); err != nil {
		return ErrResponse(err.Error())
	}

	// // 将请求加入 数据库
	// // DAO create app
	// AppList[app.ApplicationId] = app
	if _, err := model.NewApplicationDaoInstance().CreateApplication(app); err != nil {
		return ErrResponse(err.Error())
	}
	return &services.UserSubmitApplicationResponse{
		StatusCode: 200,
		StatusMsg:  "submit application successfully",
	}, nil
}

/*
检索请求列表
*/
func (userService *UserService) RetrievalApplication(ctx context.Context, req *services.UserRetrievalApplicationRequest) (*services.UserRetrievalApplicationResponse, error) {
	ErrResponse := func(errorMsg string) (*services.UserRetrievalApplicationResponse, error) {
		return &services.UserRetrievalApplicationResponse{
			StatusCode: 400,
			StatusMsg:  errorMsg,
		}, nil
	}

	// // DAO find usr by id
	// // 查询实体之后 直接操作
	// if _, ok := Users[req.UserId]; !ok {
	// 	return ErrResponse("user not find")
	// }
	user, err := model.NewUserDaoInstance().FindUserByUserId(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	res, err := AppIdListToApp(user.Applications)
	if err != nil {
		return ErrResponse(err.Error())
	}
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
func ModelUserToServicesUser(user *model.User) (*services.User, error) {
	applications, err := AppIdListToApp(user.Applications)
	if err != nil {
		return nil, err
	}
	res := services.User{
		UserId:       user.UserId,
		Applications: applications,
		Priority:     user.Priority,
	}
	return &res, nil
}

func AppIdListToApp(appIdList []int64) ([]*services.Application, error) {
	res := make([]*services.Application, 0)
	for _, v := range appIdList {
		// TODO: DAO
		// find app by app id
		// modelApp := AppList[int64(v)]
		modelApp, err := model.NewApplicationDaoInstance().FindApplicationById(int64(v))
		if err != nil {
			return nil, err
		}
		res = append(res, &services.Application{
			ApplicationId: int64(modelApp.ApplicationId),
			Context:       string(modelApp.Context),
			ReviewStatus:  bool(modelApp.ReviewStatus),
		})
	}
	return res, nil
}

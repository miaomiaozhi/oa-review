package rpcserver

import (
	"context"
	"fmt"
	dao "oa-review/dao"
	middleware "oa-review/middleware"
	services "oa-review/proto/services"
	"strconv"
	"time"
)

type UserService struct {
	services.UnimplementedUserServiceServer
}

/*
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
		return ErrResponse("error login user format")
	}

	user, err := dao.NewUserDaoInstance().FindUserByUserId(userId)
	if err != nil {
		return ErrResponse(fmt.Sprintf("error on find user by user id: %v", err.Error()))
	}
	if user.Password != userPsw {
		return ErrResponse("password incorect")
	}

	jwtToken, err := middleware.CreateUserJwtToken(userId, user.Priority)
	if err != nil {
		return ErrResponse(fmt.Sprintf("error on create user Jwt token: %v", err.Error()))
	}

	return &services.UserLoginResponse{
		StatusCode: 200,
		StatusMsg:  fmt.Sprintf("user id :%v, app size :%v", user.UserId, len(user.Applications)),
		Token:      jwtToken,
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
		}, nil
	}

	if req.UserId == "" || req.UserPassword == "" || req.UserName == "" {
		return ErrResponse("user id and user password can not be empty")
	}
	if req.Priority < 0 {
		return ErrResponse("priority illegal")
	}
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return ErrResponse("illegal user id")
	}

	// 普通用户注册
	user := dao.User{
		UserId:       userId,
		Password:     req.UserPassword,
		Name:         req.UserName,
		Applications: make([]int64, 0),
		Priority:     req.Priority,
		CreatedAt:    time.Now().UTC(),
	}

	// DAO find user by id
	exist, err := dao.NewUserDaoInstance().CheckUserExist(userId)
	if exist {
		return ErrResponse("user already exist")
	}
	if err != nil && err.Error() != "record not found" {
		return ErrResponse(err.Error())
	}

	// 审核人注册 DAO create reviewer
	if user.Priority > 0 {
		reviewer := dao.Reviewer{
			ReviewerId:   user.UserId,
			Name:         user.Name,
			Applications: user.Applications,
			Options:      make([]*dao.ReviewOption, 0),
			Priority:     user.Priority,
			CreatedAt:    user.CreatedAt,
		}
		if _, err := dao.NewReviewerDaoInstance().CreateReviewer(&reviewer); err != nil {
			return ErrResponse("Error on create reviewer" + err.Error())
		}
	}

	// // DAO create user
	// user.UserId = userId
	// Users[userId] = &user
	user.UserId = userId
	_, err = dao.NewUserDaoInstance().CreateUser(&user)
	if err != nil {
		return ErrResponse("Error on create user" + err.Error())
	}
	return &services.UserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "register successfully",
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
	// 用户鉴权
	tokenId, _, err := middleware.UserAuthorize(req.Token)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if tokenId != req.UserId {
		return ErrResponse("no authorization to get user info")
	}

	user, err := dao.NewUserDaoInstance().FindUserByUserId(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}

	servicesUser, err := daoUserToServicesUser(user)
	if err != nil {
		return ErrResponse(err.Error())
	}
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

	// 用户鉴权
	tokenId, _, err := middleware.UserAuthorize(req.Token)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if tokenId != req.UserId {
		return ErrResponse("no authorization to submit application")
	}

	// 新建申请
	appTableSize, err := dao.NewApplicationDaoInstance().TableSize()
	if err != nil {
		return ErrResponse(err.Error())
	}
	app := &dao.Application{
		ApplicationId:    appTableSize + 1,
		Context:          req.ApplicationContext,
		ReviewStatus:     false,
		UserId:           req.UserId,
		ApprovedReviewer: make(map[int64]bool), // 已通过的审核人
	}

	// 将 请求加入用户的请求列表
	userId := req.UserId
	if err := dao.NewUserDaoInstance().AddApplicationForUser(userId, app.ApplicationId); err != nil {
		return ErrResponse(err.Error())
	}

	if _, err := dao.NewApplicationDaoInstance().CreateApplication(app); err != nil {
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

	// 用户鉴权
	tokenId, _, err := middleware.UserAuthorize(req.Token)
	if err != nil {
		return ErrResponse(err.Error())
	}
	if tokenId != req.UserId {
		return ErrResponse("no authorization to retrival application")
	}

	user, err := dao.NewUserDaoInstance().FindUserByUserId(req.UserId)
	if err != nil {
		return ErrResponse(err.Error())
	}
	res, err := appIdListToApp(user.Applications)
	if err != nil {
		return ErrResponse(err.Error())
	}
	return &services.UserRetrievalApplicationResponse{
		StatusCode:   200,
		StatusMsg:    "ok",
		Applications: res,
	}, nil
}

// internal api
// 数据库中的 user 里只有 app ID List 因此要转换一下
func daoUserToServicesUser(user *dao.User) (*services.User, error) {
	applications, err := appIdListToApp(user.Applications)
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

func appIdListToApp(appIdList []int64) ([]*services.Application, error) {
	res := make([]*services.Application, 0)
	for _, v := range appIdList {
		daoApp, err := dao.NewApplicationDaoInstance().FindApplicationById(int64(v))
		if err != nil {
			return nil, err
		}
		res = append(res, &services.Application{
			ApplicationId: int64(daoApp.ApplicationId),
			Context:       string(daoApp.Context),
			ReviewStatus:  bool(daoApp.ReviewStatus),
		})
	}
	return res, nil
}

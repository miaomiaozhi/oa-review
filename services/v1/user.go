package v1

import (
	"oa-review/bean"
	"oa-review/dao"
	"oa-review/internal/wrapper"
	"oa-review/logger"
	v1_req "oa-review/models/protoreq/v1"
	v1_resp "oa-review/models/protoresp/v1"
	"strconv"
)

// TODO: validator

func UserLogin(ctx *wrapper.Context, reqBody interface{}) error {
	req := reqBody.(*v1_req.UserLoginRequest)
	logger.Debug("user info", req.UserId, req.UserPassword)
	userIdStr, userPsw := req.UserId, req.UserPassword
	// logger.Info("handle user login, user info :", userIdStr, userPsw)
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if userIdStr == "" || userPsw == "" || err != nil || userId < 0 {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户信息错误")
		return nil
	}
	logger.Debug("user info", userId)
	// 判断是否存在
	if exist, _ := dao.NewUserDaoInstance().CheckUserExist(userId); !exist {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户不存在")
		return nil
	}

	user, err := dao.NewUserDaoInstance().FindUserByUserId(userId)
	if err != nil {
		return err
	}
	if user.Password != userPsw {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户密码错误")
		return nil
	}

	// TODO: jwt
	wrapper.SendApiOKResponse(ctx, nil, "登录成功")
	return nil
}

func UserRegister(ctx *wrapper.Context, reqBody interface{}) error {
	req := reqBody.(*v1_req.UserRegisterRequest)
	userIdStr, userPsw := req.UserId, req.UserPassword
	// logger.Info("handle user login, user info :", userIdStr, userPsw)
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if userIdStr == "" || userPsw == "" || err != nil || req.UserName == "" || req.Priority < 0 || userId < 0 {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户信息错误")
		return nil
	}
	// 判断是否存在
	if exist, err := dao.NewUserDaoInstance().CheckUserExist(userId); err != nil || exist {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户 id 错误或用户已存在")
		return nil
	}

	usr := bean.User{
		Id:           userId,
		Password:     req.UserPassword,
		Name:         req.UserName,
		Applications: make(bean.Apps, 0),
		Priority:     req.Priority,
	}

	if _, err := dao.NewUserDaoInstance().CreateUser(&usr); err != nil {
		return err
	}

	wrapper.SendApiOKResponse(ctx, nil, "register success")
	logger.Info("Apiwrapper user register ok")
	return nil
}

func UserGetInfo(ctx *wrapper.Context, reqBody interface{}) error {
	// TODO validator
	req := reqBody.(*v1_req.UserGetInfoRequest)
	user, err := dao.NewUserDaoInstance().FindUserByUserId(req.UserId)
	if err != nil {
		logger.Error("数据库查询错误", err.Error())
		return err
	}

	apps := make([]*v1_resp.Application, 0, len(user.Applications))
	for _, id := range user.Applications {
		app, err := dao.NewApplicationDaoInstance().FindApplicationById(id)
		if err != nil {
			logger.Error("数据库查询错误", err.Error())
			return err
		}
		apps = append(apps, &v1_resp.Application{
			Context:      app.Context,
			ReviewStatus: app.ReviewStatus,
		})
	}
	resp := v1_resp.UserGetInfoResponse{
		Id:           user.Id,
		Name:         user.Name,
		Priority:     user.Priority,
		Applications: apps,
	}

	wrapper.SendApiOKResponse(ctx, resp, "查询成功")
	return nil
}

func UserSubmitApplication(ctx *wrapper.Context, reqBody interface{}) error {
	// TODO validator
	// 进入流程
	req := reqBody.(*v1_req.UserSubmitApplicationRequest)
	appTableSize, err := dao.NewApplicationDaoInstance().TableSize()
	if err != nil {
		return err
	}
	app := &bean.Application{
		Id:               appTableSize + 1,
		Context:          req.ApplicationContext,
		ReviewStatus:     false,
		UserId:           req.UserId,
		ApprovedReviewer: make(bean.ApproverMap),
	}
	if err := dao.NewUserDaoInstance().AddApplicationForUser(req.UserId, appTableSize+1); err != nil {
		return err
	}
	if _, err := dao.NewApplicationDaoInstance().CreateApplication(app); err != nil {
		return err
	}

	wrapper.SendApiOKResponse(ctx, nil, "提交成功")
	return nil
}

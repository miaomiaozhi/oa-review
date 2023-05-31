package v1

import (
	"encoding/json"
	"oa-review/bean"
	"oa-review/dao"
	"oa-review/internal/wrapper"
	"oa-review/logger"
	"oa-review/middleware"
	v1_req "oa-review/models/protoreq/v1"
	v1_resp "oa-review/models/protoresp/v1"
	"strconv"
)

// TODO: validator

func UserLogin(ctx *wrapper.Context, reqBody interface{}) error {
	logger.Info("handle user Login now")
	req := reqBody.(*v1_req.UserLoginRequest)

	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		wrapper.SendApiBadRequestResponse(ctx, nil, "user Id 错误")
		return nil
	}

	logger.Debug("user info", userId)
	if exist, _ := dao.NewUserDaoInstance().CheckUserExist(userId); !exist {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户不存在")
		return nil
	}

	user, err := dao.NewUserDaoInstance().FindUserByUserId(userId)
	if err != nil {
		return err
	}
	if user.Password != req.UserPassword {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户密码错误")
		return nil
	}

	jwtToken, err := middleware.GenJwtToken(user.Id, user.Name, int64(user.Priority))
	if err != nil {
		return err
	}
	auth := AuthResult{
		Token:    jwtToken,
		UserID:   user.Id,
		UserName: user.Name,
		Priority: int64(user.Priority),
	}

	jsonData, err := json.Marshal(auth)
	if err != nil {
		logger.Errorf("marshal auth to json data failed: %v", err.Error())
		return err
	}

	resp := v1_resp.UserLoginResponse{
		Token: string(jsonData),
	}
	logger.Info("json token is", resp.Token)
	wrapper.SendApiOKResponse(ctx, resp, "登录成功")
	return nil
}

func UserRegister(ctx *wrapper.Context, reqBody interface{}) error {
	logger.Info("handle user Register now")
	req := reqBody.(*v1_req.UserRegisterRequest)

	logger.Info("handle user login, user info :", req.UserId, req.UserName)
	userId, _ := strconv.ParseInt(req.UserId, 10, 64)

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
	// 审核人注册 DAO create reviewer
	if usr.Priority > 0 {
		reviewer := bean.Reviewer{
			Id:           usr.Id,
			Name:         usr.Name,
			Applications: usr.Applications,
			Options:      make([]*bean.ReviewOption, 0),
			Priority:     usr.Priority,
			CreatedAt:    usr.CreatedAt,
		}
		if _, err := dao.NewReviewerDaoInstance().CreateReviewer(&reviewer); err != nil {
			return err
		}
	}

	wrapper.SendApiOKResponse(ctx, nil, "注册成功")
	logger.Info("Apiwrapper user register ok")
	return nil
}

func UserGetInfo(ctx *wrapper.Context, reqBody interface{}) error {
	logger.Info("handle user GetInfo now")

	req := reqBody.(*v1_req.UserGetInfoRequest)
	if req.UserId != ctx.UserToken.UserID {
		wrapper.SendApiUnAuthResponse(ctx, nil, "token user Id 与请求 user id 不一致")
		return nil
	}
	// 判断是否存在
	if exist, err := dao.NewUserDaoInstance().CheckUserExist(req.UserId); err != nil || !exist {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户不存在")
		return nil
	}

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
			Context:       app.Context,
			ReviewStatus:  app.ReviewStatus,
			ApplicationId: app.Id,
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
	logger.Info("handle user SubmitApplication now")

	req := reqBody.(*v1_req.UserSubmitApplicationRequest)
	if req.UserId != ctx.UserToken.UserID {
		wrapper.SendApiUnAuthResponse(ctx, nil, "token user Id 与请求 user id 不一致")
		return nil
	}
	appTableSize, err := dao.NewApplicationDaoInstance().TableSize()
	if err != nil {
		return err
	}
	workflow, finish := GetWorkFlow()
	if workflow == nil {
		wrapper.SendApiBadRequestResponse(ctx, nil, "工作流不存在")
		return nil
	}

	if !workflow.CheckWorkFlowExist(req.UserId, req.ApplicationContext) {
		wrapper.SendApiBadRequestResponse(ctx, nil, "流程不存在")
		return nil
	}
	if finish {
		wrapper.SendApiBadRequestResponse(ctx, nil, "流程已完成")
		return nil
	}
	if workflow.IsStarted() {
		wrapper.SendApiBadRequestResponse(ctx, nil, "流程已开始")
		return nil
	}

	// 启动流程
	workflow.Start()
	app := &bean.Application{
		Id:               appTableSize + 1,
		Context:          req.ApplicationContext,
		ReviewStatus:     false,
		UserId:           req.UserId,
		ApprovedReviewer: make(bean.ApproverMap),
	}
	workflow.SetApplicationId(app.Id)
	if err := dao.NewUserDaoInstance().AddApplicationForUser(req.UserId, appTableSize+1); err != nil {
		return err
	}
	if _, err := dao.NewApplicationDaoInstance().CreateApplication(app); err != nil {
		return err
	}

	wrapper.SendApiOKResponse(ctx, nil, "提交成功")
	return nil
}

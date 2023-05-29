package v1

import (
	"oa-review/internal/wrapper"
	logger "oa-review/logger"
	v1 "oa-review/models/protoreq/v1"
	services_v1 "oa-review/services/v1"
)

type UserController struct {
}

func (UserController) UserLogin(ctx *wrapper.Context) {
	logger.Info("user controller wrapper login")
	wrapper.ApiWrapper(ctx, services_v1.UserLogin, true, &v1.UserLoginRequest{}, nil)
}

func (UserController) UserRegister(ctx *wrapper.Context) {
	logger.Info("user controller wrapper register")
	wrapper.ApiWrapper(ctx, services_v1.UserRegister, true, &v1.UserRegisterRequest{}, nil)
}

func (UserController) UserGetInfo(ctx *wrapper.Context) {
	logger.Info("user controller wrapper GetInfo")
	wrapper.ApiWrapper(ctx, services_v1.UserGetInfo, true, &v1.UserGetInfoRequest{}, nil)
}

func (UserController) UserSubmitApplication(ctx *wrapper.Context) {
	logger.Info("user controller wrapper SubmitApplication")
	wrapper.ApiWrapper(ctx, services_v1.UserSubmitApplication, true, &v1.UserSubmitApplicationRequest{}, nil)
}

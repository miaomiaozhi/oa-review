package v1

import (
	"oa-review/internal/wrapper"
	logger "oa-review/logger"
	v1 "oa-review/services/v1"
)

type UserController struct {
}

func (UserController) UserLogin(ctx *wrapper.Context) {
	logger.Info("user controller wrapper login")
	wrapper.ApiWrapper(ctx, v1.UserLogin, true, nil, nil)
}

func (UserController) UserRegister(ctx *wrapper.Context) {
	logger.Info("user controller wrapper register")
	wrapper.ApiWrapper(ctx, v1.UserRegister, true, nil, nil)
}

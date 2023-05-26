package v1

import (
	"oa-review/internal/wrapper"
	"oa-review/logger"
	v1_req "oa-review/models/protoreq/v1"
	"strconv"
)

func UserLogin(ctx *wrapper.Context, reqBody interface{}) error {

	req := v1_req.UserLoginRequest{}
	if err := ctx.ReadJSON(&req); err != nil {
		wrapper.SendApiBadRequestResponse(ctx, nil, "参数错误")
		return nil
	}
	userIdStr, userPsw := req.UserId, req.UserPassword
	logger.Info("handle user login, user info :", userIdStr, userPsw)
	_, err := strconv.ParseInt(userIdStr, 10, 64)
	if userIdStr == "" || userPsw == "" || err != nil {
		wrapper.SendApiBadRequestResponse(ctx, nil, "用户信息错误")
		return nil
	}

	//user, err := dao.NewUserDaoInstance().FindUserByUserId(userId)
	//if err != nil {
	//	return ErrResponse(fmt.Sprintf("error on find user by user id: %v", err.Error()))
	//}
	//if user.Password != userPsw {
	//	return ErrResponse("password incorect")
	//}
	//
	//jwtToken, err := middleware.CreateUserJwtToken(userId, user.Priority)
	//if err != nil {
	//	return ErrResponse(fmt.Sprintf("error on create user Jwt token: %v", err.Error()))
	//}

	wrapper.SendApiOKResponse(ctx, nil, "login success")
	return nil
}

func UserRegister(ctx *wrapper.Context, reqBody interface{}) error {
	wrapper.SendApiOKResponse(ctx, nil, "register success")
	logger.Info("Apiwrapper user register ok")
	return nil
}

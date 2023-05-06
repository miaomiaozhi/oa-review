package handler

import (
	"context"
	pb "oa-review/gateway/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	req := &pb.UserLoginRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	resp, err := C.userClient.Login(context.Background(), req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(resp)
}

func Register(ctx iris.Context) {
	req := &pb.UserRegisterRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	resp, err := C.userClient.Register(context.Background(), req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(resp)
}

func GetInfo(ctx iris.Context) {
	req := &pb.UserGetInfoRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	resp, err := C.userClient.GetInfo(context.Background(), req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(resp)
}

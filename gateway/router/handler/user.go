package handler

import (
	"fmt"
	"log"
	pb "oa-review/proto/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	req := &pb.UserLoginRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.Login(ctx, req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func Register(ctx iris.Context) {
	req := &pb.UserRegisterRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.Register(ctx, req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func GetInfo(ctx iris.Context) {
	req := &pb.UserGetInfoRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		log.Printf("Error on user get info: %v\n", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.GetInfo(ctx, req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func Submit(ctx iris.Context) {
	req := &pb.UserSubmitApplicationRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		log.Printf("Error on user submit: %v\n", err)

		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.SubmitApplication(ctx, req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func Retrieval(ctx iris.Context) {
	req := &pb.UserRetrievalApplicationRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		log.Printf("Error on user retrieval: %v\n", err)

		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.RetrievalApplication(ctx, req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func SubmitReview(ctx iris.Context) {
	req := &pb.UserSubmitReviewRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		errResponse(ctx, fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.SubmitReview(ctx, req)
	if err != nil {
		errResponse(ctx, fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func WithdrawReview(ctx iris.Context) {
	req := &pb.UserWithdrawReviewRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		errResponse(ctx, fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.WithdrawReview(ctx, req)
	if err != nil {
		errResponse(ctx, fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func errResponse(ctx iris.Context, errMsg string) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.WriteString(errMsg)
}

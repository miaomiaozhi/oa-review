package handler

import (
	"context"
	"fmt"
	"log"
	pb "oa-review/gateway/services"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	req := &pb.UserLoginRequest{}
	if err := ctx.ReadJSON(req); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	resp, err := C.userClient.Login(context.Background(), req)
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
	resp, err := C.userClient.Register(context.Background(), req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

func GetInfo(ctx iris.Context) {
	req := &pb.UserGetInfoRequest{}
	// log.Println("to here 0")

	if err := ctx.ReadJSON(req); err != nil {
		log.Printf("Error on user get info: %v\n", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(fmt.Sprintf("Error on read:%v", err.Error()))
		return
	}
	// log.Println("to here 1")
	resp, err := C.userClient.GetInfo(context.Background(), req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	// log.Println("to here 2")
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
	resp, err := C.userClient.SubmitApplication(context.Background(), req)
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
	resp, err := C.userClient.RetrievalApplication(context.Background(), req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(fmt.Sprintf("Error on connect client:%v", err.Error()))
		return
	}
	ctx.JSON(resp)
}

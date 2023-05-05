package handler

import "github.com/kataras/iris/v12"

func Login(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "login",
	})
}

func Register(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "register",
	})
}

func GetInfo(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "get info",
	})
}

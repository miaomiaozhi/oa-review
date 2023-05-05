package handler

import "github.com/kataras/iris/v12"

func Login(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "login",
	})
}

func GetInfo(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "get info",
	})
}
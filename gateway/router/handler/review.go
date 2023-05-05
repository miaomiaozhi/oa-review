package handler

import "github.com/kataras/iris/v12"

func Review(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "review",
	})
}

package handler

import "github.com/kataras/iris/v12"

func FirstReview(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "first",
	})
}

func FinalReview(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"message": "final",
	})
}

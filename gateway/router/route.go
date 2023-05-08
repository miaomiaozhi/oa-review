package router

import (
	"oa-review/gateway/router/handler"

	"github.com/kataras/iris/v12"
)

func NewRouter() *iris.Application {
	r := iris.New()
	userApi := r.Party("/user")
	{
		// midware
		userApi.Use(iris.Compression)
		// get info
		userApi.Get("/info", handler.GetInfo)
		// login
		userApi.Post("/login", handler.Login)
		// register
		userApi.Post("/register", handler.Register)
		// submit
		userApi.Post("/submit", handler.Submit)
		// retrieval application
		userApi.Get("/retrieval", handler.Retrieval)

		// User review need auth
		// submit review
		userApi.Post("/review/submit", handler.SubmitReview)
		// withdraw review
		userApi.Post("/review/withdraw", handler.WithdrawReview)
	}

	// test connect
	// TODO: delete
	r.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "ok",
		})
	})
	return r
}

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
		userApi.Get("/", handler.GetInfo)
		// login
		userApi.Post("/", handler.Login)
	}
	reviewApi := r.Party("review")
	{
		// midware
		reviewApi.Use(iris.Compression)
		// first review
		reviewApi.Post("/", handler.FirstReview)
		// final review
		reviewApi.Post("/", handler.FinalReview)
	}
	return r
}

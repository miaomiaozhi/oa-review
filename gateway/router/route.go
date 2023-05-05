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
	}
	reviewApi := r.Party("review")
	{
		// midware
		reviewApi.Use(iris.Compression)
		// review
		reviewApi.Post("/", handler.Review)
	}
	return r
}

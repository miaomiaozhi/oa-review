package router

import (
	handler "oa-review/router/handler"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

// TODO add jwt midware

func NewRouter() *iris.Application {
	r := iris.New()

	// midware
	loadMidware(r)

	userApi := r.Party("/user")
	{
		// login
		userApi.Post("/login", handler.Login)
		// register
		userApi.Post("/register", handler.Register)
	}

	// User api need authorize
	userAuthorizeApi := r.Party("/user")
	{
		// get info
		userAuthorizeApi.Get("/info", handler.GetInfo)
		// submit application
		userAuthorizeApi.Post("/submit", handler.Submit)
		// retrieval application
		userAuthorizeApi.Get("/retrieval", handler.Retrieval)

	}

	// Reviewer api need authorize
	reviewerAuthorizeApi := r.Party("/user/review")
	{
		// submit review
		reviewerAuthorizeApi.Post("/submit", handler.SubmitReview)
		// withdraw review
		reviewerAuthorizeApi.Post("/withdraw", handler.WithdrawReview)
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

func loadMidware(r *iris.Application) {
	r.Use(iris.Compression)
	r.Use(recover.New())
	r.Use(logger.New())
}

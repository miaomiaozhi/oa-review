package router

import (
	pkg "oa-review/gateway/pkg"
	"oa-review/gateway/router/handler"

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

	// User api need auth
	userAuthorizeApi := r.Party("/user", pkg.Authorize)
	{
		// get info
		userAuthorizeApi.Get("/info", handler.GetInfo)
		// submit application
		userAuthorizeApi.Post("/submit", handler.Submit)
		// retrieval application
		userAuthorizeApi.Get("/retrieval", handler.Retrieval)

		// submit review
		userAuthorizeApi.Post("/review/submit", handler.SubmitReview)
		// withdraw review
		userAuthorizeApi.Post("/review/withdraw", handler.WithdrawReview)
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

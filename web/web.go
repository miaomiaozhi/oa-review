package web

import (
	"github.com/kataras/iris/v12"
	"oa-review/logger"
	"oa-review/routers"
)

type OaReviewWeb struct{}

func newApp() *iris.Application {
	return iris.New()
}

func (OaReviewWeb) Run() {
	app := newApp()
	routers.IrisRouter{}.InitApp(app)
	app.Listen(":8080")
	logger.Info("app listening port: 8080")
}

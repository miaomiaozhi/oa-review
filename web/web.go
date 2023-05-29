package web

import (
	"oa-review/conf"
	"oa-review/logger"
	"oa-review/routers"

	"github.com/kataras/iris/v12"
)

type OaReviewWeb struct {
}

func newApp() *iris.Application {
	return iris.New()
}

func New() *OaReviewWeb {
	return &OaReviewWeb{}
}

func (*OaReviewWeb) Run(config *conf.OaReviewConf) {
	InitWorkFlow(config)
	app := newApp()
	routers.IrisRouter{}.InitApp(app)
	logger.Info("app listening port: 8080")
	app.Listen(":8080")
}

package web

import (
	"oa-review/conf"
	"oa-review/logger"
	"oa-review/routers"
	workflow "oa-review/services/v1"
	"strconv"

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
	workflow.InitWorkFlow(config)
	app := newApp()
	routers.IrisRouter{}.InitApp(app)
	port := config.GetInt("web.port", 8080)
	logger.Info("app listening port:", port)
	app.Listen(":" + strconv.Itoa(int(port)))
}

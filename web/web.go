package web

import (
	"oa-review/conf"
	"oa-review/logger"
	"oa-review/routers"
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
	InitWorkFlow(config)
	app := newApp()
	routers.IrisRouter{}.InitApp(app)
	port := config.GetInt("web.port", 8080)
	logger.Info("app listening port:", port)
	app.Listen(":" + strconv.Itoa(int(port)))
}

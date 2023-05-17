package routers

import (
	"oa-review/logger"
	v1 "oa-review/routers/api/v1"

	"github.com/kataras/iris/v12"
)

type IrisRouter struct{}

func (IrisRouter) InitApp(app *iris.Application) {
	logger.Info("init app")
	loadMiddlerware(app)
	appRouter := app.Party("/")
	{
		//// 注册鉴权路由
		//v1.RegisterAuthRouter(appRouter)
		//{
		//}
		v1.RegisterHealthRouter(appRouter)
	}
}

func loadMiddlerware(app *iris.Application) {
	logger.Info("load middle ware for app")
	//app.Use()
}

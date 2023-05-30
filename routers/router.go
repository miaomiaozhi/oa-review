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
		// 注册健康检查路由
		v1.RegisterHealthRouter(appRouter)

		// 注册User路由
		// 包括普通用户以及审核用户
		v1.RegisterUserRouter(appRouter)
	}
	logger.Info("init app success")
}

func loadMiddlerware(app *iris.Application) {
	logger.Info("load middle ware for app")
	//app.Use()
}

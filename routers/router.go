package routers

import (
	"github.com/kataras/iris/v12"
)

type IrisRouter struct{}

func (IrisRouter) InitApp(app *iris.Application) {
	loadMiddlerware(app)
	appRouter := app.Party("/oa-review")
	{
		
	}
}

func loadMiddlerware(app *iris.Application) {
	app.Use()
}
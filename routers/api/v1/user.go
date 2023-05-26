package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"net/http"
	v1_controller "oa-review/controllers/v1"
	"oa-review/internal/wrapper"
	"oa-review/logger"
)

func RegisterUserRouter(party router.Party) {
	partyNotLogin := party.Party("/user")
	{
		// without auth
		partyNotLogin.Handle(http.MethodPost, "/login", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserLogin))
		partyNotLogin.Handle(http.MethodPost, "/register", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserRegister))
	}
	//v1 := party.Party("/user")
	//{
	//	v1.Get("info", wrapper.Handler(v1_controller.HealthController{}.Test))
	//}

	logger.Info("register user router success")
}

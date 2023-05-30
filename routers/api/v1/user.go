package v1

import (
	"net/http"
	v1_controller "oa-review/controllers/v1"
	"oa-review/internal/wrapper"
	"oa-review/logger"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterUserRouter(party router.Party) {
	v1_notLogin := party.Party("/user")
	{
		// without auth
		v1_notLogin.Handle(http.MethodPost, "/login", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserLogin))
		v1_notLogin.Handle(http.MethodPost, "/register", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserRegister))

		v1_notLogin.Handle(http.MethodGet, "/info", wrapper.Handler(v1_controller.UserController{}.UserGetInfo))
		v1_notLogin.Handle(http.MethodPost, "/submit", wrapper.Handler(v1_controller.UserController{}.UserSubmitApplication))
	}
	v1_login := party.Party("/user/review")
	{
		v1_login.Handle(http.MethodPost, "/submit", wrapper.HandlerNotLogin(v1_controller.ReviewerController{}.ReviewerSubmit))
		v1_login.Handle(http.MethodPost, "/withdraw", wrapper.HandlerNotLogin(v1_controller.ReviewerController{}.ReviewerWithDraw))
	}

	logger.Info("register user router success")
}

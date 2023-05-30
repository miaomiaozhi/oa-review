package v1

import (
	"net/http"
	v1_controller "oa-review/controllers/v1"
	"oa-review/internal/wrapper"
	"oa-review/logger"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterUserRouter(party router.Party) {
	// TODO: need auth
	partyNotLogin := party.Party("/user")
	{
		// without auth
		partyNotLogin.Handle(http.MethodPost, "/login", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserLogin))
		partyNotLogin.Handle(http.MethodPost, "/register", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserRegister))

		// TODO
		partyNotLogin.Handle(http.MethodGet, "/info", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserGetInfo))
		partyNotLogin.Handle(http.MethodPost, "/submit", wrapper.HandlerNotLogin(v1_controller.UserController{}.UserSubmitApplication))
	}
	v1_reviewer := party.Party("/user/review")
	{
		// TODO
		v1_reviewer.Handle(http.MethodPost, "/submit", wrapper.HandlerNotLogin(v1_controller.ReviewerController{}.ReviewerSubmit))
		v1_reviewer.Handle(http.MethodPost, "/withdraw", wrapper.HandlerNotLogin(v1_controller.ReviewerController{}.ReviewerWithDraw))

	}
	// v1 := party.Party("/user")
	// {
	// 	v1.Handle(http.MethodGet, "/info", wrapper.Handler(v1_controller.UserController{}.UserGetInfo))
	// }

	logger.Info("register user router success")
}

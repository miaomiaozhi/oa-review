package v1

import (
	"net/http"
	"oa-review/internal/wrapper"
	v1_controller "oa-review/controllers/v1"
	"github.com/kataras/iris/v12/core/router"
)

func RegisterAuthRouter(party router.Party) {
	v1 := party.Party("/auth/v1")
	{
		v1.Handle(http.MethodGet, "/verify/token", wrapper.Handler(v1_controller.AuthController{}.VerifyToken))
		v1.Handle(http.MethodGet, "/test/header", wrapper.Handler(v1_controller.AuthController{}.TestHeader))
	}
}

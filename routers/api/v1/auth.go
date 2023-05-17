package v1

//import (
//	"net/http"
//	v1_controller "oa-review/controllers/v1"
//	"oa-review/internal/wrapper"
//
//	"github.com/kataras/iris/v12/core/router"
//)
//
//// 鉴权，将路由分发
//func RegisterAuthRouter(party router.Party) {
//	v1 := party.Party("/auth/v1")
//	{
//		// TODO change url
//		v1.Handle(http.MethodGet, "/verify/token", wrapper.Handler(v1_controller.AuthController{}.VerifyToken))
//		v1.Handle(http.MethodGet, "/test/header", wrapper.Handler(v1_controller.AuthController{}.TestHeader))
//	}
//}

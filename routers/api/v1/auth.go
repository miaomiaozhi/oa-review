package v1

//// 鉴权，将路由分发
//func RegisterAuthRouter(party router.Party) {
//	v1 := party.Party("/register")
//	{
//		// TODO change url
//		v1.Handle(http.MethodGet, "/verify/token", wrapper.Handler(v1_controller.AuthController{}.VerifyToken))
//		v1.Handle(http.MethodGet, "/test/header", wrapper.Handler(v1_controller.AuthController{}.TestHeader))
//	}
//}

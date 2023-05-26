package v1

//// 鉴权，将路由分发
//func RegisterReviewerRouter(party router.Party) {
//	v1 := party.Party("/reviewer")
//	{
//		v1.Handle(http.MethodGet, "/register", wrapper.Handler(v1_controller.AuthController{}.VerifyToken))
//	}
//}

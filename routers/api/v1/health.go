package v1

import (
	"net/http"
	v1_controller "oa-review/controllers/v1"
	"oa-review/internal/wrapper"
	"oa-review/logger"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterHealthRouter(party router.Party) {
	v1 := party.Party("/health")
	{
		v1.Handle(http.MethodGet, "/", wrapper.Handler(v1_controller.HealthController{}.Test))
	}
	logger.Info("register health router success")
}

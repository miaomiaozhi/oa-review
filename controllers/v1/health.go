package v1

import (
	wrapper "oa-review/internal/wrapper"
	"oa-review/logger"
	v1 "oa-review/services/v1"
)

type HealthController struct {
}

func (HealthController) Test(ctx *wrapper.Context) {
	logger.Info("health controller wrapper Test")
	wrapper.ApiWrapper(ctx, v1.Test, true, nil, nil)
}

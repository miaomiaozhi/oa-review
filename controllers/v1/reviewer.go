package v1

import (
	"oa-review/internal/wrapper"
	"oa-review/logger"
	v1 "oa-review/services/v1"
)

type ReviewerController struct {
}

func (ReviewerController) Test(ctx *wrapper.Context) {
	logger.Info("health controller wrapper Test")
	wrapper.ApiWrapper(ctx, v1.Test, true, nil, nil)
}

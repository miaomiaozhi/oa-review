package v1

import (
	"oa-review/internal/wrapper"
	"oa-review/logger"
	v1 "oa-review/models/protoreq/v1"
	services_v1 "oa-review/services/v1"
)

type ReviewerController struct {
}

func (ReviewerController) ReviewerSubmit(ctx *wrapper.Context) {
	logger.Info("reviewer controller wrapper submit")
	wrapper.ApiWrapper(ctx, services_v1.ReviewerSubmit, true, &v1.ReviewerSubmitRequest{}, nil)
}

func (ReviewerController) ReviewerWithDraw(ctx *wrapper.Context) {
	logger.Info("reviewer controller wrapper WithDraw")
	wrapper.ApiWrapper(ctx, services_v1.ReviewerWithDraw, true, &v1.ReviewerWithDrawRequest{}, nil)
}
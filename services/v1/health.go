package v1

import (
	"oa-review/internal/wrapper"
	"oa-review/logger"
)

func Test(ctx *wrapper.Context, reqBody interface{}) error {
	wrapper.SendApiOKResponse(ctx, nil, "hello")
	logger.Info("Apiwrapper ok")
	return nil
}

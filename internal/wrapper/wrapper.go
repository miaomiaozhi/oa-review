package wrapper

import (
	logger "oa-review/logger"

	"github.com/go-playground/validator"
)

var validate = validator.New()

type ApiHandler func(ctx *Context, reqBody interface{}) error

// ApiWrapper 传入 handler 跟请求，在这里进行参数校验等合法性检验，最后进行请求响应
func ApiWrapper(ctx *Context, handler ApiHandler, paramChecker bool, reqBody interface{}, params ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			// show panic error
			logger.Errorf("recover :%v", r)
			SendApiErrorResponse(ctx, nil, "内部错误")
		}
	}()

	if reqBody != nil {
		if len(params) == 0 {
			logger.Error("ApiWrapper 传入参数为空")
			SendApiBadRequestResponse(ctx, nil, "params empty")
			return
		}

		// 解析请求
		err := ctx.ReadJSON(reqBody)
		if err != nil {
			logger.Error("ApiWrapper 解析请求错误", err.Error())
			SendApiBadRequestResponse(ctx, nil, "parse request body error")
			return
		}

		// 参数校验
		if paramChecker {
			if err := checkParam(reqBody); err != nil {
				logger.Info("param checker error:", err.Error())
				SendApiBadRequestResponse(ctx, nil, "参数错误")
				return
			}
		}
	}
	err := handler(ctx, reqBody)
	if err != nil {
		SendApiErrorResponse(ctx, nil, "内部错误")
	}
}

func checkParam(reqBody interface{}) error {
	err := validate.Struct(reqBody)
	if err != nil {
		return err
	}
	return nil
}

package wrapper

import (
	"encoding/json"
	mlog "oa-review/logger"

	"github.com/kataras/iris/v12"
)

type Context struct {
	iris.Context

	UserName  string
	UserToken *AuthResult
	// TODO: logger info
}

type AuthResult struct {
	Certificate string `json:"certificate"` //认证
	UserID      string `json:"user_id"`     //用户id
	UserName    string `json:"user_name"`   //用户名
}

func Acquire(original iris.Context, login bool) *Context {
	// TODO add context to context pool
	ctx := &Context{
		Context: original,
	}
	// set the context to the original one in order to have access to iris's implementation.
	if login {
		ctx.UserToken = GetAutoResult(original)
		if ctx.UserToken == nil {
			ctx.StatusCode(401)
			ctx.StopExecution()
		} else {
			ctx.UserName = ctx.UserToken.UserName
		}
	}
	return ctx
}

func GetAutoResult(ctx iris.Context) *AuthResult {
	user := ctx.GetHeader("User")
	if user == "" {
		mlog.Error("header user invalid")
		return nil
	}
	authInfo := &AuthResult{}
	err := json.Unmarshal([]byte(user), &authInfo)
	if err != nil {
		mlog.Error("auth info json unmarshal failed", err.Error())
		return nil
	}
	mlog.Debugf("[AUTH_INFO] username: %v", authInfo.UserName)
	return authInfo
}

// Handler will convert our handler of func(*Context) to an iris Handler,
// in order to be compatible with the HTTP API.
func Handler(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := Acquire(original, true)
		if !ctx.IsStopped() { // 请求被终止
			h(ctx)
		}
		// TODO release
	}
}

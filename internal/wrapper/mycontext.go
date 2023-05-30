package wrapper

import (
	"encoding/json"
	logger "oa-review/logger"
	jwt "oa-review/middleware"

	"github.com/kataras/iris/v12"
)

type Context struct {
	iris.Context

	UserToken *AuthResult
	// TODO: logger info
}

type AuthResult struct {
	Token    string `json:"Token"`    // token
	UserID   int64  `json:"UserId"`   //用户id
	UserName string `json:"UserName"` //用户名
	Priority int64  `json:"Priority"` // 优先级
}

// Acquire 将 iris 的context转化成自定义的context
// 并且做授权验证
func Acquire(original iris.Context, login bool) *Context {
	// TODO add context to context pool
	ctx := &Context{
		Context: original,
	}
	// set the context to the original one in order to have access to iris's implementation.
	if login {
		ctx.UserToken = GetAuthResult(original)
		if ctx.UserToken == nil {
			ctx.StatusCode(401)
			ctx.StopExecution()
		} else {
			claim, err := jwt.ParseJwtToken(ctx.UserToken.Token)
			if err != nil {
				ctx.StatusCode(401)
				ctx.StopExecution()
			} else {
				userId, ok1 := claim["UserId"].(float64)
				userName, ok2 := claim["UserName"].(string)
				priority, ok3 := claim["Priority"].(float64)

				// 不合法
				if !ok1 || !ok2 || !ok3 ||
					int64(userId) != ctx.UserToken.UserID ||
					userName != ctx.UserToken.UserName ||
					int64(priority) != ctx.UserToken.Priority {
					ctx.StatusCode(401)
					ctx.StopExecution()
				}
			}
		}
	}
	return ctx
}

func GetAuthResult(ctx iris.Context) *AuthResult {
	user := ctx.GetHeader("User")
	if user == "" {
		logger.Error("header user invalid")
		return nil
	}
	// 得到授权信息
	authInfo := &AuthResult{}
	err := json.Unmarshal([]byte(user), &authInfo)
	if err != nil {
		logger.Error("auth info json unmarshal failed", err.Error())
		return nil
	}
	logger.Debug("user auth token ")
	return authInfo
}

// 处理需要登录请求
func Handler(handler func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := Acquire(original, true)
		if !ctx.IsStopped() { // 请求被终止
			handler(ctx)
		}
		// TODO release
	}
}

// 处理无需登录的请求
func HandlerNotLogin(handle func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := Acquire(original, false)
		if !ctx.IsStopped() { // 请求被终止
			handle(ctx)
		}
		// TODO release
	}
}

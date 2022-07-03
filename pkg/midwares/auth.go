package middleware

import (
	apibase "4clabs-server/api/base/v1"
	"4clabs-server/pkg/auth"
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"time"
)

type Conf struct {
	JwtKey     string
	MaxRefresh time.Duration
	MaxTime    time.Duration
}

func GetAuthMiddle(validator *auth.JwtUtils, contextUidSetter *auth.ContextUtils, inPath map[string]struct{}) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 不需要鉴权
				if _, ok := inPath[tr.Operation()]; !ok {
					return handler(ctx, req)
				}
				jwtToken := tr.RequestHeader().Get("Authorization")
				if jwtToken == "" {
					// 缺少可认证的token，返回错误
					return nil, apibase.ErrorNoAuthHeader("需要登录后使用")
				}
				add, token, err := validator.ValidateTokenWithAutoRefresh(jwtToken)
				if err != nil {
					return nil, apibase.ErrorAuthTokenExpired("登录已过期")
				}
				tr.ReplyHeader().Set("refresh-token", jwtToken)
				ctx = contextUidSetter.SetJwtToken(contextUidSetter.SetAddress(ctx, add), token)
				return handler(ctx, req)
			} else {
				return nil, apibase.ErrorUnknownError("")
			}
		}
	}
}

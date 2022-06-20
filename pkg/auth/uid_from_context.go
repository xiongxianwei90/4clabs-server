package auth

import (
	"context"
)

type ContextUtils struct {
	addressKey string
	jwtKey string
}

func NewUtils() *ContextUtils {
	return &ContextUtils{addressKey: "x-md-global-address", jwtKey: "x-md-global-key"}
}

func (u *ContextUtils) SetAddress(ctx context.Context, address string) context.Context {
	return context.WithValue(ctx, u.addressKey, address)
}
func (u *ContextUtils) SetJwtToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, u.jwtKey, token)
}

func (u *ContextUtils) GetAddress(ctx context.Context) string {
	if address, ok := ctx.Value(u.addressKey).(string); ok {
		return address
	}
	return ""
}
func (u *ContextUtils) GetJwtToken(ctx context.Context) string {
	if token, ok := ctx.Value(u.jwtKey).(string); ok {
		return token
	}
	return ""
}

package usecase

import (
	apibase "4clabs-server/api/base/v1"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/geth"
	"4clabs-server/app/4clabs-server/internal/ports"
	"4clabs-server/pkg/auth"
	"context"
	"fmt"
)

type Auth struct {
	authRepo ports.Auth
	jwtUtils *auth.JwtUtils
}

func NewAuth(authRepo ports.Auth, jwtUtils *auth.JwtUtils) *Auth {
	return &Auth{authRepo: authRepo, jwtUtils: jwtUtils}
}

func (a *Auth) GetSignMessage(ctx context.Context, address string) (msg string, err error) {
	return a.getMsg(ctx, address)
}

func (a *Auth) SignToLogin(ctx context.Context, address, sign, message string) (token string, err error) {
	if !geth.VerifySig(address, sign, []byte(message)) {
		return "", apibase.ErrorSignVerifyFailed("签名验证失败")
	}
	if err = a.authRepo.RefreshNonce(ctx, address); err != nil {
		return
	}
	return a.jwtUtils.NewToken(address)
}

func (a *Auth) getMsg(ctx context.Context, address string) (string, error) {
	nonce, err := a.authRepo.GetNonce(ctx, address)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("welcom to 4clabs! nonce : %s", nonce), nil
}

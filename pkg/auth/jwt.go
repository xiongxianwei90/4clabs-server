package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtUtils struct {
	Secret         string
	Exp            int64
	RefreshTimeGap int64
}

type CustomClaims struct {
	Address string
	*jwt.StandardClaims
}

func NewJwtUtils(secret string, exp, refreshTimeGap int64) *JwtUtils {
	return &JwtUtils{
		Secret:         secret,
		Exp:            exp,
		RefreshTimeGap: refreshTimeGap,
	}
}
func (u *JwtUtils) ValidateTokenWithAutoRefresh(tokenStr string) (string, string, error) {
	var err error
	var jwtToken *jwt.Token
	if jwtToken, err = jwt.ParseWithClaims(tokenStr, &CustomClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + u.Exp,
		},
	}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(u.Secret), nil
	}); err != nil {
		return "", "", err
	}
	if cc, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
		// 看下 token 是否快过期了，如果快过期了，就要生成一个新的给客户端使用
		s := jwtToken.Claims.(*CustomClaims)
		if s.ExpiresAt-time.Now().Unix() < u.RefreshTimeGap {
			newTokenStr, err := u.NewToken(cc.Address)
			if err != nil {
				return "", "", err
			}
			return cc.Address, newTokenStr, nil
		}
		return cc.Address, tokenStr, nil
	}
	return "", "", fmt.Errorf("valid token : %s", tokenStr)
}
func (u *JwtUtils) NewToken(address string) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + u.Exp,
		},
		Address: address,
	})
	newTokenStr, err := newToken.SignedString([]byte(u.Secret))
	if err != nil {
		return "", err
	}
	return newTokenStr, nil
}

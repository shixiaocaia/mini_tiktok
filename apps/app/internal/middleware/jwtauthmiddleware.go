package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"mini_tiktok/apps/app/internal/config"
	"mini_tiktok/pkg/common/result"
	"net/http"
)

type JwtAuthMiddleware struct {
	Config config.Config
}

type UserClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

const (
	tokenEmpty   = "token is empty"
	tokenExpired = "token is expired"
	tokenInvalid = "token is invalid"
)

func NewJwtAuthMiddleware(c config.Config) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{Config: c}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")

		if token == "" {
			result.JwtUnauthorizedResult(w, r, fmt.Errorf(tokenEmpty))
			return
		}

		_, err := ParseToken(token, m.Config.JwtAuth.AccessSecret)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				result.JwtUnauthorizedResult(w, r, fmt.Errorf(tokenExpired))
			} else {
				result.JwtUnauthorizedResult(w, r, fmt.Errorf(tokenInvalid))
			}
			return
		}
		next(w, r)
	}
}

func ParseToken(tokenString string, secret string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if token == nil {
		return nil, err
	}

	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, err
	}
	return userClaims, err
}

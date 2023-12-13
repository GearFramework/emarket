package auth

import (
	"context"
	"errors"
	"github.com/GearFramework/emarket/internal/models"
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const (
	TokenExpired = time.Hour * 24
	SecretKey    = "bu7HBJD&873HVHJdh*Jbhsfdfs8622Dsf"
)

var ErrNeedAuthorization = errors.New("требуется авторизация")
var ErrInvalidAuthorization = errors.New("отсутствует UID сессии пользователя")

type Auth struct {
	TokenExpired time.Duration
	SecretKey    string
	Logger       *alog.Alog
	Storage      models.Storable
}

type Claims struct {
	jwt.RegisteredClaims
	SessionUID string
}

func (a Auth) IdentityByCookie(cookie *http.Cookie) (string, error) {
	if cookie == nil {
		return "", ErrInvalidAuthorization
	}
	return "", nil
}

func (a Auth) Login(ctx context.Context, r *models.RequestLogin) {

}

//
//func (a *Auth) BuildJWT(sessionID string) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpired)),
//		},
//		SessionUID: sessionID,
//	})
//	tk, err := token.SignedString([]byte(SecretKey))
//	if err != nil {
//		return "", err
//	}
//	return tk, nil
//}
//
//func (a *Auth) GetSessionUIDFromJWT(tk string) string {
//	claims, err := getClaims(tk)
//	if err != nil {
//		return ""
//	}
//	a.logger.Infof("app session UID: %d", claims.SessionUID)
//	return claims.SessionUID
//}
//
//func (a *Auth) getClaims(tk string) (*Claims, error) {
//	claims := &Claims{}
//	token, err := jwt.ParseWithClaims(tk, claims,
//		func(t *jwt.Token) (interface{}, error) {
//			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
//			}
//			return []byte(SecretKey), nil
//		})
//	if err != nil || !token.Valid {
//		a.logger.Error(err.Error())
//		return nil, err
//	}
//	return claims, nil
//}

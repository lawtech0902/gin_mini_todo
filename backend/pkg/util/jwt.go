package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

func GenerateToken(key, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"key": key,
		// * 生效时间
		"nbf": time.Now().Unix(),
		// * 签发时间
		"iat": time.Now().Unix(),
		// * 过期时间
		// "exp":      time.Now().Add(time.Hour * 2).Unix()
	})
	
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func ParseToken(token, secret string) error {
	t, err := jwt.Parse(token, secretFunc(secret))
	if err != nil {
		return err
	} else if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return nil
	} else {
		return err
	}
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		
		return []byte(secret), nil
	}
}

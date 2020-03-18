package common

import (
	"github.com/dgrijalva/jwt-go"
	"irisProject/model"
	"time"
)

var jwtKey = []byte("a_secret_caret")

//需要转换成token的结构
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

//生成 token 函数
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * time.Hour * 24).Unix() //过期时间 7*24
	claims := Claims{UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "MrWusenshan",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//解析 token 函数
func ParesToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}

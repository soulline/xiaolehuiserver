package tokenutil

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SecretKey = "xiaolehui_kslawer"
)

/**
* 生成token
 */
func NewToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

/**
* 校验token是否合法
 */
func ParseToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretKey"), nil
	})
	fmt.Println("token : ", token)
	return token.Valid, err
}

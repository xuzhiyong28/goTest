package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	Demo1()
}

func TestDemo2(t *testing.T) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: "xuzhiyong",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 最终生成的签名
	tokenString, err := token.SignedString(jwtKey)
	if err == nil {
		fmt.Println(tokenString)
	}
}

func TestDemo3(t *testing.T) {

}

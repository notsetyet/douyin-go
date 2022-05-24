// Package jwt 生成和获取用户的 token
package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtKey = []byte("byteDance") // 硬编码
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// SetToken 颁发 token 给用户
func SetToken(ctx *gin.Context, userID uint) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil
	}
	str = tokenString
	//ctx.JSON(200, gin.H{"token": tokenString})
	return tokenString, nil
}

// GetToken 解析 token
func GetToken(ctx *gin.Context, method uint8) (int64, error) {
	var tokenString string
	// 根据 method 方法判断是 GET 请求还是 POST 请求：0 -- GET		1 -- POST
	if method == 0 {
		tokenString = ctx.Query("token")
	} else {
		tokenString = ctx.PostForm("token")
	}
	if tokenString == "" {
		return -1, errors.New("权限不足")
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return -1, errors.New("权限不足")
	}
	return int64(claims.UserId), nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, Claims, err
}

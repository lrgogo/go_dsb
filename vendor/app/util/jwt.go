package util

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"app/config"
	"time"
	"errors"
)

func verifyJWT(str string) int64 {
	t, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWT_SECRET_KEY), nil
	})
	CheckError(err)

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		uid := int64(claims["uid"])
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			CheckError(errors.New("token exp is wrong"))
		}
		return uid
	}
	return -1
}

func createAccessJWT(uid int64) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(), //3天
	})
	token, err := t.SignedString([]byte(config.JWT_SECRET_KEY))
	CheckError(err)

	return token
}

func createRefreshJWT(uid int64) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), //30天
	})
	token, err := t.SignedString([]byte(config.JWT_SECRET_KEY))
	CheckError(err)

	return token
}

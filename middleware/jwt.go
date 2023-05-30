package middleware

import (
	"fmt"
	"oa-review/conf"
	"oa-review/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenJwtToken(userId int64, userName string, priority int64) (string, error) {
	// 创建一个JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":   userId,
		"UserName": userName,
		"Priority": priority,
		"iat":      time.Now().Unix(),
	})

	// 签名JWT token
	jwtSecret := conf.GetConfig().Conf.MustGetString("web.jwt_secret")
	logger.Debug("jwt secret", jwtSecret)
	// jwtSecret := "mozezhao"
	secret := []byte(jwtSecret)
	logger.Debug("jwt secret", jwtSecret)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		logger.Error("gen jwt token failed", err.Error())
		return "", err
	}
	logger.Debug("gen jwt token success")
	logger.Debug("token:", signedToken)
	return signedToken, nil
}

func ParseJwtToken(jwtToken string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Errorf("parse jwt token failed: %v", fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := conf.GetConfig().Conf.MustGetString("web.jwt_secret")
		// jwtSecret := "mozezhao"
		secret := []byte(jwtSecret)
		return secret, nil
	})

	logger.Debug("token", jwtToken)
	if err != nil {
		return nil, err
	}

	// 提取JWT token的载荷信息
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// parse args
		_, ok1 := claims["UserId"].(float64)
		_, ok2 := claims["UserName"].(string)
		_, ok3 := claims["Priority"].(float64)
		if !ok1 || !ok2 || !ok3 {
			return nil, fmt.Errorf("parse jwt token failed: parsedToken invalid")
		}
		logger.Info("parse token ok")
		return claims, nil
	}
	return nil, fmt.Errorf("parse jwt token failed: parsedToken invalid")
}

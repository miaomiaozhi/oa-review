package pkg

import (
	"errors"
	"fmt"
	"log"
	dao "oa-review/dao"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	mySerectCode = "mozezhao"
)

// Authorize is a middleware that verifies the User JWT token in the Authorization header.
func UserAuthorize(tokenStr string) (int64, int32, error) {
	if tokenStr == "" {
		return errorMsg("user authorize token empty")
	}

	// handle authrization token
	log.Println("Authorization token", tokenStr)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the signing method is HMAC and that the secret key matches.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(mySerectCode), nil
	})
	if err != nil {
		return errorMsg(fmt.Sprintf("jwt token parse error :%v", err.Error()))
	}

	// Extract the user ID from the token.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errorMsg(fmt.Sprintf("extract the user ID from the token error :%v", err.Error()))
	}
	// Check time
	if exp := claims["exp"].(float64); !time.Unix(int64(exp), 0).After(time.Now()) {
		return errorMsg(fmt.Sprintf("jwt token expired. error"))
	}
	log.Println("claims", claims)

	userIdTmp := claims["userId"].(float64)
	userId := int64(userIdTmp)
	if int64(userId) == int64(0) {
		return errorMsg("lack of user id")
	}

	// Get the user from the database.
	exist, err := dao.NewUserDaoInstance().CheckUserExist(userId)
	if err != nil {
		return errorMsg(fmt.Sprintf("check user exist error: %v", err.Error()))
	}

	if !exist {
		return errorMsg("user doesn't exist")
	}
	return userId, 0, nil
}

// Authorize is a middleware that verifies the User JWT token in the Authorization header.
func ReviewerAuthorize(tokenStr string) (int64, int32, error) {
	if tokenStr == "" {
		return errorMsg("user authorize token empty")
	}

	// handle authrization token
	log.Println("Authorization token", tokenStr)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the signing method is HMAC and that the secret key matches.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(mySerectCode), nil
	})
	if err != nil {
		return errorMsg(fmt.Sprintf("jwt token parse error :%v", err.Error()))
	}

	// Extract the user ID from the token.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errorMsg(fmt.Sprintf("extract the user ID from the token error :%v", err.Error()))
	}
	// Check time
	if exp := claims["exp"].(float64); !time.Unix(int64(exp), 0).After(time.Now()) {
		return errorMsg(fmt.Sprintf("jwt token expired. error"))
	}
	log.Println("claims", claims)

	userIdTmp := claims["userId"].(float64)
	userId := int64(userIdTmp)
	if int64(userId) == int64(0) {
		return errorMsg("lack of user id")
	}

	// Get the user from the database.
	priorityTmp := claims["priority"].(float64)
	priority := int32(priorityTmp)
	if priority == int32(0) {
		return errorMsg("lack of priority")
	}

	// Get the user from the database.
	user, err := dao.NewUserDaoInstance().FindUserByUserId(userId)
	if err != nil {
		return errorMsg(fmt.Sprintf("error on find user: %v", err.Error()))
	}
	if user.Priority != priority {
		return errorMsg("no authorization to review")
	}
	return userId, priority, nil
}

func CreateUserJwtToken(userId int64, priority int32) (string, error) {
	// Create a new token object.
	token := jwt.New(jwt.SigningMethodHS256)

	// Set token claims.
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = int64(userId)
	claims["priority"] = int32(priority)
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and return it as a string.
	tokenString, err := token.SignedString([]byte(mySerectCode))
	if err != nil {
		log.Printf("Error on create user jwt token: %v\n", err)
		return "", err
	}
	return tokenString, nil
}

func errorMsg(errMsg string) (int64, int32, error) {
	log.Println(errMsg)
	return -1, -1, errors.New(errMsg)
}

package pkg

import (
	"errors"
	"log"
	dao "oa-review/dao"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

const (
	mySerectCode = "mozezhao"
)

func unauthorizedResponse(ctx iris.Context, errMsg string) {
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(iris.Map{
		"message": errMsg,
	})
}

// authorize is a middleware that verifies the User JWT token in the Authorization header.
func Authorize(ctx iris.Context) {
	// Get authorization header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		unauthorizedResponse(ctx, "empty authorization")
		return
	}

	// handle authrization token
	log.Println("Authorization header", authHeader)
	// Parse the token from the header.
	tokenString := authHeader[len("Bearer "):]
	log.Println("Authorization token", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the signing method is HMAC and that the secret key matches.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(string("can not parse user jwt token"))
		}
		return []byte(mySerectCode), nil
	})
	if err != nil {
		unauthorizedResponse(ctx, "invalid token")
		return
	}

	// Extract the user ID from the token.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		unauthorizedResponse(ctx, "invalid token")
		return
	}

	userId, ok := claims["userId"].(int64)
	if !ok {
		unauthorizedResponse(ctx, "invalid token")
		return
	}
	priority, ok := claims["priority"].(int32)
	if !ok {
		unauthorizedResponse(ctx, "invalid token")
		return
	}

	// Get the user from the database.
	user, err := dao.NewUserDaoInstance().FindUserByUserId(userId)
	if err != nil {
		unauthorizedResponse(ctx, err.Error())
		return
	}
	if user.UserId != userId || user.Priority != priority {
		unauthorizedResponse(ctx, "invalid token")
		return
	}

	// Store the user in the request context for use in subsequent handlers.
	ctx.Values().Set("user", user)

	// Call the next handler.
	ctx.Next()
}

func CreateUserJwtToken(userId int64) (string, error) {
	// Create a new token object.
	token := jwt.New(jwt.SigningMethodHS256)

	user, err := dao.NewUserDaoInstance().FindUserByUserId(userId)
	if err != nil {
		return "", err
	}

	// Set token claims.
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = int64(user.UserId)
	claims["priority"] = int32(user.Priority)
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and return it as a string.
	tokenString, err := token.SignedString([]byte(mySerectCode))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

// User represents a user in the system.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Credentials represents user credentials for login.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app := iris.New()

	// Add middleware for logging and recovering from panics.
	app.Use(recover.New())
	app.Use(logger.New())

	// Handle user login.
	app.Post("/login", handleLogin)

	// Authorized endpoints.
	users := app.Party("/users", authorize)
	users.Get("/{id:int}", handleGetUser)

	// Start the server.
	app.Listen(":8080")
}

// handleLogin handles user login requests.
func handleLogin(ctx iris.Context) {
	var creds Credentials
	err := ctx.ReadJSON(&creds)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid request body")
		return
	}

	// Check if the user exists and the password is correct.
	user, err := authenticateUser(creds.Username, creds.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("Invalid username or password")
		return
	}

	// Create JWT token.
	token, err := createToken(user.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Failed to create token")
		return
	}

	// Return token to the client.
	ctx.JSON(map[string]string{
		"token": token,
	})
}

// handleGetUser handles requests to get user information.
func handleGetUser(ctx iris.Context) {
	userID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid user ID")
		return
	}

	// Get the authenticated user.
	user := ctx.Values().Get("user").(*User)

	// Check if the authenticated user has permission to access the requested user.
	if user.ID != userID {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("You are not authorized to access this user")
		return
	}

	// Return user information to the client.
	ctx.JSON(user)
}

// authenticateUser authenticates a user with the given username and password.
func authenticateUser(username, password string) (*User, error) {
	// In a real application, this would usually involve checking the user's
	// credentials against a database or other authentication system.
	// For simplicity, we'll just hard-code a user with username "admin" and password "password".
	if username == "admin" && password == "password" {
		return &User{
			ID:       1,
			Username: "admin",
			Password: "",
		}, nil
	}
	return nil, errors.New("cant authenticate user")
}

// createToken creates a JWT token for the given user ID.
func createToken(userID int) (string, error) {
	// Create a new token object.
	token := jwt.New(jwt.SigningMethodHS256)

	// Set token claims.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and return it as a string.
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// authorize is a middleware that verifies the JWT token in the Authorization header.
func authorize(ctx iris.Context) {
	// Get the Authorization header.
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("Missing Authorization header")
		return
	}
	fmt.Println("token", authHeader)
	// Parse the token from the header.
	tokenString := authHeader[len("Bearer "):]
	fmt.Println("token", tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the signing method is HMAC and that the secret key matches.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("cant authenticate user")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("Invalid token")
		return
	}

	// Extract the user ID from the token.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("Invalid token")
		return
	}
	userID, ok := claims["sub"].(float64)
	if !ok {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("Invalid token")
		return
	}

	// Get the user from the database.
	user, err := getUserByID(int(userID))
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("Invalid token")
		return
	}

	// Store the user in the request context for use in subsequent handlers.
	ctx.Values().Set("user", user)

	// Call the next handler.
	ctx.Next()
}

// getUserByID gets a user from the database by ID.
func getUserByID(userID int) (*User, error) {
	// In a real application, this would usually involve retrieving the user's
	// information from a database or other data store.
	// For simplicity, we'll just hard-code a user with ID 1.
	if userID == 1 {
		return &User{
			ID:       1,
			Username: "admin",
			Password: "",
		}, nil
	}
	return nil, errors.New("cant authenticate user")
}

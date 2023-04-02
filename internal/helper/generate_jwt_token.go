package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"time"
)

func GenerateJWT(user *entity.User, key string, expiredMinute int) (string, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(time.Duration(expiredMinute) * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &entity.Claims{
		UserId:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {

		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, nil
}

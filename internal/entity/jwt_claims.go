package entity

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

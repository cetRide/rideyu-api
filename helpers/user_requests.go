package helpers

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Username string
	Password string
	jwt.StandardClaims
}

type GeneratedToken struct {
	Token string `json:"token"`
}

type AccountDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

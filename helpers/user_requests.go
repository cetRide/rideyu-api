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



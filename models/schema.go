package models

type User struct {
	Base
	Username string
	Email string
	Password string
}
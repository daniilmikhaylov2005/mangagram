package models

import "github.com/golang-jwt/jwt/v4"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

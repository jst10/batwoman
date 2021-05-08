package structs

import "github.com/dgrijalva/jwt-go"

type TokenData struct {
	UserId    uint   `json:"userId"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	SessionId uint   `json:"session_id"`
	jwt.StandardClaims
}

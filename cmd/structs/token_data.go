package structs

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenData struct {
	UserId    uint   `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
	Username  string `json:"username"`
	SessionId uint   `json:"session_id"`
	jwt.StandardClaims
}

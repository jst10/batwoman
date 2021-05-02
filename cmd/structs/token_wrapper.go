package structs

import (
	"made.by.jst10/celtra/batwoman/cmd/database"
	"time"
)

type TokenWrapper struct {
	User       *database.User
	Token      string
	Expiration *time.Time
}

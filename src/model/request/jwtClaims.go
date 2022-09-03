package request

import (
	"github.com/golang-jwt/jwt/v4"
)

type BaseClaims struct {
	UID      int64
	Username string
	Nickname string
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

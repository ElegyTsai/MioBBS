package utils

import (
	"Miyo_Assignment/global"
	"Miyo_Assignment/model/request"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired  = errors.New("Token is expired")
	TokenNotValid = errors.New("Token is not active yet")
	TokenInvalid  = errors.New("Couldn't handle this token")
)

//var (
//	once sync.Once
//	jwt *JWT
//)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GV_Config.Jwt.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.GV_Config.Jwt.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: &jwt.NumericDate{time.Now()},
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Microsecond * time.Duration(global.GV_Config.Jwt.ExpireTime))},
			Issuer:    global.GV_Config.Jwt.Issuer,
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.SigningKey)
}

//更新token
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	//global用来进行并发控制，保证一个用户只有一个token
	newToken, err, _ := global.GV_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return newToken.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if validError, ok := err.(*jwt.ValidationError); ok {
			if validError.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if validError.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValid
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}

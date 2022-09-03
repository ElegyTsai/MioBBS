package middleware

import (
	"Miyo_Assignment/global"
	"Miyo_Assignment/model/request"
	"Miyo_Assignment/model/response"
	"Miyo_Assignment/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

func JWTAUth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithAll(ctx, gin.H{"reload": true}, "身份验证失败，未登录")
			ctx.Abort()
			return
		}
		j := utils.NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithAll(ctx, gin.H{"reload": true}, "登陆过期，请重新登陆")
				ctx.Abort()
				return
			} else if err == utils.TokenNotValid {
				response.FailWithAll(ctx, gin.H{"reload": true}, "令牌尚未启用")
				ctx.Abort()
				return
			} else if err == utils.TokenInvalid {
				response.FailWithAll(ctx, gin.H{"reload": true}, "令牌无效")
				ctx.Abort()
				return
			}
		}
		if claims.ExpiresAt.Unix()-time.Now().Unix() < global.GV_Config.Jwt.BufferTime {
			updateNewToken(ctx, claims, j, token)
		}
		//更新token
		ctx.Next()

	}
}

func updateNewToken(c *gin.Context, claims *request.CustomClaims, j *utils.JWT, token string) {
	claims.ExpiresAt = &jwt.NumericDate{time.Now().Add(time.Microsecond * time.Duration(global.GV_Config.Jwt.ExpireTime))}
	newToken, _ := j.CreateTokenByOldToken(token, *claims)
	newClaims, _ := j.ParseToken(newToken)
	c.Header("new-token", newToken)
	c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
	//附上了新的token
	c.Set("claims", newClaims)

}

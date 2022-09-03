package middleware

import (
	"Miyo_Assignment/model/response"
	"Miyo_Assignment/service"
	"Miyo_Assignment/utils"
	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.CasbinService

func AuthenticationHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		method := c.Request.Method
		claims, err := utils.GetClaims(c)
		if err != nil {
			response.FailWithAll(c, gin.H{"reload": true}, "令牌无效")
			c.Abort()
		}
		userID := claims.UID
		e := casbinService.Casbin()
		success, _ := e.Enforce(userID, url, method)

		if !success {
			response.FailWithAll(c, gin.H{}, "权限不足")
			c.Abort()
		}
		c.Next()
	}
}

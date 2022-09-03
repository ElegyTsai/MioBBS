package utils

import (
	"Miyo_Assignment/global"
	"Miyo_Assignment/model/request"
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := c.GetHeader("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GV_LOG.Errorf("解析用户Token失败")
		return nil, err
	}
	return claims, err

}

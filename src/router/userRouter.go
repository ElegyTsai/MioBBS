package router

import (
	v1 "Miyo_Assignment/api/v1"
	"github.com/gin-gonic/gin"
)

func (s *ApiRouter) UserRouterInit(router *gin.RouterGroup) {
	userRouter := router.Group("user")

	systemApi := v1.ApiGroupApp.SystemApi
	{
		userRouter.POST("register", systemApi.Register)
		userRouter.POST("login", systemApi.Login)
	}

}

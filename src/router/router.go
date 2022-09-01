package router

import (
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func InitRouter() *gin.Engine {
	Router := gin.Default()
	//router.Use()
	apiRouter := &ApiRouter{}

	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GVA_LOG.Info("register swagger handler")
	publicRouter := Router.Group("")
	apiRouter.UserRouterInit(publicRouter)

	return Router
}

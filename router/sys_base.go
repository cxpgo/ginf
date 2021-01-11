package router

import (
	v1 "github.com/cxpgo/ginf/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong",})
		})
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}

	return BaseRouter
}

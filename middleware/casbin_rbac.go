package middleware

import (
	"fmt"
	"github.com/cxpgo/ginf/global"
	"github.com/cxpgo/ginf/model/request"
	"github.com/cxpgo/ginf/model/response"
	"github.com/cxpgo/ginf/service"
	"github.com/gin-gonic/gin"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		fmt.Printf("sub=%s,obj=%s,act=%s\n",sub,obj,act)
		e := service.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.Config.System.Env == "dev" || success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}

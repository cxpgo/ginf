package v1

import (
	"fmt"
	"github.com/cxpgo/ginf/model"
	"github.com/cxpgo/ginf/model/request"
	"github.com/cxpgo/ginf/model/response"
	"github.com/cxpgo/ginf/service"
	"github.com/cxpgo/ginf/utils"
	"github.com/cxpgo/golib/lib"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			lib.Log.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			//tokenNext(c, *user)
			fmt.Printf("user:%+v",user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

package inits

import (
	"fmt"
	_ "github.com/cxpgo/ginf/docs"
	"github.com/cxpgo/ginf/global"
	"github.com/cxpgo/ginf/router"
	"github.com/cxpgo/golib/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := router.Routers()
	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.Config.System.Post)
	s := initServer(address, Router)

	// 保证欢迎信息文本最后顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	lib.Log.Info("server run success on "+address)
	fmt.Printf(`
	欢迎使用 ginf
	当前版本:V0.0.1
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html

`, address)
	lib.Log.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

package core

import (
	"context"
	"fmt"
	_ "github.com/cxpgo/ginf/docs"
	"github.com/cxpgo/ginf/global"
	"github.com/cxpgo/ginf/router"
	"github.com/cxpgo/golib/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server interface {
	ListenAndServe() error
}

var HttpSrvHandler *http.Server

func RunWindowsServer() {

	//*gin.Engine
	router := router.Routers()
	//静态文件
	//Router.Static("/form-generator", "./resource/page")

	//服务器地址
	address := fmt.Sprintf(":%d", global.Config.System.Post)

	//启动httpserver
	HttpSrvHandler = HttpServerRun(router, address)

	//显示欢迎文字
	showWecomeText(address)

	//优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	HttpServerStop(HttpSrvHandler)
}

func HttpServerRun(router *gin.Engine, address string) (httpServer *http.Server) {
	gin.SetMode(lib.GConfig.Http.GinModel)
	httpServer = InitServer(router, address)
	go func() {
		lib.Log.Printf(" [INFO] HttpServerRun:%s", address)
		if err := httpServer.ListenAndServe(); err != nil {
			lib.Log.Errorf(" [ERROR] HttpServerRun:%s err:%v", address, err)
		}
	}()
	return
}

func HttpServerStop(s *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		lib.Log.Errorf(" [ERROR] HttpServerStop err:%v", err)
	}
	lib.Log.Printf(" [INFO] HttpServerStop stopped")
}

func InitServer(router *gin.Engine, address string, ) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    time.Duration(lib.GConfig.Http.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(lib.GConfig.Http.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.GConfig.Http.MaxHeaderBytes),
	}
}

func showWecomeText(address string) {
	// 保证欢迎信息文本最后顺序输出
	time.Sleep(10 * time.Microsecond)
	lib.Log.Info("server run success on " + address)
	fmt.Printf(`
	欢迎使用 ginf
	当前版本:V0.0.1
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html

`, address)
}

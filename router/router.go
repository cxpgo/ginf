package router

import (
	_ "github.com/cxpgo/ginf/docs"
	"github.com/cxpgo/ginf/middleware"
	"github.com/cxpgo/golib/lib"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"


)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	////web静态资源
	//Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	////https
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	lib.Log.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	lib.Log.Info("use middleware cors")
	//启动swagger 第一次运行1: swag init 2: import _ "xxx/xx/xx/docs"
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	lib.Log.Info("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用 ps:/server1/base /server2/base
	PublicGroup := Router.Group("")
	{
		InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	PrivateGroup.Use(middleware.CasbinHandler())
	{
		//InitApiRouter(PrivateGroup)                   // 注册功能api路由
		InitJwtRouter(PrivateGroup)                   // jwt相关路由
		InitUserRouter(PrivateGroup) // 注册用户路由
		InitMenuRouter(PrivateGroup) // 注册menu路由
		//router.InitEmailRouter(PrivateGroup)                 // 邮件相关路由
		//router.InitSystemRouter(PrivateGroup)                // system相关路由
		//router.InitCasbinRouter(PrivateGroup)                // 权限相关路由
		//router.InitWorkflowRouter(PrivateGroup)              // 工作流相关路由
		//router.InitCustomerRouter(PrivateGroup)              // 客户路由
		//router.InitAutoCodeRouter(PrivateGroup)              // 创建自动化代码
		//router.InitAuthorityRouter(PrivateGroup)             // 注册角色路由
		//router.InitSimpleUploaderRouter(PrivateGroup)        // 断点续传（插件版）
		//router.InitSysDictionaryRouter(PrivateGroup)         // 字典管理
		//router.InitSysOperationRecordRouter(PrivateGroup)    // 操作记录
		//router.InitSysDictionaryDetailRouter(PrivateGroup)   // 字典详情管理
		//router.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	}
	lib.Log.Info("router register success")
	return Router
}

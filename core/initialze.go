package core

import (
	"github.com/cxpgo/ginf/global"
	"github.com/cxpgo/golib/lib"
)

func Init() {
	initConfig()
}

func initConfig() {
	lib.InitConfig(getConfigPath(), global.Config)
	//fmt.Printf("config=%+v",global.Config)
}

func getConfigPath() (path string) {
	path = global.ConfigPath + lib.GConfig.System.Env + global.ConfigName
	return
}



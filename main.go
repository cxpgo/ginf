package main

import (
	"github.com/cxpgo/ginf/core"
	"github.com/cxpgo/golib/lib"
)
func main(){
	lib.InitAll()
	defer lib.Destroy()

	core.Init()
	core.RunWindowsServer()

}

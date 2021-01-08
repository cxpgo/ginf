package main

import (
	"github.com/cxpgo/ginf/inits"
	"github.com/cxpgo/golib/lib"
)
func main(){
	lib.InitAll()
	defer lib.Destroy()

	inits.Init()
	inits.RunWindowsServer()

}

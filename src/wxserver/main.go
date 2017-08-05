package main

import (
	_ "wxserver/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

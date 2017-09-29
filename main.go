package main

import (
	_ "quickstart/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/upload/clumn_logo", "./upload/clumn_logo")
	beego.Run()
}

package routers

import (
	"quickstart/controllers"
	"quickstart/controllers/admin"
	"quickstart/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")

	//后台路由
	ns := beego.NewNamespace("admin",
		beego.NSRouter("/", &admin.IndexHandle{}),
		beego.NSRouter("/login", &admin.LoginHandle{}, "*:Login"),
	)
	beego.AddNamespace(ns)

	//API Modules
	jns := beego.NewNamespace("api",
		beego.NSRouter("/", &api.IndexApiHandle{}),
	)

	beego.AddNamespace(jns)
}

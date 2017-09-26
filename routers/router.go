package routers

import (
	"quickstart/controllers"
	"quickstart/controllers/admin"
	"quickstart/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/?:page:int/", &controllers.MainController{})
	beego.Router("/i/?:id:int/", &controllers.MainController{}, "*:GetInfo")
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")

	beego.Router("/up/f", &controllers.ServiceUpController{}, "*:ToPost")

	//后台路由
	ns := beego.NewNamespace("admin",
		beego.NSRouter("/", &admin.IndexHandle{}),
		beego.NSRouter("/login", &admin.LoginHandle{}, "*:Login"),
		beego.NSRouter("/tube", &admin.IndexHandle{}, "*:Tube"),

		//栏目
		beego.NSRouter("/column/index", &admin.ClumnHandle{}, "*:Index"),
		beego.NSRouter("/column/add", &admin.ClumnHandle{}, "*:Add"),

		//相册
		beego.NSRouter("/p/i/?:page:int/", &admin.AlbumsHandle{}, "*:List"),
		beego.NSRouter("/p/e/?:page:int/:id:int/", &admin.AlbumsHandle{}, "*:Info"),
	)
	beego.AddNamespace(ns)

	//API Modules
	jns := beego.NewNamespace("api",
		beego.NSRouter("/", &api.IndexApiHandle{}),
	)

	beego.AddNamespace(jns)
}

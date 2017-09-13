package admin

import (
	"strconv"
	//"log"
	"quickstart/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	userid         int64
	username       string
	controllerName string
	actionName     string
	admindir       string
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()

	//去掉controller的后缀handle以及全部小写
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-6])
	//小写
	this.actionName = strings.ToLower(actionName)
	this.auth()
}

func (this *BaseController) auth() {
	this.admindir = beego.AppConfig.String("admindir")

	//登录和登出除外
	if this.actionName != "login" && this.actionName != "logout" {

		//把cookie信息转换成slice
		ck := strings.Split(this.Ctx.GetCookie("token"), "|")

		if ok := len(ck); ok == 2 {
			idStr, password := ck[0], ck[1]
			//反idStr转换成int64
			userId, _ := strconv.ParseInt(idStr, 10, 0)

			if userId > 0 {
				var adminInfo models.AdminInfo
				adminInfo.Id = userId

				//根据id读取一条数据
				if adminInfo.Read() == nil && password == models.Md5("Token|"+adminInfo.Password) {
					this.userid = adminInfo.Id
					this.username = adminInfo.Username
				}
			}

		}
		if this.userid == 0 {
			this.Redirect(this.admindir+"/login", 302)
		}
	}
}

/*
	获取访问者IP
*/
func (this *BaseController) getClientIP() string {
	return this.Ctx.Input.IP()
}

func (this *BaseController) getTime() time.Time {
	return time.Now()
}

//@desc:是否为Post请求
//@return boolear

func (this *BaseController) isPost() bool {
	return this.Ctx.Input.Method() == "POST"
}

package admin

import (
	"log"
	"strconv"
	"strings"

	"quickstart/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type IndexHandle struct {
	BaseController
}

type LoginHandle struct {
	BaseController
}

type FromData struct {
	UserName string `valid:"Required"`
	Passwrod string `valid:"Required;MaxSize(12);MinSize(6)"`
}

func (f *FromData) Valid(v *validation.Validation) {
	if strings.Index(f.UserName, "admin") != -1 {
		v.SetError("UserName:", "名称里不能包含admin信息")
	}
}

func (this *IndexHandle) Get() {
	this.Ctx.WriteString("Index")
	return
}

//后台登录
func (this *LoginHandle) Login() {
	/*var adminInfo models.AdminInfo
	adminInfo.Username = "wangpeng"
	adminInfo.Password = models.Md5("123456")
	adminInfo.Lastlogintime = this.getTime()
	adminInfo.Lastloginip = this.getClientIP()
	adminInfo.Addtime = this.getTime()
	adminInfo.Insert()
	return*/
	//初使化flash
	beego.ReadFromRequest(&this.Controller)

	if ok := this.isPost(); ok == true {
		//创造一个新的flash
		flash := beego.NewFlash()

		//获取Post过来的数据进行span去除
		userName := strings.TrimSpace(this.GetString("username"))
		pwd := strings.TrimSpace(this.GetString("password"))

		//validation struct给一个变量
		valid := validation.Validation{}
		f := FromData{userName, pwd}

		//得到数据后进行验证
		b, err := valid.Valid(&f)

		//是否读取错误
		if err != nil {
		}

		if !b {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
				msg := err.Key + ":" + err.Message
				flash.Error(msg)
				flash.Store(&this.Controller)
				this.Redirect("/admin", 302)
				return
			}
		}

		//数据验证通过
		var adminInfo models.AdminInfo
		adminInfo.Username, adminInfo.Password = userName, models.Md5(pwd)

		//账号密码错误
		if adminInfo.Read("username") != nil || adminInfo.Password != models.Md5(pwd) {
			flash.Error("账号或者密码错误.")
			flash.Store(&this.Controller)
			this.Redirect("/admin", 302)
			return
		}

		//验证通过准备放入到cookie中
		adminInfo.Logintimes += 1
		adminInfo.Lastloginip = this.getClientIP()
		adminInfo.Lastlogintime = this.getTime()
		adminInfo.Update()

		authKey := models.Md5("Token|" + adminInfo.Password)

		this.Ctx.SetCookie("token", strconv.FormatInt(int64(adminInfo.Id), 10)+"|"+authKey)
		this.Redirect("/admin", 302)
	}

	this.TplName = "admin/login.tpl"
}

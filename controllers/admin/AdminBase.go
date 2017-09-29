package admin

import (
	"math/rand"
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

//@desc:分页数据组合
//@parames:
//	pagesize->分页显示多少条数据
//	page->当前页数
//	counts->总共有多少条数据
//	first->是否显示第一页
//	path->分页路径
//@return string
func (this *BaseController) PagesList(pagesize, page, counts int64, first bool, path string) string {
	//判断总数为0
	if counts <= 0 {
		return ""
	}

	//根据条件计算出的一共为多少页
	var pagecount int64
	pagecount = 0

	//拿pagesize跟counts取模，如果等于0那就说明能整除
	if counts%pagesize == 0 {
		pagecount = counts / pagesize
	} else {
		//会有小数
		pagecount = (counts / pagesize) + 1
	}

	if pagecount < 2 {
		return "<ul class='pagination'><li><a href='javascript:;'><span>共1页</span></a></li></ul>"
	}

	//上一页
	prevPage := ""
	if page > 1 {
		tempP := path + "/" + strconv.FormatInt(page-1, 10)
		prevPage = "<ul class='pagination'><li><a href='" + tempP + "'>上一页</a></li></ul>"
	}

	var l string
	l += "<ul class='pagination pagination-group'>"
	if page < pagecount {
		var i int64
		for i = 1; i <= 5; i++ {
			if page+i == pagecount || page+i > pagecount {
				continue
			}
			tempC := path + "/" + strconv.FormatInt(page+i, 10)
			l += "<li><a href='" + tempC + "'>" + strconv.FormatInt(page+i, 10) + "</a></li>"
		}
	} else if page == pagecount || page > pagecount {
		l += "<li><a href='javascript:;'>最后一页</a></li>"
	}
	l += "</ul>"

	//下一页
	nextPage := ""
	if page < pagecount {
		tempN := path + "/" + strconv.FormatInt(page+1, 10)
		nextPage = "<ul class='pagination'><li><a href='" + tempN + "'>下一页</a></li></ul>"
	}

	return prevPage + l + nextPage
}

//生成随机字符串
func (this *BaseController) GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

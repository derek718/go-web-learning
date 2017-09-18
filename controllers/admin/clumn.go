package admin

import (
	//	"log"
	"quickstart/models"
	"strconv"
)

//	"github.com/astaxie/beego"

type ClumnHandle struct {
	BaseController
}

//栏目列表
func (this *ClumnHandle) Index() {
	var (
		clumn models.Clumn
		list  []*models.Clumn
	)
	query := clumn.Query()
	query = query.OrderBy("-id")
	query.All(&list, "id", "parentid", "clumn_name")
	this.Data["list"] = list
	//	this.Ctx.WriteString("heloo")
	this.TplName = "admin/clumn_index.tpl"
}

//添加栏目
func (this *ClumnHandle) Add() {
	//把所有一级栏目都列出来
	var (
		clumn models.Clumn
		list  []*models.Clumn
	)
	clumn.Query().Filter("parentid", 0).OrderBy("-id").All(&list, "id", "parentid", "clumn_name")

	if ok := this.isPost(); ok == true {

		clumnName := this.GetString("clumn_name")
		parentid := this.Input().Get("parentid")

		intparentid, _ := strconv.Atoi(parentid)
		clumn.ClumnName = clumnName
		clumn.Parentid = int64(intparentid)
		clumn.Insert()

		this.Redirect(this.admindir+"column/index", 302)
	}

	this.Data["list"] = list
	this.TplName = "admin/add.tpl"
}

//编辑栏目信息
func (this *ClumnHandle) Edit() {}

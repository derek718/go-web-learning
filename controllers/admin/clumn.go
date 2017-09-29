package admin

import (
	//	"log"
	"quickstart/models"
	"strconv"
)

type ClumnHandle struct {
	BaseController
}

//栏目列表
func (this *ClumnHandle) Index() {
	var (
		page     int64
		pagesize int64 = 6
		offset   int64
		pager    string
		clumn    models.Clumn
		list     []*models.Clumn
	)
	//拿到Get参数
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)

	if page < 1 {
		page = 1
	}

	//偏移量
	offset = (page - 1) * pagesize

	query := clumn.Query()
	counts, _ := query.Count()
	if counts > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PagesList(pagesize, page, counts, false, this.admindir+"column/index")

	this.Data["list"] = list
	this.Data["pager"] = pager
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
		tmpName := this.GetRandomString(12)
		clumnName := this.GetString("clumn_name")
		parentid := this.Input().Get("parentid")
		desc := this.GetString("desc")

		//获取文件
		f, _, _ := this.GetFile("uploadfile")
		f.Close()
		fileName := "./upload/clumn_logo/" + tmpName + ".jpg"
		this.SaveToFile("uploadfile", fileName)

		intparentid, _ := strconv.Atoi(parentid)
		clumn.ClumnName = clumnName
		clumn.Parentid = int64(intparentid)
		clumn.Desc = desc
		clumn.Thumbs = fileName
		clumn.Insert()

		this.Redirect(this.admindir+"column/index", 302)
	}

	this.Data["list"] = list
	this.TplName = "admin/add.tpl"
}

//编辑栏目信息
func (this *ClumnHandle) Edit() {
	var (
		m    models.Clumn
		list []*models.Clumn
		info models.Clumn
	)
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	query := m.Query()
	query.Filter("parentid", 0).OrderBy("-id").All(&list)
	query.Filter("id", id).One(&info)

	//处理post
	if ok := this.isPost(); ok {
		tmpName := this.GetRandomString(12)
		clumnName := this.GetString("clumn_name")
		parentidStr := this.Input().Get("parentid")
		parentid, _ := strconv.ParseInt(parentidStr, 10, 64)
		desc := this.GetString("desc")

		fileName := info.Thumbs
		//获取文件
		f, _, err := this.GetFile("uploadfile")
		if err == nil {
			f.Close()
			fileName = "./upload/clumn_logo/" + tmpName + ".jpg"
			this.SaveToFile("uploadfile", fileName)
		}

		//读取数据
		info.ClumnName = clumnName
		info.Desc = desc
		info.Parentid = parentid
		info.Thumbs = fileName
		info.Update()
		this.Redirect(this.admindir+"column/index", 302)
	}

	this.Data["list"] = list
	this.Data["info"] = info
	this.TplName = "admin/add.tpl"
}

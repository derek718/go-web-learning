package admin

import (
	"quickstart/models"
	"strconv"
)

type TagHandle struct {
	BaseController
}

//标签列表
func (this *TagHandle) Index() {
	var (
		page     int64
		pagesize int64 = 50
		offset   int64
		pager    string
		m        models.TagInfo
		list     []*models.TagInfo
	)

	//拿到Get参数
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)

	if page < 1 {
		page = 1
	}
	//偏移量
	offset = (page - 1) * pagesize
	query := m.Query()
	counts, _ := query.Count()
	if counts > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}

	pager = this.PagesList(pagesize, page, counts, false, this.admindir+"tag/index")

	// this.Ctx.WriteString(pager)
	this.Data["list"] = list
	this.Data["pager"] = pager
	this.TplName = "admin/tag_index.tpl"
}

//添加标签
func (this *TagHandle) Add() {
	if ok := this.isPost(); ok == true {
		var m models.TagInfo
		m.TagName = this.GetString("tag_name")
		typestr := this.GetString("types")
		types, _ := strconv.ParseInt(typestr, 10, 64)
		m.Types = types
		m.Inset()
		this.Redirect(this.admindir+"tag/index", 302)
	}
	this.TplName = "admin/tag_add.tpl"
}

//编辑标签
func (this *TagHandle) Edit() {

}

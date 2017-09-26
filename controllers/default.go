package controllers

import (
	"quickstart/models"
	"strconv"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	//需求是所有栏目列表，第栏目下有一张缩略图
	var (
		page     int64
		pagesize int64 = 36
		offset   int64
		pager    string
		m        models.PhotoGallery
		list     []*models.PhotoGallery
	)

	//page
	pageStr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pageStr, 10, 64)

	if page < 1 {
		page = 1
	}

	var _list = make([][]*models.PhotoGallery, 0)
	num := 0

	//计数偏移量
	offset = (page - 1) * pagesize
	query := m.Query()
	counts, _ := query.Count()
	if counts > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PagesList(pagesize, page, counts, false, this.admindir)

	var tmpM = make([]*models.PhotoGallery, 0)
	for _, v := range list {

		num += 1
		tmpM = append(tmpM, v)
		if num == 6 {
			_list = append(_list, tmpM)
			num = 0
			tmpM = nil
		}
	}
	this.Data["pager"] = pager
	this.Data["list"] = _list
	this.TplName = "index.tpl"
}

func (this *MainController) GetInfo() {
	idstr := this.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idstr, 10, 64)

	var (
		m    models.PhotoGalleryImgs
		info []*models.PhotoGalleryImgs
	)

	info = m.GetListInfo(id)

	this.Data["info"] = info
	this.TplName = "p_info.tpl"
}

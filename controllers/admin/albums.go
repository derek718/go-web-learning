package admin

import (
	"quickstart/models"
	"strconv"
)

type AlbumsHandle struct {
	BaseController
}

//列表页面
func (this *AlbumsHandle) List() {
	var (
		page     int64
		pagesize int64 = 10
		offset   int64
		pager    string
		list     []*models.PhotoGallery
		m        models.PhotoGallery
	)

	//拿到Get参数
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)

	if page < 1 {
		page = 1
	}

	offset = (page - 1) * pagesize
	query := m.Query()

	//总共多少条数据
	counts, _ := query.Count()
	if counts > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}

	pager = this.PagesList(pagesize, page, counts, false, this.admindir+"p/i")
	this.Data["pager"] = pager
	this.Data["list"] = list

	this.TplName = "admin/list.tpl"
}

//列表详情
func (this *AlbumsHandle) Info() {
	var (
		id    int64
		mlist models.PhotoGallery
		ml    models.PhotoGalleryImgs
		m     models.PhotoGallery
	)

	//拿到Get参数
	idStr := this.Ctx.Input.Param(":id")
	id, _ = strconv.ParseInt(idStr, 10, 64)

	//去查询图集表
	qm := m.Query()
	err := qm.Filter("id", id).One(&mlist)
	if err != nil {
		this.TplName = "admin/p_info.tpl"
		return
	}

	list := ml.GetListInfo(id)

	//图集下的所有图片
	this.Data["list"] = list
	this.Data["title"] = mlist.Title
	this.TplName = "admin/p_info.tpl"
}

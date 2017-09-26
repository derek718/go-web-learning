package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type PhotoGalleryImgs struct {
	Id    int64
	Names string `orm:"size(120)"`
	Hashs string `orm:"size(64)"`
	//	PhotoGallery *PhotoGallery `orm:"rel(many)"`
}

func (this *PhotoGalleryImgs) TableName() string {
	return "photo_gallery_imgs"
}

func (this *PhotoGalleryImgs) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *PhotoGalleryImgs) GetListInfo(id int64) []*PhotoGalleryImgs {
	var msp []*PhotoGalleryImgs
	m := orm.NewOrm()
	strID := strconv.FormatInt(id, 10)
	sql := fmt.Sprintf("select id,names from photo_gallery_imgs where photo_gallery_id=" + strID)
	_, _ = m.Raw(sql).QueryRows(&msp)

	return msp
}

package models

import (
	"github.com/astaxie/beego/orm"
)

type PhotoGallery struct {
	Id     int64
	Title  string `orm:"size(120)"`
	Desc   string `orm:"size(300)"`
	Thumbs string `orm:"size(100)"`
	//	PhotoGalleryImgs *PhotoGalleryImgs `orm:"reverse(many)"`
}

func (this *PhotoGallery) TableName() string {
	return "photo_gallery"
}

func (this *PhotoGallery) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

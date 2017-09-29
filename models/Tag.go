package models

import "time"
import "github.com/astaxie/beego/orm"

type TagInfo struct {
	Id      int64
	TagName string `orm:"size(30)"`
	//标签类型 1:服饰类型，2:美女特征，3：模特人物，4:职业类型，5:动物类型，6:套图，7:搞笑，8:素材
	Types   int64
	Addtime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *TagInfo) TableName() string {
	return "tag_info"
}

func (this *TagInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *TagInfo) Inset() error {
	if _, ok := orm.NewOrm().Insert(this); ok != nil {
		return ok
	}
	return nil
}

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Clumn struct {
	Id        int64
	ClumnName string `orm:"size(32)"`
	Parentid  int64
	Addtime   time.Time `orm:"auto_now_add;type(datetime)"`
}

func (this *Clumn) TableName() string {
	return "clumn"
}

func (this *Clumn) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Clumn) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type AdminInfo struct {
	Id            int64
	Username      string    `orm:"size(20)"`
	Password      string    `orm:"size(32)"`
	Lastlogintime time.Time `orm:"auto_now_add;type(datetime)"`
	Logintimes    int64
	Lastloginip   string
	Addtime       time.Time `orm:"auto_now_add;type(datetime)"`
}

//@return string `表名名称`
func (this *AdminInfo) TableName() string {
	return "admin_info"
}

//@desc `插入一条数据`
func (this *AdminInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}

	return nil
}

//根据传入的字段进行检索一条数据
func (this *AdminInfo) Read(filed ...string) error {
	if err := orm.NewOrm().Read(this, filed...); err != nil {
		return err
	}

	return nil
}

func (this *AdminInfo) Update(filed ...string) error {
	if _, err := orm.NewOrm().Update(this, filed...); err != nil {
		return err
	}

	return nil
}

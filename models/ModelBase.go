package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")

	//注册数据表
	orm.RegisterModel(new(AdminInfo), new(Clumn))
	//mysql
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+")/"+mysqldb+"?charset=utf8&loc=Asia%2FShanghai")
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)

	return strings.ToLower(fmt.Sprintf("%x", m.Sum(nil)))
}

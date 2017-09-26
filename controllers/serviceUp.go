package controllers

import (
	//	"fmt"
	//	"io/ioutil"
	//	"os"
	//	"strings"
	//"fmt"
	//"mime/multipart"

	"github.com/astaxie/beego"
)

type ServiceUpController struct {
	beego.Controller
}

func (c *ServiceUpController) ToPost() {
	f, h, _ := c.GetFile("uploadfile")

	f.Close()
	c.SaveToFile("uploadfile", "./upload/test.png")
	c.Ctx.WriteString(h.Filename)

	c.TplName = "serviceup.tpl"
}

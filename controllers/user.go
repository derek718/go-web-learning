package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("hello")
}

func (c *UserController) Login() {
	fmt.Sprintln("dss")
	c.Ctx.WriteString(c.Ctx.Input.Param("key"))
}

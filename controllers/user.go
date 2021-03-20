package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Create() {
	this.Layout = "layouts/app.html"
	this.Data["Title"] = "注册"
	this.TplName = "users/create.html"
}

package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Layout = "layouts/app.html"
	this.Data["Title"] = "首页"
	this.TplName = "static_pages/home.html"
}

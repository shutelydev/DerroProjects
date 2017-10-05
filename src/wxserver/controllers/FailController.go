package controllers

import (
	"github.com/astaxie/beego"
)

type FailController struct {
	beego.Controller
}

func (this *FailController) Get() {
	this.TplName = "failpage.tpl"
}

package controllers

import (
	"github.com/astaxie/beego"
)

type SuccessController struct {
	beego.Controller
}

func (this *SuccessController) Get() {
	this.TplName = "successpage.tpl"
}

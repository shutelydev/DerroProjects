package routers

import (
	"wxserver/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//beego.Router("/", &controllers.MainController{})
	beego.Router("/weixin", &controllers.WeiXinController{}, "*:WeixinHandler")

}

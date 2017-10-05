package routers

import (
	"wxserver/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//beego.Router("/", &controllers.MainController{})
	beego.Router("/weixin", &controllers.WeiXinController{}, "*:WeixinHandler")

	//booking post
	beego.Router("/api/booking/:roomId", &controllers.BookingController{})

	//booking get
	beego.Router("/booking/?:date", &controllers.BookingController{})

	//booking success page
	beego.Router("/status/success", &controllers.SuccessController{})

	//booking failed page
	beego.Router("/status/fail", &controllers.FailController{})

}

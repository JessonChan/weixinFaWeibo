package main

import (
	"./controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/weixin", &controllers.WeixinController{})
	beego.Router("/weibo", &controllers.WeiboController{})
	beego.Run()
}

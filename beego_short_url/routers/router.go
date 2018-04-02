package routers

import (
	"beego_short_url/controllers"
	"github.com/astaxie/beego"
)

func init() {
	err := controllers.InitDb()
	if err != nil{
		beego.Error("init db is error:",err)
	}
    beego.Router("/trans/long2short", &controllers.ShortUrlController{},"post:Long2Short")
    beego.Router("/trans/short2long", &controllers.ShortUrlController{},"post:Short2Long")
    beego.Router("/shorturl", &controllers.ShortUrlController{},"get:ShortUrlList")
    beego.Router("/jump", &controllers.ShortUrlController{},"get:Jump")

}

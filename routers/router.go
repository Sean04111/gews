package routers

import (
	"gews/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//for the hot news main window
	beego.Router("/hotmain", &controllers.HotController{}, "get:GetHot")
	//for the login site
	beego.AutoRouter(&controllers.LoginController{})
	beego.Router("/", &controllers.LoginController{}, "get:Login")
	//for the register site
	beego.AutoRouter(&controllers.RegisterController{})
	beego.AutoRouter(&controllers.ChatController{})
}

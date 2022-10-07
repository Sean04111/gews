package routers

import (
	"gews/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.AutoRouter(&controllers.ArticleController{})
}

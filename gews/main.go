package main

import (
	"gews/middleware"
	_ "gews/routers"
	"github.com/beego/beego/v2/adapter/cache"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	Ca, _ = cache.NewCache("memory", `{"interval":60}`)
)

func main() {
	middleware.InsertNews()
	beego.Run()
}

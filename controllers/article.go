package controllers

import (
	"gews/DB"
	"gews/middleware"
	"gews/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type HotController struct {
	web.Controller
}

//attention:
//we could just get the hot news and their links
//and just place them in the main site
//but to use the database
//I got to handle it like this
//To get the news form the database although hot news could change every single second!
//It is dumb to insert data into database and get them out few second later
func (this *HotController) GetHot() {
	middleware.InsertNews()
	this.TplName = "HotMain.tpl"
	var Hoty [60]models.News
	for i := 1; i < 50; i++ {
		Hoty[i], _ = DB.GetFromId(i)
	}
	Hot := Hots{A1: Hoty[1], A2: Hoty[2], A3: Hoty[3], A4: Hoty[4], A5: Hoty[5], A6: Hoty[6], A7: Hoty[7], A8: Hoty[8], A9: Hoty[9], A10: Hoty[10], A11: Hoty[11], A12: Hoty[12], A13: Hoty[13], A14: Hoty[14], A15: Hoty[15], A16: Hoty[16], A17: Hoty[17], A18: Hoty[18], A19: Hoty[19], A20: Hoty[20]}
	this.Data["HotNews"] = Hot
	o := orm.NewOrm()
	//dumbest way
	for i := 0; i < 60; i++ {
		o.Delete(&Hoty[i])
	}
}

type Hots struct {
	A1  models.News
	A2  models.News
	A3  models.News
	A4  models.News
	A5  models.News
	A6  models.News
	A7  models.News
	A8  models.News
	A9  models.News
	A10 models.News
	A11 models.News
	A12 models.News
	A13 models.News
	A14 models.News
	A15 models.News
	A16 models.News
	A17 models.News
	A18 models.News
	A19 models.News
	A20 models.News
}

package controllers

import (
	"bufio"
	"encoding/json"
	"gews/DB"
	"gews/models"
	"github.com/beego/beego/v2/server/web"
	"os"
	"strconv"
)

type ArticleController struct {
	web.Controller
}
type Respon struct {
	Status  int         `json:"status"`
	Artcles [30]Newnews `json:"artcles"`
}
type Newnews struct {
	PreNews        models.News
	ArticalLink    string `json:"articalLink"`
	ArticalImgLink string `json:"articalImgLink"`
}

func (this *ArticleController) GetArticle() {
	var res Respon
	res.Status = 0
	for i := 1; DB.CheckId(i) == true; i++ {
		name1 := "static/imgsite/" + strconv.Itoa(i) + ".txt"
		name2 := "static/test/" + strconv.Itoa(i) + ".txt"
		file, _ := os.Open(name1)
		reader := bufio.NewReader(file)
		imglink, _, _ := reader.ReadLine()
		prenews, _ := DB.GetFromId(i)
		res.Artcles[i].PreNews = prenews
		res.Artcles[i].ArticalLink = name2
		res.Artcles[i].ArticalImgLink = string(imglink)
	}
	resjson, _ := json.Marshal(res)
	this.Ctx.Output.Body(resjson)
}

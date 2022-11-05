package controllers

import (
	"fmt"
	"gews/DB"
	"github.com/beego/beego/v2/server/web"
	"time"
)

type LoginController struct {
	web.Controller
}

func (this *LoginController) Login() {
	web.SetStaticPath("window.css", "css")
	this.TplName = "Login.html"
}

type Input struct {
	Id       int    `form:"id"`
	Password string `form:"password"`
	Name     string `form:"name"`
}
type SayHello struct {
	Sayhi string
	Today string
}

func (this *LoginController) Check() {
	input := new(Input)
	err := this.ParseForm(input)
	if err != nil {
		fmt.Println("Login:Parse from form error: ", err)
	}
	right, name := DB.CheckUsers(input.Id, input.Password)
	input.Name = name
	if right == true {
		var say = new(SayHello)
		say.Today = time.Now().Weekday().String()
		if time.Now().Hour() < 12 {
			say.Sayhi = "上午好"
		}
		if 12 < time.Now().Hour() && time.Now().Hour() < 19 {
			say.Sayhi = "下午好"
		}
		if 19 < time.Now().Hour() {
			say.Sayhi = "晚上好"
		}
		this.TplName = "Welcome.tpl"
		this.Data["User"] = input
		this.Data["Hello"] = say
	} else {
		this.TplName = "login_wrong.html"
	}
}

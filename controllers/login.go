package controllers

import (
	"fmt"
	"gews/DB"
	"github.com/beego/beego/v2/server/web"
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

func (this *LoginController) Check() {
	input := new(Input)
	err := this.ParseForm(input)
	if err != nil {
		fmt.Println("Login:Parse from form error: ", err)
	}
	right, name := DB.CheckUsers(input.Id, input.Password)
	input.Name = name
	if right == true {
		this.TplName = "Welcome.tpl"
		this.Data["User"] = input
	} else {
		this.TplName = "login_wrong.html"
	}
}

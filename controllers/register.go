package controllers

import (
	"fmt"
	"gews/DB"
	"github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	web.Controller
}

func (this *RegisterController) Register() {
	this.TplName = "Register.html"
}

func (this *RegisterController) Do() {
	input := new(Input)
	err := this.ParseForm(input)
	if err != nil {
		fmt.Println("Register:Parse form error :", err)
	}
	err = DB.InsertUsers(input.Id, input.Password, input.Name)
	if err != nil {
		fmt.Println("Register:Insert Users error:", err)
	} else {
		this.TplName = "RegisterOK.tpl"
		this.Data["User"] = input
	}
}

package controllers

import (
	"fmt"
	"gews/chat"
	"github.com/beego/beego/v2/server/web"
	"google.golang.org/grpc"
	"strconv"
)

type ChatController struct {
	web.Controller
}
type Words struct {
	Content string `form:"content"`
}

func (this *ChatController) Say() {
	name := this.Ctx.Input.Params()["Name"]
	id, _ := strconv.Atoi(this.Ctx.Input.Params()["Id"])
	conn1, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	you := chat.NewUser(id, name, conn1)
	input := new(Words)
	err := this.ParseForm(input)
	if err != nil {
		fmt.Println(err)
	}
	you.SaytoAll(input.Content)
}

type Messa struct {
	M1  chat.Messages
	M2  chat.Messages
	M3  chat.Messages
	M4  chat.Messages
	M5  chat.Messages
	M6  chat.Messages
	M7  chat.Messages
	M8  chat.Messages
	M9  chat.Messages
	M10 chat.Messages
	M11 chat.Messages
	M12 chat.Messages
}

func (this *ChatController) ChatMain() {
	this.TplName = "Chat.tpl"
	name := this.Ctx.Input.Params()["Name"]
	id, _ := strconv.Atoi(this.Ctx.Input.Params()["Id"])
	conn1, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	you := chat.NewUser(id, name, conn1)
	get := you.GetAllMass(12)
	all := Messa{M1: get[0], M2: get[1], M3: get[2], M4: get[3], M5: get[4], M6: get[5], M7: get[6], M8: get[7], M9: get[8], M10: get[9], M11: get[10], M12: get[11]}
	this.Data["M"] = all
	this.Data["User"] = you.Self
}

package controllers

import (
	"fmt"
	"gews/chat"
	"gews/models"
	"github.com/beego/beego/v2/server/web"
	"google.golang.org/grpc"
	"strconv"
)

type Cos struct {
	Id   int
	Name string
}
type ChatController struct {
	web.Controller
}

type Words struct {
	Content string `form:"text"`
}

func (this *ChatController) Say() {
	name := this.Ctx.Input.Params()["0"]
	id, _ := strconv.Atoi(this.Ctx.Input.Params()["1"])
	conn1, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	you := chat.NewUser(id, name, conn1)
	input := new(Words)
	err := this.ParseForm(input)
	if err != nil {
		fmt.Println(err)
	}
	you.SaytoAll(input.Content)
	//I was thinking about to send a feedback to the client by using session or cookie
	this.ChatMain()

}

type Messa struct {
	M1  models.Messages
	M2  models.Messages
	M3  models.Messages
	M4  models.Messages
	M5  models.Messages
	M6  models.Messages
	M7  models.Messages
	M8  models.Messages
	M9  models.Messages
	M10 models.Messages
	M11 models.Messages
	M12 models.Messages
	M13 models.Messages
	M14 models.Messages
	M15 models.Messages
	M16 models.Messages
	M17 models.Messages
	M18 models.Messages
	M19 models.Messages
	M20 models.Messages
}

//Right here
//the length of get depends on the num of the messages in the database !
//so this is a bug anyway!!!
func (this *ChatController) ChatMain() {
	this.TplName = "Chat.tpl"
	name := this.Ctx.Input.Params()["0"]
	id, _ := strconv.Atoi(this.Ctx.Input.Params()["1"])
	conn1, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	you := chat.NewUser(id, name, conn1)
	get := you.GetAllMass()
	all := Messa{M1: get[0], M2: get[1], M3: get[2], M4: get[3], M5: get[4], M6: get[5], M7: get[6], M8: get[7], M9: get[8], M10: get[9], M11: get[10], M12: get[11], M13: get[13], M14: get[14], M15: get[15], M16: get[16], M17: get[17], M18: get[18], M19: get[19], M20: get[20]}
	this.Data["M"] = all
	this.Data["User"] = you.Self
}

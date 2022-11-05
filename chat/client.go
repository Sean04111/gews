package chat

import (
	"context"
	"fmt"
	pb "gews/chat/proto"
	"gews/models"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

type User struct {
	ClientChat pb.ChatClient
	Self       *pb.User
}

//you can set you database by these vars
var (
	DB_database string = "sql_news"
	DB_username string = "sql_news"
	DB_password string = "666"
	DB_host     string = "121.36.131.50"
	DB_port     string = "3306"
)

func (this *User) SaytoAll(words string) {
	Now := time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day()) + " " + strconv.Itoa(time.Now().Hour()) + ":" + strconv.Itoa(time.Now().Minute()) + ":" + strconv.Itoa(time.Now().Second())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	MassageNum, _ := this.ClientChat.GetMessNum(ctx, &pb.UserId{Id: this.Self.Id})
	_, err2 := this.ClientChat.SendAll(ctx, &pb.Message{Id: MassageNum.Messnum,
		Content:     words,
		Speakername: this.Self.Name,
		Time:        Now,
	})
	if err2 != nil {
		fmt.Println("[chat] SendAll error : ", err2)
	}
}
func (this *User) GetMassFromId(id int) *pb.Message {
	var Re *pb.Message
	return Re
}
func (this *User) GetMassFromName(speakername string) *pb.Message {
	find := new(pb.Message)

	return find
}

//The messages here is backwards
//The latest messages will be placed at the first place !!
func (this *User) GetAllMass() []models.Messages {
	var Re []models.Messages
	for i := 50; i >= 1; i-- {
		var kv models.Messages
		o := orm.NewOrm()
		qsu := o.QueryTable("messages")
		_, err := qsu.Filter("message_id", i).All(&kv)
		if err != nil {
			continue
		}
		if kv.MessageId != 0 {
			Re = append(Re, kv)
		}
	}
	return Re
}

func NewUser(id int, name string, grpcclientconn *grpc.ClientConn) *User {
	Sean := new(User)
	self := new(pb.User)
	self.Id = int64(id)
	self.Name = name
	Sean.Self = self
	cliC := pb.NewChatClient(grpcclientconn)
	Sean.ClientChat = cliC
	return Sean
}

var SqlName string = "sql_news"
var Password string = "666"
var Ip string = "121.36.131.50:3306"
var DBname string = "sql_news"

func init() {

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println("[orm] Register Driver error : ", err)
	}
	err = orm.RegisterDataBase("default", "mysql", SqlName+":"+Password+"@tcp("+Ip+")/"+DBname+"?charset=utf8")
	if err != nil {
		fmt.Println("[orm] Register Data Base error : ", err)
	}
	/*
		err = orm.RunSyncdb("default", false, false)
		if err != nil {
			fmt.Println("[orm] Create Table error : ", err)
		}
	*/
}

/*
func main() {
	conn1, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	fmt.Println(conn1.Target())
	if err != nil {
		fmt.Println("Dial error : ", err)
	}
	if conn1 == nil {
		fmt.Println("wrong")
		return
	}
	defer func(conn1 *grpc.ClientConn) {
		err := conn1.Close()
		if err != nil {
			fmt.Println("Failed to close the conn")
		}
	}(conn1)
	s := NewUser(12, "Atoi", conn1)
	fmt.Println(s.GetAllMass(5))
}
*/

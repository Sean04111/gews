package DB

import (
	"fmt"
	"gews/models"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//
//
//ini the date base
func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println("Driver register error :", err)
	}
	err = orm.RegisterDataBase("default", "mysql", "sql_news:666@tcp(121.36.131.50:3306)/sql_news?charset=utf8")
	if err != nil {
		fmt.Println("Register DataBase error : ", err)
	}
	//orm.RunSyncdb("default", true, true)
}

//
//
//
//functions for news
func Insert(news models.News) error {
	o := orm.NewOrm()
	_, err := o.Insert(&news)
	return err
}
func GetFromId(articlie_id int) (models.News, error) {
	var resp models.News
	o := orm.NewOrm()
	qsu := o.QueryTable("news")
	_, err := qsu.Filter("article_id", articlie_id).All(&resp)
	if err != nil {
		fmt.Println("Get from id error :", err)
	}
	return resp, err
}
func CheckId(id int) bool {
	o := orm.NewOrm()
	qsu := o.QueryTable("news")
	exist := qsu.Filter("article_id", id).Exist()
	return exist
}
func GetFromTitle(title string) models.News {
	var resp models.News
	o := orm.NewOrm()
	qsu := o.QueryTable("news")
	_, err := qsu.Filter("title", title).All(&resp)
	if err != nil {
		fmt.Println("Get from title error:", err)
	}
	return resp
}

//
//
//
//functions for users
func InsertUsers(id int, password string) error {
	bkey := []byte(password)
	hashedpassword, err := bcrypt.GenerateFromPassword(bkey, 10)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		storekey := string(hashedpassword)
		o := orm.NewOrm()
		_, err := o.Insert(&models.Users{
			Id:         id,
			Hashed_key: storekey,
		})
		return err
	}
}
func CheckUsers(id int, password string) bool {
	bkey := []byte(password)

	o := orm.NewOrm()
	qsu := o.QueryTable("Users")

	var find models.Users
	var findhashed []byte
	exist := qsu.Filter("id", id).Exist()
	if exist == false {
		fmt.Println("Account input does not exist !")
		return false
	} else {
		qsu.Filter("id", id).All(&find, "hashed_key")
		findhashed = []byte(find.Hashed_key)
	}
	err := bcrypt.CompareHashAndPassword(findhashed, bkey)
	if err == nil {
		return true
	} else {
		return false
	}
}

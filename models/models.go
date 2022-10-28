package models

import "github.com/beego/beego/v2/client/orm"

type Users struct {
	Id         int    `orm:"auto"json:"id"orm:"column(id)"`
	Hashed_key string `json:"hashed_Key"orm:"column(hashed_key)"`
	Name       string `orm:"column(name)"`
}

type News struct {
	Article_id int    `orm:"auto"json:"articleid"orm:"column(article_id)"`
	Article    string `json:"article"orm:"column(article)"` //for the links of each
	Title      string `json:"title"orm:"column(title)"`     //for the links of each
}

func init() {
	orm.RegisterModel(new(Users), new(News))
}

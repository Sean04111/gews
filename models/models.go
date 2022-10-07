package models

import "github.com/beego/beego/v2/client/orm"

type Users struct {
	Id         int    `orm:"auto"json:"id"orm:"column(id)"`
	Hashed_key string `json:"hashed_Key"orm:"column(hashed_key)"`
}

type News struct {
	Article_id int    `orm:"auto"json:"articleid"orm:"column(article_id)"`
	Article    string `json:"article"orm:"column(article)"`
	Likes      int    `json:"likes"orm:"column(likes)"`
	Checknum   int    `json:"checknum"orm:"column(checknum)"`
	Title      string `json:"title"orm:"column(title)"`
}

func init() {
	orm.RegisterModel(new(Users), new(News))
}

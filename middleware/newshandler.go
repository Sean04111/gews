package middleware

import (
	"fmt"
	"gews/DB"
	"gews/Spider"
	"gews/models"
)

func InsertNews() {
	url := "http://resou.today/art/20.html"
	titles, links := Spider.VisitMain(url)
	for i := 0; len(titles[i]) != 0; i++ {
		var news models.News
		news.Title = titles[i]
		news.Article_id = i + 1
		news.Article = links[i]
		if err := DB.Insert(news); err != nil {
			fmt.Println("id ", i, "insert error:", err)
		}
	}
}

package middleware

import (
	"fmt"
	"gews/DB"
	"gews/Spider"
	"gews/models"
	"strconv"
)

func InsertNews() {
	url := "https://www.sina.com.cn/"
	titles, links := Spider.VisitSinaMain(url)
	for i := 1; len(links[i]) != 0; i++ {
		var news models.News
		text, img := Spider.VisitSingleNew(links[i])
		Spider.WriteArticle(text, i)
		Spider.WriteImg(img, i)
		news.Article_id = i
		news.Article = "static/text/" + strconv.Itoa(i) + ".txt"
		news.Likes = 0
		news.Checknum = 0
		news.Title = titles[i]
		if err := DB.Insert(news); err != nil {
			fmt.Println("id ", i, "insert error:", err)
		}
	}
}

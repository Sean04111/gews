package Spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//return the news title and their links
func VisitSinaMain(url string) ([50]string, [50]string) {
	var alltext [50]string
	var allhref [50]string
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(reqest)
	docu, _ := goquery.NewDocumentFromReader(response.Body)
	docu.Find("#syncad_0 > ul > li > a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		text := s.Text()
		alltext[i] = text
		allhref[i] = href
	})
	return alltext, allhref
}

//return the text and the site of the imgs
func VisitSingleNew(url string) ([1000]string, [50]string) {
	var text [1000]string
	var img [50]string
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println("http client error : ", err)
	}
	documents, _ := goquery.NewDocumentFromReader(response.Body)
	documents.Find("#article > p").Each(func(i int, selection *goquery.Selection) {
		te := selection.Text()
		text[i] = te
	})
	documents.Find("#article > div > img").Each(func(i int, selection *goquery.Selection) {
		imgsrc_, _ := selection.Attr("src")
		imgsrc := "https" + imgsrc_
		img[i] = imgsrc
	})
	return text, img
}

//
//plan B
func Fetch(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "ttcid=e2124a2f89f64c7793da7623cbbab40734; _ga=GA1.2.660092466.1651074958; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091312926552081924%2522%252C%2522user_unique_id%2522%253A%25227091312926552081924%2522%252C%2522timestamp%2522%253A1651074958228%257D; sid_guard=31979794b33a6fc2b9a20c6363a5f9ed%7C1651889153%7C31536000%7CSun%2C+07-May-2023+02%3A05%3A53+GMT; uid_tt=e730ac0086d71e1aa46412015f088c51; uid_tt_ss=e730ac0086d71e1aa46412015f088c51; sid_tt=31979794b33a6fc2b9a20c6363a5f9ed; sessionid=31979794b33a6fc2b9a20c6363a5f9ed; sessionid_ss=31979794b33a6fc2b9a20c6363a5f9ed; sid_ucp_v1=1.0.0-KGQ2N2U0OGIzNWQ1OTU2MjkxMGI3ZTkxNzk2YTY4MWU1Y2I5M2RjZDgKFwjHp4Ce0YytAxCBqNeTBhiwFDgCQPEHGgJsZiIgMzE5Nzk3OTRiMzNhNmZjMmI5YTIwYzYzNjNhNWY5ZWQ; ssid_ucp_v1=1.0.0-KGQ2N2U0OGIzNWQ1OTU2MjkxMGI3ZTkxNzk2YTY4MWU1Y2I5M2RjZDgKFwjHp4Ce0YytAxCBqNeTBhiwFDgCQPEHGgJsZiIgMzE5Nzk3OTRiMzNhNmZjMmI5YTIwYzYzNjNhNWY5ZWQ; _tea_utm_cache_2608={%22utm_source%22:%22xitongxiaoxi%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22chuxia%22}; passport_csrf_token=93e09b2b47c5d95d646e80cad84dc95e; passport_csrf_token_default=93e09b2b47c5d95d646e80cad84dc95e; MONITOR_WEB_ID=d1aa8d4f-0754-4b62-8d40-d192603574e2; _gid=GA1.2.189140324.1664174688; _jj_ext=1; MONITOR_DEVICE_ID=1555a3ed-47fa-4246-a197-e85b2713a08e; tt_scid=rzk1iS4FRt3v5vKDlJmhKAuSDavrUiyFkhhDzXTx3hc3W9i6MN9KD16WAMH3Z8Dq7b8d")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http Get err :", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("http status code :", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error :", err)
		return ""
	}

	return string(body)
}

//
//plan B
func Parse(html string) {
	html = strings.Replace(html, "\n", "", -1)
	regexp_sidebar := regexp.MustCompile(`<ul class="nav bs-docs-sidenav">(.*?)</ul>`)
	sidebar := regexp_sidebar.FindString(html)
	regexp_link := regexp.MustCompile(`href="(.*?)"`)
	links := regexp_link.FindAllString(sidebar, -1)
	fmt.Println(links)
}
func WriteArticle(text [1000]string, id int) error {
	var name string
	name = "static/text/" + strconv.Itoa(id) + ".txt"
	file, err := os.Create(name)
	for i := 1; len(text[i]) != 0; i++ {
		file.WriteString(text[i])
		file.WriteString("\n")
	}
	file.Close()
	return err
}
func WriteImg(img [50]string, id int) error {
	var name string
	name = "static/imgsite/" + strconv.Itoa(id) + ".txt"
	file, err := os.Create(name)
	for i := 1; len(img[i]) != 0; i++ {
		file.WriteString(img[i])
		file.WriteString("\n")
	}
	file.Close()
	return err
}

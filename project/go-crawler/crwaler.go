package go_crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

/***
爬虫 https://strconv.com/page/2/
获取top250
*/
func fetch(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}

/***
解析body
*/
func parseBody(body string) {
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := rp.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		fmt.Println(idRe.FindStringSubmatch(item[1])[1],
			titleRe.FindStringSubmatch(item[1])[1])
	}
}

func parseBodyByGoQuery(body string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	doc.Find("ol.grid_view li").Find(".hd").Each(func(index int, ele *goquery.Selection) {
		movieUrl, _ := ele.Find("a").Attr("href")
		fmt.Println(strings.Split(movieUrl, "/")[4], ele.Find(".title").Eq(0).Text())
	})
}

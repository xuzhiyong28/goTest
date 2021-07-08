package goquery

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

/**
网络爬虫 - 爬取豆瓣
**/
func Demo1() {
	var movies []DoubanMovie
	pages := GetPages(BaseUrl)
	for _, page := range pages {
		body := GetHttpsBody(strings.Join([]string{BaseUrl, page.Url}, ""))
		doc, err := goquery.NewDocumentFromReader(body)
		body.Close()
		if err != nil {
			fmt.Println(err)
		}
		movies = append(movies, ParseMovies(doc)...)
	}
	fmt.Println(movies)
}

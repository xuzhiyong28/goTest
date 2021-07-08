package goquery

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const BaseUrl = "https://movie.douban.com/top250"

type DoubanMovie struct {
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

type Page struct {
	Page int
	Url  string
}


func GetHttpsBody(url string) io.ReadCloser {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	reqest, _ := http.NewRequest("GET", url, nil)
	reqest.Header.Add("User-Agent", "'Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	reqest.Header.Add("Referer","https://movie.douban.com/")
	reqest.Header.Add("Connection","keep-alive")
	res, _ := client.Do(reqest)
	return res.Body
}


// 获取分页
func GetPages(url string) []Page {
	body := GetHttpsBody(url)
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println(err)
	}
	return ParsePages(doc)
}

// 分析分页
func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

// 分析电影数据
func ParseMovies(doc *goquery.Document) (movies []DoubanMovie) {
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()
		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")
		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")
		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]
		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])
		star := s.Find(".bd .star .rating_num").Text()
		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")
		quote := s.Find(".quote .inq").Text()
		movie := DoubanMovie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}
		movies = append(movies, movie)
	})

	return movies
}



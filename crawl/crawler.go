package crawl

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type EasyCrawler struct {
	Depth       int
	ReadUrlList []string

	depth_counter int
}

func (crawler *EasyCrawler) Crawl(depth int, u string) {

	depthCountChanel := make(chan bool)
	crawler.crawl(u, depthCountChanel)
}

func (crawler *EasyCrawler) crawl(u string, d chan bool) {
	start := time.Now()

	c := make(chan Content)

	var content Content = <-c
	crawler.depth_counter++
	fmt.Println(crawler.depth_counter, "週目！ 長さは", len(content.Urls))
	for _, url := range content.Urls {
		go crawler.getContentFromUrl(url, c)
	}
	crawler.saveHost(u)

	end := time.Now()
	fmt.Printf("実行時間%f秒\n", (end.Sub(start)).Seconds())
	crawler.crawl(u, d)
}

func (crawler *EasyCrawler) getContentFromUrl(u string, c chan Content) {
	var urls = []string{}

	baseUrl, urlParseError := url.Parse(u)
	if urlParseError != nil {
		fmt.Println("url parse error:", urlParseError)
		c <- Content{Url: "", Urls: []string{}, Body: ""}
	}

	resp, httpGetError := http.Get(baseUrl.String())
	if httpGetError != nil {
		fmt.Println("http error:", httpGetError)
		c <- Content{Url: "", Urls: []string{}, Body: ""}
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	html, htmlGetError := doc.Html()
	if htmlGetError != nil {
		fmt.Println("html extract error:", htmlGetError)
		c <- Content{Url: "", Urls: []string{}, Body: ""}
	}

	// parse a tag and href
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			reqUrl, err := url.Parse(href)
			for _, v := range urls {
				if reqUrl.String() == v {
					return
				}
			}
			if err == nil {
				urls = append(urls, reqUrl.String())
			}
		}
	})
	time.Sleep(1 * time.Second)
	c <- Content{Url: baseUrl.String(), Urls: urls, Body: html}
}

func (crawler *EasyCrawler) saveHost(u string) {
	for _, v := range crawler.ReadUrlList {
		if u == v {
			return
		}
	}
	crawler.ReadUrlList = append(crawler.ReadUrlList, u)
}

func (crawler *EasyCrawler) SetCallBack() {
	fmt.Println("set CallBack!")
}
package crawl

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type EasyCrawler struct {
	Depth       int
	InitialUrl  string
	ReadUrlList []string
}

func (crawler *EasyCrawler) Crawl(u string) {
	baseUrl, err := url.Parse(u)
	//var content = Content{}
	resp, err := http.Get(baseUrl.String())
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	//urls = make([]string, 0)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			reqUrl, err := baseUrl.Parse(href)
			if err == nil {
				fmt.Println("url一覧", reqUrl, err)
				crawler.ReadUrlList = append(crawler.ReadUrlList, reqUrl.String())
				//urls = append(urls, reqUrl.String())
			}
		}
	})
}

func (crawler *EasyCrawler) SetCallBack() {
	fmt.Println("set CallBack!")
}

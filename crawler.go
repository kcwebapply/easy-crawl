package easyCrawl

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type EasyCrawler struct {
	Depth int

	readUrlList       []string
	callBackInterface CallBackInterface
	logger            Logging
}

func (crawler *EasyCrawler) Crawl(u string) error {
	if crawler.callBackInterface == nil {
		return errors.New("please implements CallBackInterface and set that object by `SetCallBack` method.")
	}

	if &crawler.logger == nil {
		crawler.logger = Logging{logging: false}
	}

	var internalDEPTHCounter = 0
	content := Content{Url: u, Urls: []string{u}, Body: ""}
	crawler.crawl([]Content{content}, &internalDEPTHCounter)
	return nil
}

func (crawler *EasyCrawler) crawl(contentList []Content, internalDEPTHCounter *int) {
	// initialize
	start := time.Now()
	c := make(chan Content)
	*internalDEPTHCounter++
	newURLList := crawler.newURLList(contentList)
	crawler.logger.logDepth(*internalDEPTHCounter, len(newURLList))

	// get contents parallel
	for _, url := range newURLList {
		go crawler.getContentFromUrl(url, c)
	}

	// 並列取得したコンテンツ郡を取得済みホストに保存し、Contentに変換。
	newContentList := []Content{}
	for i := 0; i < len(newURLList); i++ {
		content := <-c
		crawler.logger.logCrawlDone(content.Url)
		crawler.callBackInterface.Callback(content.Url, content.Urls, content.Body)
		newContentList = append(newContentList, content)
		crawler.saveHost(content.Url)
	}
	// 終了処理。実行時間を記述。
	end := time.Now()
	crawler.logger.logTime((end.Sub(start)).Seconds())

	if crawler.Depth > *internalDEPTHCounter {
		crawler.crawl(newContentList, internalDEPTHCounter)
	}
}

func (crawler *EasyCrawler) getContentFromUrl(u string, c chan Content) {
	var urls = []string{}

	baseUrl, urlParseError := url.Parse(u)
	if urlParseError != nil {
		//fmt.Println("url parse error:", urlParseError)
		c <- Content{Url: "", Urls: []string{}, Body: ""}
		return
	}

	resp, httpGetError := http.Get(baseUrl.String())
	if httpGetError != nil {
		//fmt.Println("http error:", httpGetError)
		c <- Content{Url: "", Urls: []string{}, Body: ""}
		return
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	html, htmlGetError := doc.Html()
	if htmlGetError != nil {
		//fmt.Println("html extract error:", htmlGetError)
		c <- Content{Url: "", Urls: []string{}, Body: ""}
		return
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

func (crawler *EasyCrawler) crawlChecked(u string) bool {
	for _, v := range crawler.readUrlList {
		if u == v {
			return true
		}
	}
	return false
}

func (crawler *EasyCrawler) newURLList(contentList []Content) []string {
	urlList := []string{}
	newURLList := []string{}
	for _, content := range contentList {
		for _, url := range content.Urls {
			urlList = append(urlList, url)
		}
	}
	for _, url := range urlList {
		if !crawler.crawlChecked(url) {
			newURLList = append(newURLList, url)
		}
	}
	return newURLList
}

func (crawler *EasyCrawler) saveHost(u string) {
	if !crawler.crawlChecked(u) {
		crawler.readUrlList = append(crawler.readUrlList, u)
	}
}

func (crawler *EasyCrawler) SetCallBack(callBackInterface CallBackInterface) {
	crawler.callBackInterface = callBackInterface
}

func (crawler *EasyCrawler) SetLogging(enabled bool) {
	crawler.logger = Logging{logging: enabled}
}

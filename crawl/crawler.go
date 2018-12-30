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
	content := Content{Url: u, Urls: []string{u}, Body: ""}
	crawler.crawl(content, depthCountChanel)
}

func (crawler *EasyCrawler) crawl(content Content, d chan bool) {
	// 諸々初期実装。深さカウンターを+1した。
	start := time.Now()
	c := make(chan Content)
	crawler.depth_counter++
	fmt.Println(crawler.depth_counter, "週目！Hostは ", content.Url, " 長さは", len(content.Urls))

	// goroutineで並列でコンテンツ取得
	newURLList := crawler.newURLList(content.Urls)
	for _, url := range newURLList {
		fmt.Println("新url:", url)
		go crawler.getContentFromUrl(url, c)
	}

	// 並列取得したコンテンツ郡を取得済みホストに保存し、Contentに変換。
	contentList := []Content{}
	for i := 0; i < len(newURLList); i++ {
		content := <-c
		contentList = append(contentList, content)
		//fmt.Println("取得コンテンツUrl:", content.Url)
		crawler.saveHost(content.Url)
	}
	// 終了処理。実行時間を記述。
	end := time.Now()
	fmt.Printf("実行時間%f秒\n", (end.Sub(start)).Seconds())

	// 取得したcontent一覧毎に、crawlを行う。
	for _, value := range contentList {
		crawler.crawl(value, d)
	}
	//crawler.crawl(u, d)
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

func (crawler *EasyCrawler) crawlChecked(u string) bool {
	for _, v := range crawler.ReadUrlList {
		if u == v {
			return true
		}
	}
	return false
}

func (crawler *EasyCrawler) newURLList(urlList []string) []string {
	newURLList := []string{}
	for _, url := range urlList {
		if !crawler.crawlChecked(url) {
			newURLList = append(newURLList, url)
		}
	}
	return newURLList
}

func (crawler *EasyCrawler) saveHost(u string) {
	if !crawler.crawlChecked(u) {
		//fmt.Println("このurlを保存しました", u)
		crawler.ReadUrlList = append(crawler.ReadUrlList, u)
	}
}

func (crawler *EasyCrawler) SetCallBack() {
	fmt.Println("set CallBack!")
}

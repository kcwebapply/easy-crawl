package easycrawl

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// EasyCrawler is struct for crawling web pages
type EasyCrawler struct {
	Depth int

	readURLList       []string
	callBackInterface CallBackInterface
	logger            Logging
}

// Crawl method
func (crawler *EasyCrawler) Crawl(u string) error {
	if crawler.callBackInterface == nil {
		return errors.New("Please implements CallBackInterface and set that object by `SetCallBack` method")
	}

	if &crawler.logger == nil {
		crawler.logger = Logging{logging: false}
	}

	var internalDEPTHCounter = 0
	content := Content{URL: u, Urls: []string{u}, Body: ""}
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
		go crawler.getContentFromURL(url, c)
	}

	// receive Content and call `Callback` method, and save ContentData in crawled List.
	newContentList := []Content{}
	for i := 0; i < len(newURLList); i++ {
		content := <-c
		if content.URL == "" {
			continue
		}
		crawler.logger.logCrawlDone(content.URL)
		crawler.callBackInterface.Callback(content.URL, content.Urls, content.Body)
		newContentList = append(newContentList, content)
		crawler.saveHost(content.URL)
	}
	// time logging
	end := time.Now()
	crawler.logger.logTime((end.Sub(start)).Seconds())

	// recursive crawling.
	if crawler.Depth > *internalDEPTHCounter {
		crawler.crawl(newContentList, internalDEPTHCounter)
	}
}

func (crawler *EasyCrawler) getContentFromURL(u string, c chan Content) {
	var urls = []string{}
	baseURL, urlParseError := url.Parse(u)
	if urlParseError != nil {
		//fmt.Println("url parse error:", urlParseError)
		c <- Content{URL: "", Urls: []string{}, Body: ""}
		return
	}

	resp, httpGetError := http.Get(baseURL.String())
	if httpGetError != nil {
		//fmt.Println("http error:", httpGetError)
		c <- Content{URL: "", Urls: []string{}, Body: ""}
		return
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	html, htmlGetError := doc.Html()
	if htmlGetError != nil {
		//fmt.Println("html extract error:", htmlGetError)
		c <- Content{URL: "", Urls: []string{}, Body: ""}
		return
	}

	// parse a tag and href
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			reqURL, err := url.Parse(href)
			for _, v := range urls {
				if reqURL.String() == v {
					return
				}
			}
			if err == nil {
				urls = append(urls, reqURL.String())
			}
		}
	})
	c <- Content{URL: baseURL.String(), Urls: urls, Body: html}
}

func (crawler *EasyCrawler) crawlChecked(u string) bool {
	for _, v := range crawler.readURLList {
		if u == v {
			return true
		}
	}
	return false
}

func (crawler *EasyCrawler) newURLList(contentList []Content) []string {
	urlList := []string{}
	urlSetList := []string{}
	newURLList := []string{}
	// get all url contained in all contents
	for _, content := range contentList {
		for _, url := range content.Urls {
			urlList = append(urlList, url)
		}
	}
	// delete duplicate url
	m := make(map[string]struct{})
	for _, url := range urlList {
		// check
		if _, ok := m[url]; !ok {
			m[url] = struct{}{}
			urlSetList = append(urlSetList, url)
		}
	}

	for _, url := range urlSetList {
		if !crawler.crawlChecked(url) {
			newURLList = append(newURLList, url)
		}
	}
	return newURLList
}

func (crawler *EasyCrawler) saveHost(u string) {
	if !crawler.crawlChecked(u) {
		crawler.readURLList = append(crawler.readURLList, u)
	}
}

// SetCallBack method is method for set callback method which will be called when contents is acquired
func (crawler *EasyCrawler) SetCallBack(callBackInterface CallBackInterface) {
	crawler.callBackInterface = callBackInterface
}

// SetLogging method is for setting printing log or not
func (crawler *EasyCrawler) SetLogging(enabled bool) {
	crawler.logger = Logging{logging: enabled}
}

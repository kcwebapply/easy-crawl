package main

import (
	"fmt"

	. "github.com/kcwebapply/easy-crawl/crawl"
)

func main() {
	fmt.Println("test")
	crawler := EasyCrawler{Depth: 2}
	callBackImpl := CallBackImpl{}
	crawler.SetCallBack(callBackImpl)
	crawler.Crawl("http://spring-boot-reference.jp/")
}

type CallBackImpl struct {
}

func (callbackImpl CallBackImpl) Callback(url string, urls []string, body string) {
	fmt.Println("きたーーー", url)
}

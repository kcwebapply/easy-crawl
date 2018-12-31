package main

import (
	"fmt"

	"github.com/kcwebapply/easy-crawl"
)

func main() {
	fmt.Println("test")
	crawler := easyCrawl.EasyCrawler{Depth: 2}
	callBackImpl := CallBackImpl{}
	crawler.SetCallBack(callBackImpl)
	crawler.Crawl("http://spring-boot-reference.jp/")
}

type CallBackImpl struct {
}

func (callbackImpl CallBackImpl) Callback(url string, urls []string, body string) {
	fmt.Println("きたー : ", url)
}

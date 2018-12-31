package main

import (
	"github.com/kcwebapply/easy-crawl"
)

func main() {
	crawler := easyCrawl.EasyCrawler{Depth: 3}
	callBackImpl := CallBackImpl{}
	crawler.SetCallBack(callBackImpl)
	crawler.SetLogging(true)
	crawler.Crawl("http://spring-boot-reference.jp/")
}

type CallBackImpl struct {
}

func (callbackImpl CallBackImpl) Callback(url string, urls []string, body string) {
	//fmt.Println("callBack Method Called! : ", url)
}

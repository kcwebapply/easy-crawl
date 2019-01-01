package main

import (
	easycrawl "github.com/kcwebapply/easy-crawl"
)

func main() {
	crawler := easycrawl.EasyCrawler{Depth: 3}
	callBackImpl := CallBackImpl{}
	crawler.SetCallBack(callBackImpl)
	crawler.SetLogging(true)
	crawler.Crawl("http://spring-boot-reference.jp/")
}

// CallBackImpl is example struct implementing CallBackInterface
type CallBackImpl struct {
}

// Callback method is example of  callback method
func (callbackImpl CallBackImpl) Callback(url string, urls []string, body string) {
	//fmt.Println("callBack Method Called! : ", url)
}

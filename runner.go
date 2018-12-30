package main

import (
	"fmt"

	. "github.com/kcwebapply/easy-crawl/crawl"
)

func main() {
	fmt.Println("test")
	crawler := EasyCrawler{}
	crawler.Crawl(10, "http://spring-boot-reference.jp/")

	fmt.Println("------")

	/*for i := 0; i < 10; i++ {
		url := crawler.ReadUrlList[i]
		fmt.Println("対象urlは", url)
		//crawler.Crawl(url)
	}*/
}

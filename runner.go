package main

import (
	"fmt"

	. "github.com/kcwebapply/easy-crawl/crawl"
)

func main() {
	fmt.Println("test")
	crawler := EasyCrawler{}
	crawler.Crawl("http://spring-boot-reference.jp/")

	fmt.Println("------")
	fmt.Println(crawler.ReadUrlList)
}

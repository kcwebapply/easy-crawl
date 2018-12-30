package main

import (
	"fmt"

	. "github.com/kcwebapply/easy-crawl/crawl"
)

func main() {
	fmt.Println("test")
	crawler := EasyCrawler{DepthCounter: 3}
	crawler.Crawl("http://spring-boot-reference.jp/")
}

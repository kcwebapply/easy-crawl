# easy-crawl
![Go](https://img.shields.io/badge/Language-Go-6699FF.svg)
![apache licensed](https://img.shields.io/badge/License-Apache_2.0-d94c32.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/easy-crawl)](https://goreportcard.com/report/github.com/kcwebapply/easy-crawl)

easy-crawl is library for crawling smoothly and set callback method easily.

## Installation

To install easy-crawl package, you need to install Go and set your Go workspace first.

1. Download and install it:

```sh
$ go get -u github.com/kcwebapply/easy-crawl
```

2. Import it in your code:

```go
import "github.com/kcwebapply/easy-crawl"
```


## Usage 
```Go
func main() {
  // initialize Easycrawler{} with crawling depth.
  crawler := easyCrawl.EasyCrawler{Depth: 3} 
  
  // you should implements CallBackInterface and set it in SetCallBack() method. 
  // CallBack() method is called when crawler get html contents by request .
  callBackImpl := CallBackImpl{}  
  crawler.SetCallBack(callBackImpl)
  
  // you can monitor how crawling is being done by call SetLogging() and set `true`.
  crawler.SetLogging(true)
  
  // crawling!
  crawler.Crawl("http://spring-boot-reference.jp/")
}


type CallBackImpl struct {
}

func (callbackImpl CallBackImpl) Callback(url string, urls []string, body string) {
   // implements as you like . 
}
```


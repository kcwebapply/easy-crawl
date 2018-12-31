package easyCrawl

import "fmt"

type Logging struct {
	logging bool
}

func (logging *Logging) logCrawlDone(url string) {
	if !logging.logging {
		return
	}
	fmt.Println("crawl url : ", url)
}

func (logging *Logging) logDepth(depthCounter int, urlNum int) {
	if !logging.logging {
		return
	}
	fmt.Println(depthCounter, "週目！", " 長さは", urlNum)
}

func (logging *Logging) logTime(time float64) {
	if !logging.logging {
		return
	}
	fmt.Printf("実行時間%f秒\n", time)
}

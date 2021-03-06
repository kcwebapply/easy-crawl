package easycrawl

import "fmt"

// Logging struct implements logging methods.
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
	fmt.Println("depth : ", depthCounter, " numberOfContents : ", urlNum)
}

func (logging *Logging) logTime(time float64) {
	if !logging.logging {
		return
	}
	fmt.Printf("crawling done.  %f seconds\n", time)
}

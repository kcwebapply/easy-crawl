package crawl

type CallBackInterface interface {
	Callback(url string, urls []string, body string)
}

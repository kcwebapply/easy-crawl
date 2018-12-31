package easyCrawl

type CallBackInterface interface {
	Callback(url string, urls []string, body string)
}

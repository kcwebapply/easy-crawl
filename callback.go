package easycrawl

// CallBackInterface is interface for setting callback method as user like.
type CallBackInterface interface {
	Callback(url string, urls []string, body string)
}

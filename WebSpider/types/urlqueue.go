package types

// URLQueue represents a queue of URLs.
// URLQueue represents a queue of URLs.
type URLChanQueue struct {
	URLch   chan string
	Visited Set
}

func NewURLQueue(initialURL string) *URLChanQueue {
	visited := make(Set)
	urlch := make(chan string, 32)
	urlch <- initialURL
	return &URLChanQueue{
		URLch:   urlch,
		Visited: visited,
	}
}

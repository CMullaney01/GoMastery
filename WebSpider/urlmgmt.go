package main

// URLQueue represents a queue of URLs.
// URLQueue represents a queue of URLs.
type URLChanQueue struct {
	URLch   chan string
	visited Set
}

func NewURLQueue(initialURL string) *URLChanQueue {
	visited := make(Set)
	return &URLChanQueue{
		URLch:   make(chan string, 32),
		visited: visited,
	}
}

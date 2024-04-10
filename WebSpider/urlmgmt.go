package main

type visitedURL struct {
	visited Set
}

// URLQueue represents a queue of URLs.
type URLQueue struct {
	URLs []string
}

func NewURLQueue(initialURL string) *URLQueue {
	return &URLQueue{
		URLs: make([]string, 0),
	}
}

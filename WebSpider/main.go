package main

// what do I want to crawl?
func main() {
	initialURL := "google.com"
	urls := NewURLQueue(initialURL)

	Crawl(urls)
}

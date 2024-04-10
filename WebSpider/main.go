package main

// what do I want to crawl?
func main() {
	initialURL := "google.com"
	urls := NewURLQueue(initialURL)

	parser := &DefaultHTMLParser{}
	fetcher := &DefaultFetcher{}

	Crawl(urls, parser, fetcher)

	close(urls.URLch)
}

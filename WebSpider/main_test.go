package main

import (
	"testing"
	"webspider/crawler"
	"webspider/fetchers"
	"webspider/parsers"
	"webspider/types"
)

// TestMainCrawl tests the main function's crawling process
func TestMainCrawl(t *testing.T) {
	initialURL := "google.com"
	urls := types.NewURLQueue(initialURL)
	parser := parsers.NewMockParser()
	fetcher := fetchers.NewMockFetcher()
	// Call the main function
	crawler.Crawl(urls, parser, fetcher)

	// Check if the expected number of URLs have been visited
	expectedVisitedURLs := 32
	if urls.Visited.Size() != expectedVisitedURLs {
		t.Errorf("Unexpected number of visited URLs. Expected: %d, Got: %d", expectedVisitedURLs, urls.Visited.Size())
	}
	close(urls.URLch)
	// Add more assertions as needed based on your application's behavior
}

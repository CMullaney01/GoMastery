package main

import (
	"testing"
	"webspider/fetchers"
	"webspider/parsers"
	"webspider/types"
)

// TestMainCrawl tests the main function's crawling process
func TestMainCrawl(t *testing.T) {
	initialURL := "google.com"
	urls := types.NewURLQueue(initialURL)
	parser := parsers.NewDefaultParser()
	fetcher := fetchers.NewFakeFetcher() // Use the fake fetcher to get example data
	// Call the main function
	Crawl(urls, parser, fetcher)

	// Check if the expected number of URLs have been visited
	expectedVisitedURLs := 1 // Since we have only one fake URL in the test
	if urls.Visited.Size() != expectedVisitedURLs {
		t.Errorf("Unexpected number of visited URLs. Expected: %d, Got: %d", expectedVisitedURLs, urls.Visited.Size())
	}

	// Add more assertions as needed based on your application's behavior
}

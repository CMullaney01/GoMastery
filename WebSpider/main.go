package main

import (
	"fmt"
	"sync"
	"webspider/fetchers"
	"webspider/parsers"
	"webspider/types"
)

// what do I want to crawl?
func Crawl(urls *types.URLChanQueue, parser parsers.HTMLParser, fetcher fetchers.Fetcher) {
	var wg sync.WaitGroup
	defer wg.Wait()

	// Channel to signal when to stop processing URLs
	stopCh := make(chan struct{})
	defer close(stopCh)

	for url := range urls.URLch {
		// Check if the URL has been visited
		if urls.Visited.Contains(url) {
			continue
		}

		// Check if we've visited 32 URLs
		if urls.Visited.Size() >= 32 {
			// Signal to stop processing URLs
			close(stopCh)
			break
		}

		// Add the URL to visited list
		urls.Visited.Add(url)

		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			select {
			case <-stopCh:
				// Stop processing if the stop signal is received
				return
			default:
				body, err := fetcher.Fetch(url)
				if err != nil {
					fmt.Println("Error fetching URL:", url, err)
					return // Skip to the next URL if there's an error
				}
				parser.ExtractURLs(body, urls.URLch)
				parser.ParseBody(body)
				fmt.Println("URL:", url)
				fmt.Println("Body:", body)
			}
		}(url)
	}
}

func main() {
	initialURL := "google.com"
	urls := types.NewURLQueue(initialURL)
	parser := parsers.NewDefaultParser()
	fetcher := fetchers.NewDefaultFetcher()

	Crawl(urls, parser, fetcher)

	close(urls.URLch)
}

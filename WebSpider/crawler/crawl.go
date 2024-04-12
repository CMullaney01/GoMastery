package crawler

import (
	"fmt"
	"log"
	"sync"
	"time"
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

	// Timeout duration
	timeoutDuration := time.Second

	// Initialize lastVisited time
	lastVisited := time.Now()

	for url := range urls.URLch {
		// Check if timeout duration has passed since the lastVisited time
		//To Do: This currently doesnt work
		if time.Since(lastVisited) > timeoutDuration {
			log.Println("Timeout occurred")
			stopCh <- struct{}{}
			break
		}
		// Check if the URL has been visited
		log.Println(url)
		if urls.Visited.Contains(url) {
			continue
		}

		// Check if we've visited 32 URLs
		if urls.Visited.Size() >= 32 {
			log.Println("max num Urls visited")
			// Signal to stop processing URLs
			stopCh <- struct{}{}
			break
		}

		// Add the URL to visited list and update lastVisited time
		urls.Visited.Add(url)
		lastVisited = time.Now()

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

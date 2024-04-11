package main

import (
	"fmt"
	"webspider/fetchers"
	"webspider/parsers"
	"webspider/types"
)

// what do I want to crawl?
func Crawl(urls *types.URLChanQueue, parser parsers.HTMLParser, fetcher fetchers.Fetcher) {
	for url := range urls.URLch {
		if urls.Visited.Contains(url) {
			continue
		}
		urls.Visited.Add(url)
		go func(url string) {
			body, err := fetcher.Fetch(url)
			if err != nil {
				fmt.Println("Error fetching URL:", url, err)
				return // Skip to the next URL if there's an error
			}
			parser.ExtractURLs(body, urls.URLch)
			parser.ParseBody(body)
			fmt.Println("URL:", url)
			fmt.Println("Body:", body)
		}(url)
		//prevent infinite crawling -- only visit 32 URLS
		if urls.Visited.Size() > 32 {
			break
		}
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

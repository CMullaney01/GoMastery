package main

import "fmt"

// what do I want to crawl?
func Crawl(urls *URLChanQueue, parser HTMLParser, fetcher Fetcher) {
	for url := range urls.URLch {
		if urls.visited.Contains(url) {
			continue
		}
		urls.visited.Add(url)
		go func(url string) {
			body, err := fetcher.Fetch(url, urls.URLch, parser)
			if err != nil {
				fmt.Println("Error fetching URL:", url, err)
				return // Skip to the next URL if there's an error
			}
			fmt.Println("URL:", url)
			fmt.Println("Body:", body)
		}(url)
		//prevent infinite crawling -- only visit 32 URLS
		if urls.visited.Size() > 32 {
			break
		}
	}
}

func main() {
	initialURL := "google.com"
	urls := NewURLQueue(initialURL)
	parser := NewDefaultParser()
	fetcher := NewDefaultFetcher()

	Crawl(urls, parser, fetcher)

	close(urls.URLch)
}

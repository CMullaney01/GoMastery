package main

import (
	"webspider/crawler"
	"webspider/fetchers"
	"webspider/parsers"
	"webspider/types"
)

func main() {
	initialURL := "google.com"
	urls := types.NewURLQueue(initialURL)
	parser := parsers.NewDefaultParser()
	fetcher := fetchers.NewDefaultFetcher()

	crawler.Crawl(urls, parser, fetcher)

	close(urls.URLch)
}

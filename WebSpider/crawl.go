package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// HTMLParser interface for parsing HTML content
type HTMLParser interface {
	ExtractURLs(content string, urlch chan<- string)
	ParseBody(content string)
}

// Fetcher interface for fetching web pages -- great for allowing us to use a fake fetcher for testing
type Fetcher interface {
	Fetch(url string, urlch chan<- string, parser HTMLParser) (body string, err error)
}

type DefaultHTMLParser struct{}
type DefaultFetcher struct{}

func (p *DefaultHTMLParser) ExtractURLs(content string, urlch chan<- string) {
	tokenizer := html.NewTokenizer(strings.NewReader(content))

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					urlch <- attr.Val
				}
			}
		}
	}
}

func (p *DefaultHTMLParser) ParseBody(content string) {}

// We want this function to return the body of the website and find new URLS adding them to our channel
func (f *DefaultFetcher) Fetch(url string, urlch chan<- string, parser HTMLParser) (body string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Parse HTML content to extract URLs
	parser.ExtractURLs(string(content), urlch)

	return string(content), nil
}

// Given the current url implement Crawling Logic what would be cool is if we could use a llm to process the information in the body
// Also might be cool to crawl each url in the set and get our favourite urls, assign a class of importance (how many urls reference)
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

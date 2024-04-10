package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// extractURLs parses HTML content and extracts URLs
func extractURLs(content string) []string {
	urls := make([]string, 0)
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
					urls = append(urls, attr.Val)
				}
			}
		}
	}

	return urls
}

// we want this fuinction to return the body of the website url and all urls in the current body
func fetch(url string) (body string, urls []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	// Parse HTML content to extract URLs
	urls = extractURLs(string(content))

	return string(content), urls, nil
}

// Given the current url implement Crawling Logic what would be cool is if we could use a llm to process the information in the body
// Also might be cool to crawl each url in the set and get our favourite urls, assign a class of importance (how many urls reference)
func Crawl(urls *URLQueue) {
	for _, url := range urls.URLs {
		body, fetchedURLs, err := fetch(url)
		if err != nil {
			fmt.Println("Error fetching URL:", url, err)
			continue // Skip to the next URL if there's an error
		}
		fmt.Println("URL:", url)
		fmt.Println("Body:", body)
		fmt.Println("Fetched URLs:", fetchedURLs)
	}
}

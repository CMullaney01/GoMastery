package main

import (
	"io"
	"log"
	"net/http"
)

// Fetcher interface for fetching web pages -- great for allowing us to use a fake fetcher for testing
type Fetcher interface {
	Fetch(url string, urlch chan<- string, parser HTMLParser) (body string, err error)
}

type DefaultFetcher struct{}

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

func NewDefaultFetcher() Fetcher {
	return &DefaultFetcher{}
}

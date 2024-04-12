package fetchers

import (
	"io/ioutil"
	"log"
)

// MockFetcher is a Mock implementation of the Fetcher interface
type MockFetcher struct{}

// Fetch simulates fetching a web page by reading its content from a file
func (f *MockFetcher) Fetch(url string) (body string, err error) {
	// Path to the Mock HTML page
	MockPagePath := "./testdata/Mockpage.html"

	// Read the content of the Mock HTML page
	content, err := ioutil.ReadFile(MockPagePath)
	if err != nil {
		log.Println("Error reading Mock page:", err)
		return "", err
	}

	return string(content), nil
}

// NewMockFetcher creates a new instance of MockFetcher with the specified file path
func NewMockFetcher() *MockFetcher {
	return &MockFetcher{}
}

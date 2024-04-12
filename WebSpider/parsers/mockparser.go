package parsers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type MockParser struct {
	urls  []string
	index int
}

func (p *MockParser) ExtractURLs(content string, urlch chan<- string) {
	end := p.index + 3
	if end > len(p.urls) {
		end = len(p.urls)
	}
	for _, url := range p.urls[p.index:end] {
		urlch <- url
	}
	p.index = end
}

func (p *MockParser) ParseBody(content string) {}

func NewMockParser() HTMLParser {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	log.Printf("Current directory: %s", currentDir)
	urlPath := "./parsers/testdata/mockurls.json"
	data, err := ioutil.ReadFile(urlPath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Define a struct to unmarshal the JSON data
	var jsonData struct {
		URLs []string `json:"urls"`
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON data: %v", err)
	}

	// Print the URLs
	// log.Println("URLs:")
	// for _, url := range jsonData.URLs {
	// 	log.Println(url)
	// }

	// Add the URLs to a []string slice
	urlSlice := make([]string, len(jsonData.URLs))
	copy(urlSlice, jsonData.URLs)

	return &MockParser{
		urls:  urlSlice,
		index: 0,
	}
}

package parsers

import (
	"strings"

	"golang.org/x/net/html"
)

//unlike food parser, we want to implement our own functions for both the parse body and the Extract URLs, this is done here
type SportHTMLParser struct{}

func (p *SportHTMLParser) ExtractURLs(content string, urlch chan<- string) {
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

func (p *SportHTMLParser) ParseBody(content string) {}

func NewSportParser() HTMLParser {
	return &SportHTMLParser{}
}

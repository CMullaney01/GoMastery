package parsers

import (
	"strings"

	"golang.org/x/net/html"
)

// HTMLParser interface for parsing HTML content
type HTMLParser interface {
	ExtractURLs(content string, urlch chan<- string)
	ParseBody(content string)
}

type DefaultHTMLParser struct{}

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

func NewDefaultParser() HTMLParser {
	return &DefaultHTMLParser{}
}

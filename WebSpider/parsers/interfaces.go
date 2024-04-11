package parsers

// HTMLParser interface for parsing HTML content
type HTMLParser interface {
	ExtractURLs(content string, urlch chan<- string)
	ParseBody(content string)
}

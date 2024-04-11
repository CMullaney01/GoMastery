package fetchers

// Fetcher interface for fetching web pages -- great for allowing us to use a fake fetcher for testing
type Fetcher interface {
	Fetch(url string) (body string, err error)
}

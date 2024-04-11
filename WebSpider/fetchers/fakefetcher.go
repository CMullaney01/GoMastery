package fetchers

type FakeFetcher struct{}

// Fake fetcher function used for testing
func (f *FakeFetcher) Fetch(url string) (body string, err error) {
	return "", nil
}

func NewFakeFetcher() Fetcher {
	return &FakeFetcher{}
}

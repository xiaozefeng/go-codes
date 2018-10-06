package main

import (
	"fmt"
	"log"
)

func main() {
}

type Fetcher interface {

	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum  of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// todo: fetch urls in parallel.
	// todo: don't fetch the same url twice.
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Printf("found: %s, %q\n", url, body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
}

type fakeFetcher struct{}

func (fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	panic("implement me")
}


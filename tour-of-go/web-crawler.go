package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var visitedUrls = make(Cache) // just a concurrency-safe map w/ get() & get() methods
var mu sync.Mutex

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if visitedUrls.get(url) == true {
		fmt.Printf("Already visited %v. Skipping...\n", url)
		return
	}
	fmt.Printf("Crawling %v...\n", url)
	visitedUrls.set(url, true)

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Found: %s %q\n", body, urls)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

type Cache map[string]bool

func (cache Cache) get(value string) bool {
	mu.Lock()
	defer mu.Unlock()
	return cache[value]
}

func (cache Cache) set(key string, value bool) {
	mu.Lock()
	defer mu.Unlock()
	cache[key] = value
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		body: "The Go Programming Language", // these specifers seem optional
		urls: []string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		body: "Packages",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

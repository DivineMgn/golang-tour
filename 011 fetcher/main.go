// https://go-tour-ru-ru.appspot.com/concurrency/10

package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	Crawl("http://golang.org/", 6, fetcher)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	cache := &URLCache{data: make(map[string]bool, 0)}
	results := make(chan string)

	waiter := &sync.WaitGroup{}

	waiter.Add(1)
	go _crawlInDepth(url, depth, depth, fetcher, results, cache, waiter)

	go func() {
		waiter.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func _crawlInDepth(url string, totalDepth, depth int, fetcher Fetcher, results chan<- string, cache *URLCache, waiter *sync.WaitGroup) {
	defer waiter.Done()

	if depth <= 0 || cache.Contains(url) {
		return
	}

	cache.Add(url)

	body, urls, err := fetcher.Fetch(url)

	ident := strings.Repeat("#", totalDepth-depth)
	if err != nil {
		results <- fmt.Sprintf("%s not found: %s", ident, url)
		return
	}

	results <- fmt.Sprintf("%s found: %s %s", ident, url, body)

	for _, childURL := range urls {
		waiter.Add(1)
		go _crawlInDepth(childURL, totalDepth, depth-1, fetcher, results, cache, waiter)
	}
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// Fetcher - some class
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

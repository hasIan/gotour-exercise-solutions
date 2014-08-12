// This is broken; I hate this exercise; it's a double-black diamond.

package main

import (
	"fmt"
)

var pr = fmt.Println

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	visited := make(map[string]bool)
	finished := make(chan bool)

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	finished <- true
	go doCrawl(url, depth, fetcher, visited, finished)
	<-finished

}

func doCrawl(
	url string,
	depth int,
	fetcher Fetcher,
	visited map[string]bool,
	finished chan bool) {

	if depth <= 0 {
		return
	}

	if visited[url] {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	visited[url] = true

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	inflight := 0
	inflightChan := make(chan bool)

	for _, u := range urls {
		inflight++
		go doCrawl(u, depth-1, fetcher, visited, inflightChan)
	}

	for ; inflight > 0; inflight-- {
		<-inflightChan
	}

	finished <- true
}

func alreadyVisited(candidateUrl string, visited []string) bool {
	for _, url := range visited {
		if url == candidateUrl {
			return true
		}
	}

	return false
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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

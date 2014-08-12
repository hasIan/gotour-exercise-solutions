// This is broken; I hate this exercise; it's a double-black diamond.

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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	state := newState()
	defer state.close()
	go doCrawl(url, depth, fetcher, state)
	state.wait()
	fmt.Println("Finished crawling", url)
}

type state struct {
	mu       *sync.RWMutex
	visited  map[string]struct{}
	finishCh chan struct{}
}

func newState() *state {
	return &state{
		mu:       new(sync.RWMutex),
		visited:  make(map[string]struct{}),
		finishCh: make(chan struct{}),
	}
}

func (s *state) isVisited(url string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.visited[url]
	return exists
}

func (s *state) finish(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.visited[url] = struct{}{}
	s.finishCh <- struct{}{}
}

func (s *state) wait() {
	<-s.finishCh
}

func (s *state) close() {
	close(s.finishCh)
}

func doCrawl(url string, depth int, fetcher Fetcher, state *state) {
	defer state.finish(url)
	if depth <= 0 || state.isVisited(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	inflight := 0
	for _, u := range urls {
		inflight++
		go doCrawl(u, depth-1, fetcher, state)
	}
	for ; inflight > 0; inflight-- {
		state.wait()
	}
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

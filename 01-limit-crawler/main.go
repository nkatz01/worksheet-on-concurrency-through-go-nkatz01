
/*
 Your task is to change the code to limit the crawler to at most one
 page per second, while maintaining concurrency (in other words,
 Crawl() must be called concurrently)
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

 const rateLimit = time.Second

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a real crawler.
// It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup, throttle <-chan time.Time) {


defer	wg.Done()


	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {

		// Do not remove the `go` keyword, as Crawl() must be called concurrently
	 	<-throttle
		go Crawl(u, depth-1, wg,  throttle )
	}

//	<-throttle


	return
}

func main() {
	var wg sync.WaitGroup
	//const rateLimit = time.Second
	//throttle := time.Tick(rateLimit)

	wg.Add(1)
//	<-throttle
	throttle := time.Tick(rateLimit)
	Crawl("http://golang.org/", 4, &wg, throttle)

	wg.Wait()
}


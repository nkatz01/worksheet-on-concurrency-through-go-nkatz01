/*
 Given a producer-consumer scenario, where a producer reads in tweets from a
 mock stream and a consumer processes the data.

 Your task is to modify the code so that the producer as well
 as the consumer can run concurrently.
*/

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan<- Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweets)
			return
		}

		tweets <- *tweet
	}
}

func consumer(tweets <-chan Tweet) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}

}


func main() {
	start := time.Now()
	stream := GetMockStream()
	ch := make(chan Tweet)

	go producer(stream, ch)
	consumer(ch)

	fmt.Printf("Process took %s\n", time.Since(start))
}

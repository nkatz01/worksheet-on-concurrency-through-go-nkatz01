package main

import (
	"errors"
	"strings"
	"time"
)



// Tweet defines the simlified representation of a tweet
type Tweet struct {
	Username string
	Text     string
}
// Stream is a mock stream for demonstration purposes, not threadsafe
type Stream struct {
	pos    int
	tweets []Tweet
}

var mockdata = []Tweet{
	{
		"davecheney",
		"#golang top tip: if your unit tests import any other package you wrote, including themselves, they're not unit tests.",
	}, {
		"beertocode",
		"Backend developer, doing frontend featuring the eternal struggle of centering something. #coding",
	}, {
		"ironzeb",
		"Re: Popularity of Golang in China: My thinking nowadays is that it had a lot to do with this book and author https://github.com/astaxie/build-web-application-with-golang",
	}, {
		"beertocode",
		"Looking forward to the #gopher meetup in Hsinchu tonight with @ironzeb!",
	}, {
		"vampirewalk666",
		"I just wrote a golang slack bot! It reports the state of github repository. #Slack #golang",
	},
}

// GetMockStream returns a mock stream for demonstration purposes
func GetMockStream() Stream {
	return Stream{0, mockdata}
}


// ErrEOF returns on End of File error
var ErrEOF = errors.New("End of File")

// Next returns the next Tweet in the stream, returns EOF error if
// there are no more tweets
func (s *Stream) Next() (*Tweet, error) {

	// simulate delay
	time.Sleep(320 * time.Millisecond)
	if s.pos >= len(s.tweets) {
		return &Tweet{}, ErrEOF
	}

	tweet := s.tweets[s.pos]
	s.pos++

	return &tweet, nil
}



// IsTalkingAboutGo is a mock process which pretends to be a sophisticated
// procedure to analyse whether a particular tweet is referring to go (or not).
func (t *Tweet) IsTalkingAboutGo() bool {
	// simulate delay
	time.Sleep(330 * time.Millisecond)

	hasGolang := strings.Contains(strings.ToLower(t.Text), "golang")
	hasGopher := strings.Contains(strings.ToLower(t.Text), "gopher")

	return hasGolang || hasGopher
}



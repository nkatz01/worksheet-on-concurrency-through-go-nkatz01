package main

import (
	"sync"
	"time"

)

// User defines the UserModel.
// Use this to check whether a User is a Premium user or not.

type User struct {
	mu sync.Mutex
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users.
// Returns false if a process had to be killed
func worker(done chan struct{},process func()) {
	  process()
	close(done)
}

func HandleRequest(process func(), u *User) bool {


done := make(chan struct{})

	if u.IsPremium == false {

		if u.TimeUsed >= 10 {

			return false
		} else {

			timeDuration := 10*time.Second - time.Duration(u.TimeUsed*1E9)
			startTime := time.Now()

			timer1 := time.NewTimer(timeDuration)

			go worker(done, process)
			for {
				select {
				case <-timer1.C:
					u.mu.Lock()

					u.TimeUsed = 10;
					u.mu.Unlock()
					return false

				case <-done:
					 u.mu.Lock()

					u.TimeUsed = u.TimeUsed + int64(time.Now().Sub(startTime).Seconds())
					u.mu.Unlock()
					return true
				}
			}

		}
	} else {
			   process()
			  return true
		  }

	}


func main() {
	RunMockServer()
}

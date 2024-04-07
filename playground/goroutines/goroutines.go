package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Goroutine() {

	start := time.Now()
	// these could be http requests
	userName := fetchUser()
	respch := make(chan any, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikes(userName, respch, wg)
	go fetchMatch(userName, respch, wg)

	wg.Wait() // block until wg finished
	close(respch)

	for resp := range respch {
		fmt.Println("resp: ", resp)
	}

	// fmt.Println("likes: ", likes)
	// fmt.Println("match: ", match)
	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11
	wg.Done()
}

func fetchMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- "ANNA"
	wg.Done()
}

package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	// doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	// doneChan <- true
	// close(doneChan) // I close the connection here just because I know this is gonna be the slowest operation
}

func main() {
	done := make(chan bool)

	greet("Nice to meet you!", done)
	greet("How are you?", done)
	slowGreet("How ... are ... you ...?", done) // this is blocking the next execution until it gets resolved
	greet("I hope you're liking the course!", done)

	for range done {

	}
}

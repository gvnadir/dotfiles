package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
	// doneChan <- true
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	// doneChan <- true
	// close(doneChan) // I close the connection here just because I know this is gonna be the slowest operation
}

func main() {
	// done := make(chan bool)

	greet("Nice to meet you!")
	greet("How are you?")
	slowGreet("How ... are ... you ...?") // this is blocking the next execution until it gets resolved
	greet("I hope you're liking the course!")

	// for range done {

	// }
}

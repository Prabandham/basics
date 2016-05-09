package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Started Worker !!!!")
	time.Sleep(time.Second)
	fmt.Print("Done doing the job.")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	fmt.Println("Back in the main function now.")
	<-done
}

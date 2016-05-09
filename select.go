package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "This is the first message, This take one second to complete."
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "This is the second message, This will take two seconds to complete."
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("Got message from channel one", msg)
		case msg := <-c2:
			fmt.Println("Got message from channel two", msg)
		}
	}
}

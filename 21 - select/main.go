package main

import (
	"fmt"
	"time"
)

func message(c chan string, s string) {
	for {
		time.Sleep(500 * time.Millisecond)
		c <- s
	}
}
func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go message(chan1, "hi")
	go message(chan2, ":)")

	for {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1, "\n")
		case msg2 := <-chan2:
			fmt.Println(msg2, "\n")
		default:
			fmt.Println("Default", "\n")
		}
	}
}

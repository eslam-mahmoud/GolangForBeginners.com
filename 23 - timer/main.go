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
	// chan1 := make(chan string)
	// chan2 := make(chan string)

	// go message(chan1, "hi")
	// go message(chan2, ":)")

	// timer := time.NewTimer(2 * time.Second)
	// for {
	// 	select {
	// 	case msg1 := <-chan1:
	// 		fmt.Println(msg1, "\n")
	// 	case msg2 := <-chan2:
	// 		fmt.Println(msg2, "\n")
	// 	case <-timer.C:
	// 		fmt.Println("Timer ended")
	// 		return
	// 	}
	// }

	// timer := time.NewTimer(2 * time.Second)
	// <-timer.C
	// fmt.Println("Timer kicked")

	//  you can cancel the timer before it kick
	timer := time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("Timer kicked")
	}()
	stoped := timer.Stop()
	if stoped {
		fmt.Println("Timer stopped")
	}
	time.Sleep(3 * time.Second)
}

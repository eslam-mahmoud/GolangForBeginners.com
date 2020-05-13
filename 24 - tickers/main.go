// to do something repeatedly at regular intervals we use ticker
package main

import (
	"fmt"
	"time"
)

func echo(done <-chan bool, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Ticker tick ", t.Second())
			// get data from API
		case <-done:
			ticker.Stop()
			return
		}
	}
}

func main() {
	// ticker := time.NewTicker(1 * time.Second)
	// timer := time.NewTimer(5 * time.Second)
	// for {
	// 	select {
	// 	case t := <-ticker.C:
	// 		fmt.Println("Ticker tick ", t.Second())
	// 		// get data from API
	// 	case <-timer.C:
	// 		ticker.Stop()
	// 		return
	// 	}
	// }

	done := make(chan bool)
	go echo(done, time.Second*1)

	time.Sleep(5 * time.Second)
	done <- true
	fmt.Println("Ticker stopped")
}

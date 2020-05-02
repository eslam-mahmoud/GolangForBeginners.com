// waitGroups wait for multiple goroutines to finish
package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we’ll run in every goroutine. Note that a WaitGroup must be passed to functions by pointer.

func save(id int, wg *sync.WaitGroup) {
	// wg.Done() notify the WaitGroup that this call is finished
	// "defer" delay executing this till function "save()" return
	// remove 1 from go routines count in WaitGroup
	defer wg.Done()

	fmt.Printf("start saving user %d\n", id)
	time.Sleep(time.Second)
	fmt.Printf("finish savign user %d\n", id)
}

func main() {
	// waitGroup used to block this function from returning till all async actions finish
	var wg sync.WaitGroup

	// start 5 goroutines
	for i := 1; i <= 5; i++ {
		// tell the waitGroup that you will start another goroutine
		// adds 1 to the count of go routines running
		wg.Add(1)
		// start the routine and send wg as param to decrease the count when it finishs
		go save(i, &wg)
	}

	// wait till WaitGroup counter get to 0
	// Block the main function from exit
	wg.Wait()
	fmt.Println("main function will exit now")
}

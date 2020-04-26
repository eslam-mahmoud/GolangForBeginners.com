// A goroutine is a lightweight thread of execution managed by the Go runtime.
package main

import (
	"fmt"
	"time"
)

// print is a function that will print s 5 times
func print(s string) {
	for x := 0; x < 5; x++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("(%v : %v) ", s, x)
	}
}

func main() {
	// // call the funtion 2 times
	// // will print each word 5 times in order
	// print("Hello")
	// // "world" will not be printed till "Hello" finish
	// print("world")

	// // add "go " before function name to make it run in a thread
	// // this is running the function in Goroutine
	// // will execute both of them concurrently/in parallel/asynchronously
	// go print("Hello")
	// go print("world")

	// // another way to run code in a go routine
	// go func(s string) {
	// 	for x := 0; x < 5; x++ {
	// 		time.Sleep(100 * time.Millisecond)
	// 		fmt.Printf("(%v : %v) ", s, x)
	// 	}
	// }("Go")

	// // it is bad practise to access same variable from multiple routines
	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }
	//
	// this should be
	for i := 0; i < 5; i++ {
		go func(x int) {
			time.Sleep(time.Second * 2)
			fmt.Println(x)
		}(i)
	}

	// force the application to wait for 2 seconds before exit
	// we should be using WaitGroups
	time.Sleep(time.Second * 2)
}

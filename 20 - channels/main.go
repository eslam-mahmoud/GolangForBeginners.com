//channels are the way to comunicate between goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

// wait 2 seconds before sending on channel
func m(messages chan string) {
	time.Sleep(2 * time.Second)
	messages <- "Hi :)"
}

func main() {
	/***** Channels *****/
	// // init channel
	// messages := make(chan string)
	// // send channel as argument
	// go m(messages)
	// // code will be blocked till someone send something on that channel
	// msg := <-messages
	// fmt.Println(msg)

	// // Wrong way to use channels
	// // could will be blocked
	// messages := make(chan string)
	// // now code will be blocked wating for someone/somewhere else
	// // to read the message before compiler go to the next line
	// messages <- "Hi"
	// // will never get to this line
	// fmt.Println(<-messages)

	/***** Channel buffering *****/
	// // we can use buffered channels to take messages
	// // without anyone waiting on them
	// messages := make(chan string, 2)
	// messages <- "Hi"
	// messages <- ":)"
	// fmt.Println("code can execute without anyone reading the message yet")
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// /***** rang & close Channels *****/
	// messages := make(chan string, 2)
	// messages <- "Hi"
	// messages <- ":)"
	// //close(messages)
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	close(messages)
	// }()
	// fmt.Println("code can execute without anyone reading the message yet")
	// for message := range messages {
	// 	fmt.Println(message)
	// }
	// fmt.Println("after range")

	/***** Channel Direction *****/
	// using wait group to prevent the app from exiting before
	// goroutines finish
	var wg sync.WaitGroup
	ping := make(chan bool)
	pong := make(chan bool)

	go playPing(ping, pong, &wg)
	go playPong(ping, pong, &wg)

	// this will kik start one of the playPing()
	pong <- true
	wg.Wait()
}

// using "chan<- bool" in the params define this channel as write only channel
// using "<-chan bool" in the params define this channel as read only channel
func playPing(pingChan chan<- bool, pongChan <-chan bool, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		defer wg.Done()

		time.Sleep(500 * time.Millisecond)
		// wait to read on pongChan
		<-pongChan
		fmt.Println("Ping")
		pingChan <- true
	}
}
func playPong(pingChan <-chan bool, pongChan chan<- bool, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		defer wg.Done()

		time.Sleep(500 * time.Millisecond)
		<-pingChan
		fmt.Println("Pong")
		pongChan <- true
	}
}

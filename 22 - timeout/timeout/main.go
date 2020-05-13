package main

import (
	"fmt"
	"time"

	"timeout/database"
)

func main() {
	channel := make(chan bool)

	go database.Save("hi :)", channel)

	for {
		select {
		case <-channel:
			fmt.Println("Saved\n")
			return
		case <-time.After(3 * time.Second):
			fmt.Println("This is taking too long :(\n")
			return
		}
	}
}

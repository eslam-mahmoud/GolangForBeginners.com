package database

import "time"

// Save saves the string in DB
func Save(s string, receiver chan<- bool) {
	time.Sleep(2 * time.Second)
	receiver <- true
}

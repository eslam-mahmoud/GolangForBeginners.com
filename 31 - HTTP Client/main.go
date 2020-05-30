package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	// GET request
	{
		// get the data from previous example "30 - HTTP Server"
		resp, err := http.Get("http://localhost:8090/users")
		panicOnError(err)
		defer resp.Body.Close()

		// get stautus code
		fmt.Println("Status code:", resp.StatusCode)

		// Similar as example "27 - read file"
		reader := bufio.NewReader(resp.Body)
		s := bufio.NewScanner(reader)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		// s.Err() does not return io.EOF
		err = s.Err()
		panicOnError(err)
		// we can read it into struct similar to example "26 - JSON"
	}
	// HEAD request
	// head request is not enabled on that end point
	// will get status code 405 without error
	{
		// get the data from previous example "30 - HTTP Server"
		resp, err := http.Head("http://localhost:8090/users")
		panicOnError(err)
		defer resp.Body.Close()

		// get stautus code
		fmt.Println("Status code:", resp.StatusCode)
	}
	// Post request
	// create user
	{
		json := `{"name":"Mahmoud","expertise":["GoLang"]}`
		b := strings.NewReader(json)
		resp, err := http.Post("http://localhost:8090/users", "application/json", b)
		panicOnError(err)
		defer resp.Body.Close()

		// get stautus code
		fmt.Println("Status code:", resp.StatusCode)
	}
}

func panicOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

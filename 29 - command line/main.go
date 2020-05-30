package main

import (
	"fmt"
	"os"
)

func main() {
	// read command line arguments
	{
		// 0 index is the path to the command line
		args := os.Args
		path := os.Args[0]
		fmt.Printf("Path: %v\n", path)
		args = os.Args[1:]
		for i, v := range args {
			// value always string
			fmt.Printf("Index: %v, Value: %v\n", i, v)
		}

		// build app before running this example
		// go build main.go
		// ./main 1 eslam "PHP"
	}

	// read command line flags
	{
		// // return pointer
		// fName := flag.String("fname", "", "first name as string")
		// score := flag.Int("score", 0, "score as int")
		// start := flag.Bool("start", false, "start as boolian")

		// var lName string
		// flag.StringVar(&lName, "lname", "", "last name as string")

		// // parses from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
		// flag.Parse()

		// fmt.Printf("First name: %v, Score: %v, Last name: %v, Start: %v\n", *fName, *score, lName, *start)
		// fmt.Println("Rest of the arguments", flag.Args())

		// // RUN
		// // go run main.go -fname="Eslam" -lname="Mahmoud" -score=9 many other args
	}
	// Get values from env
	{
		// // getenv return string
		// fmt.Println("GAME_SCORE:", os.Getenv("GAME_SCORE"))
		// os.Setenv("GAME_SCORE", "1")
		// fmt.Println("GAME_SCORE:", os.Getenv("GAME_SCORE"))

		// fmt.Println("All env values")
		// for _, e := range os.Environ() {
		// 	envRow := strings.Split(e, "=")
		// 	fmt.Printf("Key: %v, Value: %v\n", envRow[0], envRow[1])
		// }
	}
}

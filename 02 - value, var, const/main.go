package main

import "fmt"

func main() {
	////////////////////////////////////
	// Testing print different values //
	////////////////////////////////////

	// string
	// fmt.Println("Hello world :)")
	// fmt.Println("1")
	// fmt.Println("Hello", "world", ":)")

	// int
	// fmt.Println(1)
	// fmt.Println(1, 2)
	// fmt.Println(4 / 2)

	// string & int
	// fmt.Println("hello", 2)
	// fmt.Println("hello", 2, 3)
	// fmt.Println("hello", 2, "world")

	// // float
	// fmt.Println("this is float number 7.0/2.0", 7.0/2.0)
	// fmt.Println("this is not float number 7/2 = ", 7/2)
	// fmt.Println("this is float number ", 7.0/2)
	// fmt.Println("this is float number ", 7/2.0)

	// bool
	// fmt.Println("this is a boolian", true)
	// fmt.Println("this is a boolian", false)

	///////////////////////
	// Testing variables //
	///////////////////////

	// String variable
	// var name string
	// name = "Eslam"
	// fmt.Println("hello", name)

	// Reassign value to string variable
	// var name string
	// name = "Eslam"
	// fmt.Println("hello", name)
	// name = "Mahmoud"
	// fmt.Println("hello", name)

	// Another way to define variable
	// var name = "Eslam"
	// fmt.Println("hello", name)
	// name = "Mahmoud"
	// fmt.Println("hello", name)

	// Redefining variable with the same name will cause error
	// var name = "Eslam"
	// fmt.Println("hello", name)
	// var name = "Mahmoud"
	// fmt.Println("hello", name)

	// Another way to define string variable
	// name := "eslam"
	// fmt.Println("hello", name)

	// // Error undefined: name
	// name = "eslam"
	// fmt.Println("hello", name)

	// Define int variable
	// score := 10
	// fmt.Println("player score", score)

	// // Changing the type of the value in variable will cause error
	// name := "Eslam"
	// fmt.Println("hello", name)
	// name = 5
	// fmt.Println("hello", name)

	// Different value types
	// game := "XO"         // string
	// players := 2         // int
	// gameStarted := false // boolian
	// fmt.Println("playing", game, "with", players, "players, game started", gameStarted)

	// default values
	// var name string
	// var number int
	// var floatNumber float32 // float64
	// var b bool
	// fmt.Println("string default value:", name)
	// fmt.Println("int default value:", number)
	// fmt.Println("float default value:", floatNumber)
	// fmt.Println("boolian default value:", b)

	// // Constant
	// const gameName = "XO"
	// fmt.Println("Game name is", gameName)

	// // Error cannot assign to gameName
	// const gameName = "XO"
	// fmt.Println("Game name is", gameName)
	// gameName = "game"
	// fmt.Println("Game name is", gameName)

	// Constant
	// const gameName string
	// gameName = "XO"
	// fmt.Println("Game name is", gameName)

	////////////////////
	// Advanced types //
	////////////////////

	// // Struct
	// type person struct {
	// 	name string
	// 	age  int
	// }
	// p := person{name: "Eslam"}
	// fmt.Println("Name:", p.name)

	// Channels
	messages := make(chan string)
	go func() { messages <- "Hi :)" }() // send message
	msg := <-messages                   // read message
	fmt.Println(msg)
}

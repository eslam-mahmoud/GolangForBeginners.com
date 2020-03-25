package main

import (
	"fmt"
)

func main() {
	playerName := "Mr. x"
	playerName = "Eslam"
	if playerName == "Eslam" {
		fmt.Println("Hi", playerName)
	}
	fmt.Println("Welcome")

	// playerName := "Mr. x"
	// // playerName := "Eslam"
	// if playerName == "Eslam" {
	// 	fmt.Println("Hi")
	// } else {
	// 	fmt.Println("Who are you?")
	// }

	// playerName := "Mr. Y"
	// if playerName == "Eslam" {
	// 	fmt.Println("Hi")
	// } else if playerName == "Mr. Y" {
	// 	fmt.Println("long time no see")
	// } else if playerName == "Mr. Y" {
	// 	fmt.Println("Hi", playerName)
	// } else {
	// 	fmt.Println("Who are you?")
	// }

	// playerName := "Mr. Y"
	// if playerName != "Eslam" {
	// 	fmt.Println("Who are you!")
	// }

	// if true {
	// 	fmt.Println("this was true")
	// } else {
	// 	fmt.Println("this is not true")
	// }

	// age := 21
	// if age == 21 {
	// 	fmt.Println("age is ok")
	// }

	// age := 21
	// if age >= 18 {
	// 	fmt.Println("You can drive a car")
	// }

	// playerName := "Mr. X"
	// age := 21
	// if age >= 18 && playerName == "Eslam" {
	// 	// if true && false {
	// 	fmt.Println("You can drive this car")
	// } else {
	// 	fmt.Println("You can not drive this car")
	// }

	/*
		Truth table
		true  && true = true
		true && false = false
		false && true = false
		false && false = false
	*/

	// playerName := "Mr. Y"
	// age := 21
	// if age > 18 && playerName == "Eslam" {
	// 	fmt.Println("You can drive this car")
	// } else if age > 18 && playerName == "Mr. X" {
	// 	fmt.Println("you need permission")
	// } else {
	// 	fmt.Println("You can not drive this car")
	// }

	// playerName := "Mr. X"
	// age := 21
	// if age >= 18 {
	// 	if playerName == "Eslam" {
	// 		fmt.Println("You can drive this car")
	// 	} else if playerName == "Mr. X" {
	// 		fmt.Println("you need permission")
	// 	} else {
	// 		fmt.Println("Who are you!")
	// 	}
	// } else {
	// 	fmt.Println("You can not drive this car")
	// }

	/*
		true || true = true
		true || false = true
		false || true = true
		false || false = false
	*/
	// playerName := "Mr. X"
	// age := 21
	// if age > 18 || playerName == "Eslam" {
	// 	fmt.Println("You can drive this car")
	// } else {
	// 	fmt.Println("You can not drive this car")
	// }

	// playerName := "Mr. X"
	// switch playerName {
	// case "Eslam":
	// 	fmt.Println("Hi")
	// 	fmt.Println("Eslam")
	// case "Ahmed":
	// 	fmt.Println("Hey")
	// default:
	// 	fmt.Println("Who")
	// }

	// age := 55
	// switch {
	// case age > 18:
	// 	fmt.Println("age > 18")
	// 	fallthrough
	// case age > 50:
	// 	fmt.Println("age > 50")
	// case age > 51:
	// 	fmt.Println("age > 51")
	// }

	// if time.Now().Weekday() == time.Friday || time.Now().Weekday() == time.Saturday {

	// }
	// switch time.Now().Weekday() {
	// case time.Friday, time.Saturday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's a weekday")
	// }

}

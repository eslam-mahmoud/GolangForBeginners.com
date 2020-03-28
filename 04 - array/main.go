// Arrays is a group of data/values with same type
// Arrays in go have preset size and can not expand

package main

import (
	"fmt"
)

func main() {
	// Here we create a players array, the size is 4 and type is string
	// so players is a variable that will have 4 strings in it
	// when we print the array as it is now we will get the default string value "" (empty string)
	var players [4]string
	fmt.Println(players)

	// Here we create an array the array size/length is 3 and type is int & default value will be zeros
	var a [3]int
	fmt.Println(a)

	// In all programming language that i know we start counting from 0
	// if array length is 4 then first index start at position 0
	// to assign any value to any index you can do it like this
	// we can skip any index and it will have the default value
	players[0] = "Eslam"
	players[1] = "Mahmoud"
	players[3] = "Mr.Y"
	fmt.Println(players)

	// another way to set the values in array is when you define it
	players = [4]string{"Eslam", "Mahmoud", "Mr.X", "Mr.Y"}
	fmt.Println(players)

	// to use short definition we can do it like this
	// with or without values
	names := [2]string{}
	fmt.Println(names)

	// if at any point in your application you needed to get the length of array as variable
	// we can use `len()` function
	playersLen := len(players)
	fmt.Println("players array length is:", playersLen)

	// Working with 2D array
	// 2d array help you have group of groups of data
	// or you can think of it as a thing to help you create 2D structure example 3x3 XO game board
	// this mean we have array of length 3 and type of (array of length 3 and type string)
	xoBoard := [3][3]string{}
	fmt.Println(xoBoard)
	xoBoard = [3][3]string{
		[3]string{"-", "-", "x"},
		[3]string{"-", "o", "-"},
		[3]string{"x", "o", "-"}, // note we need to add that , to enable new line
	}
	// fmt.Println(xoBoard)
	fmt.Println(xoBoard[0])
	fmt.Println(xoBoard[1])
	fmt.Println(xoBoard[2])

	// To change value in 2d Array we change the access it by x,y position
	// or you can say by row,colomn position
	// 0,0 is the top left element
	xoBoard[0][1] = "x"
	// fmt.Println(xoBoard)
	fmt.Println(xoBoard[0])
	fmt.Println(xoBoard[1])
	fmt.Println(xoBoard[2])
}

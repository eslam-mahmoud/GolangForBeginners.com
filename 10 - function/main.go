// Functions are the way to group and re use same logic/code
// functions take zero or more arguments, and return zero or more
package main

import (
	"errors"
	"fmt"
)

// function that does not return anything
func p(s string) {
	fmt.Println(s)
}

// function that does not return anything
// and take as many params as we send and print parameters length
func l(params ...interface{}) {
	fmt.Println(len(params))
}

// function that take two int arguments and return the sum
func add(x int, y int) int {
	return x + y
}

// lets group types if it is the same
func sum(x, y int) int {
	return x + y
}

// return multiple results
func concat(a, b string) (string, int) {
	r := a + b
	return r, len(r)
}

// named return types
func concat2(a, b string) (r string, l int) {
	r = a + b
	l = len(r)
	// will return the named vars
	return

	// this as if we did
	// var r string
	// var l int
	// r = a + b
	// l = len(r)
	// return r, l
}

// function that return string & nil or default string value & error
// this pattern is widely used in Go
func returnError(b bool) (string, error) {
	if b {
		return "", errors.New("This is an error")
	}
	return "Ok", nil
}

func main() {
	p("Hello world")

	l("Hello world", 2, "hi", true, 4.4)

	v := add(1, 6)
	fmt.Println("add:", v)

	fmt.Println("sum:", sum(1, 2))

	r, l := concat("Eslam", "Mahmoud")
	fmt.Println("concat:", r, l)

	r, l = concat2("Eslam", "Mahmoud")
	fmt.Println("concat2:", r, l)

	r, err := returnError(true)
	if err != nil {
		fmt.Println("We have error:", err.Error())
	} else {
		fmt.Println("We got:", r)
	}

	// function as type
	var x func(int, int) int
	x = sum
	fmt.Println("sum:", x(1, 2))
}

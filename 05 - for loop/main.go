// For loop gives us the ability to loop and execute something multiple times
package main

import "fmt"

func main() {
	// classic most basic type of loops
	// you defind start point/number; max muber for how many times the loop will run; step
	fmt.Println("for i := 0; i <= 4; i++")
	for i := 0; i <= 4; i++ {
		fmt.Println("Hi", i)
	}

	// Another way to run loops
	// it will be your responsibility to handle the step and the init value
	fmt.Println("for x <= 3")
	x := 1
	for x <= 3 {
		fmt.Println(x)
		// x = x + 1
		x++
	}
	fmt.Println("X value after the loop", x)

	// Infinite loop
	// for {
	// 	fmt.Println("will run forever")
	// }

	// We can skip just one step by using "continue"
	// will jump to the begining of the next iteration
	fmt.Println("if a%2 == 0")
	for a := 0; a <= 5; a++ {
		if a == 3 {
			continue
		}
		fmt.Println(a)
	}

	// How to break loop (stop it at any point and exit)
	// "break" exit the loop and continue executing the rest of the code after the loop
	for {
		fmt.Println("will run only ones, because of break")
		break
	}

	// loop within loop
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println("x, y:", x, y)
		}
	}

	// break loop within loop
	fmt.Println("break loop within loop")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				break
			}
			fmt.Println("x, y:", x, y)
		}
	}

	// break loop within loop using labels
	fmt.Println("break loop within loop using labels")
breakpoint:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				break breakpoint
			}
			fmt.Println("x, y:", x, y)
		}
	}

	// break loop within loop using var
	fmt.Println("break loop within loop using var")
	b := false
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				b = true
				break
			}
			fmt.Println("x, y:", x, y)
		}

		if b == true {
			break
		}
	}

	// "return" exit the loop and the outer function, which will cause the app to close in our case
	// "return" will exit all the loops if it is in a nested loop
	for {
		fmt.Println("will run only ones, because of return")
		return
	}

	// This line will never get displayed, because of the return in the previous example
	fmt.Println("This line will never get displayed")
}

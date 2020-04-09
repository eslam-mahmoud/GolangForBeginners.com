// variable scope is the definition of where a particular variable is accessible in the code
package main

import "fmt"

var name = "eslam"

func main() {
	// "name" is defined in file scope so it can be access from anywhere in the file
	fmt.Println(name)

	// // inside any function we can group some of the code using { }
	// a := 1
	// {
	// 	fmt.Println("A In", a)
	// 	{
	// 		fmt.Println("A Deep", a)
	// 	}
	// }
	// fmt.Println("A Out", a)

	// // if you edit variable in inner scope it will be reflected in the outer scope
	// b := 1
	// {
	// 	b = 2
	// 	fmt.Println("B In", b)
	// }
	// fmt.Println("B Out", b)

	// // if you redefine variable in inner scope it will act as new variable not related to the one in the outer scope
	// c := 1
	// {
	// 	c := 2
	// 	fmt.Println("C In", c)
	// }
	// fmt.Println("C Out", c)

	// // this will cause errro "error undefined: x" as var x is not fully defined yet
	// x, y := 1, x

	// // this will work
	// x := 1
	// y := x
	// fmt.Println("X, Y", x, y)

	// // This will cause error "undefined: d" as we can not access variable that was not defined yet
	// fmt.Println("D", d)
	// d := "anything"

	// // Types declared any where in the outer scope can be access from within the function
	// // note does not matter where in the file it was defined (before/after this function)
	// g := game{
	// 	started: false,
	// 	winner: player{
	// 		name:  "Eslam",
	// 		score: 100,
	// 	},
	// }
	// fmt.Println("G", g)

	// // can not access private types created in other functions
	// // this will return error "undefined: privateType"
	// t()
	// t := privateType{}
}

type game struct {
	started bool
	winner  player
}

type player struct {
	name  string
	score int
}

func t() {
	type privateType struct {
	}
	t := privateType{}
	_ = t
}

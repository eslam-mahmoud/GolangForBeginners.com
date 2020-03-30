// Slice is a strong widely used data type in go
// it is a more flexible than arrays, dynamically-sized
// it does not store any data by itself, it describes a section of an underlying array
package main

import "fmt"

func main() {
	// unlike array we do not need to set size
	keywords := []string{}
	fmt.Println(keywords)

	// can be defined in another way
	var x []string
	fmt.Println(x)

	// we can init slice with values
	src := []string{"x", "y", "z"}
	fmt.Println(src)

	// we can get/set the values by index
	fmt.Println(src[0])
	src[1] = "b"
	fmt.Println(src)

	// unlike arryas we extend slice
	src = append(src, "d")
	fmt.Println(src)

	// Another way to init slice with predefined size
	dest := make([]string, 2)
	fmt.Println(dest)

	// we can copy value between slices
	// if the destination is smaller than the source in size
	// copy will fill up the destination then stops
	copy(dest, src)
	fmt.Println(dest)
	// if we changes the source after copy it does not affect the destination
	src[0] = "a"
	fmt.Println(src)
	fmt.Println(dest)

	// Since Slice is mainly a pointer to underline array
	// slicing a slice does not create a copy of values
	// this will create a slice that point to the same values in the original array that "src" was pointing at
	// s will point to the indexes 1,2 only. index 3 is excluded
	s := src[1:3]
	fmt.Println(s)
	// updating the source will affect all the slices from source
	src[2] = "c"
	fmt.Println(src)
	fmt.Println(s)

	// src[:x] will create a slice point to values start at index 0 end at x-1
	// src[x:y] will create a slice point to values start at index x end at y-1
	// src[x:] will create a slice point to values start at index x end at len(s)-1
	fmt.Println(src)
	fmt.Println("src[:3]", src[:3])
	fmt.Println("src[1:3]", src[1:3])
	fmt.Println("src[2:]", src[2:])

	// 2D slice is simillar to 2D array, but it can expand
	gameMap := [][]string{
		make([]string, 3),
		make([]string, 5),
		[]string{"-", "-"},
	}
	fmt.Println(gameMap)
	gameMap[0] = make([]string, 6)
	gameMap[1] = make([]string, 10)
	fmt.Println(gameMap)
	gameMap[1][2] = "-"
	fmt.Println(gameMap)

	// Good read
	// Go Slices: usage and internals https://blog.golang.org/slices-intro
}

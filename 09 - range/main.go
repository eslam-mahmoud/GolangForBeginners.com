// Range is used within the the for loop,
// iterates over a array, slice or map to get key/value
package main

import "fmt"

func main() {
	// Range on array
	arr := [4]string{"Eslam", "Mahmoud", "X", "Y"}
	for k, v := range arr {
		fmt.Println("Range on array:, key:", k, "Value:", v)
	}
	fmt.Println("\n")

	// Range on slice
	slc := []string{"Eslam", "Mahmoud", "X"}
	for key, value := range slc {
		fmt.Println("Range on slice:, key:", key, "Value:", value)
	}
	fmt.Println("\n")

	// Range on map, order is not guaranteed
	mp := map[string]string{
		"admin": "Eslam",
		"user":  "Mahmoud",
	}
	for key, value := range mp {
		fmt.Println("Range on map:, key:", key, "Value:", value)
	}
	fmt.Println("\n")

	// Range on map of slices
	mapOfSlices := map[string][]string{
		"Admins": []string{"Eslam", "Mahmoud"},
		"Users":  []string{"X", "Y", "Z"},
	}
	for key, value := range mapOfSlices {
		fmt.Println("Range on map of slices:, key:", key, "Value:", value)
		for k, v := range value {
			fmt.Println("Range on the slice in map of slices:, key:", k, "Value:", v)
		}
	}
}

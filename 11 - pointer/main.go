// Pointers help us send the memory address of variable around
// instead of sending copy from the variable
package main

import "fmt"

func main() {
	name := "Eslam"
	fmt.Println(name)

	// sending a copy of the variable value
	updateCopy(name)
	fmt.Println(name)

	// we send pointer to the variable by adding "&" before it
	updateRefrence(&name)
	fmt.Println(name)

	// Print a pointer will print the memory address
	fmt.Println("Pointer value", &name)

	// // We can not send send variable from type if function take pointer of the same type, and the other way around
	// // This will cause error "cannot use name (type string) as type *string in argument to updateRefrence"
	// updateRefrence(name)

	// if parameter is pointer we can send nil as argument
	defaultValueForPointer(nil)

	var pValue *string
	pValue = &name
	fmt.Println(*pValue)
}

func updateCopy(s string) {
	s = "Mahmoud"
}

func updateRefrence(s *string) {
	// adding "*" before variable dereference it, mean we are now using the original variable
	// Assigning a value to a dereferenced pointer changes the value at the referenced address.
	fmt.Println("will print the memory address", s)
	fmt.Println("will print the value after dereference using *s", *s)
	*s = "Mahmoud"

	// // This just change the variable s to point to another variable
	// n := "Mahmoud"
	// s = &n
}

func defaultValueForPointer(s *string) {
	fmt.Println("nil value:", s)
}

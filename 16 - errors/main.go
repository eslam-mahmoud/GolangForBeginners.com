// "error" is a type to represent any issue in the execution of function
// invalid input, problem connecting to external service ...
//
// "error" type is an interface that have one function
// type error interface {
// 	Error() string
// }
package main

import (
	"errors"
	"fmt"
)

// this is a function that will return error if got true as parameter
func getError(b bool) (err error) {
	if b == true {
		// err = errors.New("this is an error :(")
		err = fmt.Errorf("this is an error :(")
	}
	return
}

// custom error struct that implement error interface
// we use this to return custom dynamic error message
type invalidInput struct {
	filedName string
}

// implement "error" interface
func (i invalidInput) Error() string {
	return fmt.Sprintf("invlaid input on field '%v'", i.filedName)
}

// function that should validate input and return error
// it works because our custom type implement "error" interface
func validateInput(field, value string) (err error) {
	if value == "" {
		err = invalidInput{filedName: field}
	}
	return
}

func main() {
	// // using variable scope
	// {
	// 	// this is how we check on error
	// 	err := getError(true)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	// {
	// 	// call to validateInput function with valid input
	// 	err := validateInput("firstName", "eslam")
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }
	// {
	// 	// call to validateInput function with invalid input
	// 	err := validateInput("lastName", "")
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	// {
	// 	// check on error value
	// 	// err := invalidLength("")
	// 	err := invalidLength("error")
	// 	// err := getError(true)
	// 	if err == errInvalidLength {
	// 		fmt.Println("We know err value using 'err == errInvalidLength'")
	// 	} else {
	// 		fmt.Println("we do not know about this error")
	// 	}
	// }
	// {
	// 	// check on error type using "type assertion"
	// 	// this works only with custom types like invalidInput
	// 	// does not work with variable like errInvalidLength
	// 	err := validateInput("name", "")
	// 	_, ok := err.(invalidInput)
	// 	if ok {
	// 		fmt.Println("We know err type using 'type assertion'")
	// 	} else {
	// 		fmt.Println("we do not know about this error")
	// 	}
	// }

	// // Errors in Go 1.13
	// // https://blog.golang.org/go1.13-errors
	// {
	// 	// check on error value
	// 	// The errors.Is function compares an error to a value.
	// 	err := invalidLength("")
	// 	if errors.Is(err, errInvalidLength) {
	// 		fmt.Println("We know err value using 'errors.Is'")
	// 	} else {
	// 		fmt.Println("we do not know about this error")
	// 	}
	// }
	// {
	// 	// check on error type
	// 	err := validateInput("name", "")
	// 	if errors.As(err, &invalidInput{}) {
	// 		fmt.Println("We know err type using 'errors.As'")
	// 	} else {
	// 		fmt.Println("we do not know about this error")
	// 	}
	// }

	// // error unwrap from https://blog.golang.org/go1.13-errors
	// // Frequently a function passes an error up the call stack while adding information to it
	// // If e1.Unwrap() returns e2, then we say that e1 wraps e2, and that you can unwrap e1 to get e2.
	// {
	// 	err := fmt.Errorf("could not save your input: %w", errInvalidLength)
	// 	if errors.Is(err, errInvalidLength) {
	// 		fmt.Println("We know err value using 'errors.Is'", err.Error())
	// 		fmt.Println("this is the wrapped message:", errors.Unwrap(err))
	// 	} else {
	// 		fmt.Println("we do not know about this error")
	// 	}
	// }

	{
		// if we got errInvalidLength as return from a sub funtion
		// we warp it in our error and return it up the chain
		err := errSaving{err: errInvalidLength}
		// This should be in the parent function
		// check on the type using "errors.As" and unwarp the error
		// to get the sub error
		if errors.As(err, &errSaving{}) {
			fmt.Println("This is the parent error:'", err.Error())
			fmt.Println("This is the deep error:", errors.Unwrap(err))
		} else {
			fmt.Println("we do not know about this error")
		}
	}
}

type errSaving struct {
	err error
}

// implement "error" interface
func (e errSaving) Error() string {
	return fmt.Sprintf("error saving inputs")
}

func (e errSaving) Unwrap() error {
	return e.err
}

var errInvalidLength = errors.New("Invalid length")

func invalidLength(s string) error {
	if len(s) < 1 {
		return errInvalidLength
	} else if s == "error" {
		return errors.New("Invalid input")
	}
	return nil
}

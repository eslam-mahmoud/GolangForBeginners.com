// Interfaces are types, it define a set of method signatures
// that should be implemented by struct
package main

import (
	"errors"
	"fmt"
)

// walker type define something that can walk
type walker interface {
	walk(p point) error
	getPosition() point
}

// talker type define something that can talk
type talker interface {
	talk(s string)
}

type point struct {
	x int
	y int
}

// human struct type
type human struct {
	position point
}

// implement interface walker on human
// interface implementation comes implicitly when struct implement all the functions in interface type
func (h *human) walk(p point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("invalid point")
	}
	h.position = p
	fmt.Println("Human walked to", h.position)
	return nil
}
func (h *human) getPosition() point {
	return h.position
}

// implement interface talker on human
func (h *human) talk(s string) {
	fmt.Println("Human talking", s)
}

type animal struct {
	position point
}

// implement interface walker on animal
func (a *animal) walk(p point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("invalid point")
	}
	a.position = p
	fmt.Println("animal walked to", a.position)
	return nil
}
func (a *animal) getPosition() point {
	return a.position
}

func main() {
	eslam := &human{}
	dog := &animal{}

	steps := []point{
		point{1, 1},
		point{2, 2},
		point{3, 3},
	}
	err := move(eslam, steps)
	if err != nil {
		fmt.Println(err)
	}
	err = move(dog, steps)
	if err != nil {
		fmt.Println(err)
	}
	err = move(dog, []point{point{-1, 2}})
	if err != nil {
		fmt.Println(err)
	}

	// Functions should take interface as parameter type
	// that makes it work better with diferent type of sturcts
	// this will work
	moveHuman(eslam, steps)
	// this will cause errors
	// cannot use dog (type *animal) as type *human in argument to moveHuman
	// moveHuman(dog, steps)
}

// function takes interface as parameter type
// can take any struct type implements the interface
func move(w walker, points []point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

// function takes struct as type for the parameter
// can not take any other struct even if it implements the required functions
func moveHuman(w *human, points []point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(newUser("eslam", 10))
	fmt.Println(newUser("mahmoud", 11))
}

var id = 0

type productInterface interface {
	buy() error
	sell() error
}
type product struct {
	name string
}

func (p *product) buy() error {
	fmt.Println("buy product")
	return nil
}
func (p *product) sell() error {
	fmt.Println("sell product")
	return nil
}

type user struct {
	id       int
	name     string
	age      int
	products []product
}

func newUser(name string, age int) *user {
	id++
	return &user{
		id:   id,
		name: name,
		age:  age,
	}
}

func (u *user) buyProduct(p productInterface) error {
	return p.buy()
}

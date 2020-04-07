// Structs are a collection of fields.
// They’re useful for grouping data together to form records.
// the closest thing in Go to objects/OOP
package main

import "fmt"

type player struct {
	firstName string
	lastName  string
	score     int
}
type game struct {
	players map[string]*player
}

func (g *game) getWinner() *player {
	var maxScore int
	var winner *player
	for _, value := range g.players {
		if value.score > maxScore {
			winner = value
			maxScore = value.score
		}
	}
	return winner
}

type chain struct {
	base
	value int
	next *chain
}
func (c *chain) sayOk() {
	fmt.Println(":)")
}


type base struct {
}

func (b *base) sayHi() {
	fmt.Println("Hi")
}
func (b *base) sayOk() {
	fmt.Println("Ok")
}


func main() {
	g := &game{
		players: make(map[string]*player),
	}

	p1 := player{}
	fmt.Println("p1", p1)
	g.players["p1"] = &p1
	fmt.Println("game", g, "\n")

	var p2 player
	p2 = player{}
	p2.score = 5
	fmt.Println("p2", p2)
	g.players["p2"] = &p2
	fmt.Println("game", g, "\n")

	p3 := player{
		lastName:  "Mahmoud",
		firstName: "Eslam",
		// score will have the default value
	}
	fmt.Println("p3", p3)
	p3.score = 10
	fmt.Println("p3", p3)
	g.players["p3"] = &p3
	fmt.Println("game", g, "\n")


	fmt.Println("g.players[2]", g.players["p2"])
	fmt.Println("*g.players[2]", *g.players["p2"], "\n")

	winner := g.getWinner()
	fmt.Println("winner", winner)
	fmt.Println("*winner", *winner, "\n")

	// b := base{}
	// b.sayHi()
	// b.sayOk()

	c1 := chain{value:100}
	fmt.Println("c1", c1)
	c1.sayHi()
	c2 := &chain{value:200}
	c1.next = c2
	fmt.Println("c1", c1)
	fmt.Println("c1.next", c1.next)
	c1.next.sayHi()
	c1.sayOk()
}

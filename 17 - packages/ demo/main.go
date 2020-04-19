// Any Go program is made of packages.
// Programs start running in package main.
// Main package name defined in go.mod file
// Every sub directory is a sub package
package main

import (
	"demo/lcr"
	"fmt"
)

func main() {
	fmt.Println("Welcome to LCR dice game :D")
	g := lcr.NewGame()

	fmt.Println("Please enter how many players will play the game?")
	fmt.Println("Note: enter number more than 2.")

	// need players count and will take it as input
	var playersCount int
	for {
		fmt.Scanln(&playersCount)
		if playersCount < 3 {
			fmt.Println("Note: enter number more than 2.")
			continue
		}
		break
	}

	for i := 0; i < playersCount; i++ {
		playerName := fmt.Sprintf("P%v", i)
		p := g.Join(playerName)
		fmt.Println(fmt.Sprintf("player: %v joined.", p.Name()))
	}

	for {
		turn := g.Turn()

		// check if player have tokens
		// if not skip him
		if turn.Tokens() == 0 {
			fmt.Println(fmt.Sprintf("\nplayer %v, you have 0 tokens.", turn.Name()))
			continue
		}

		// if yes ask him to hit enter to roll dice
		fmt.Println(fmt.Sprintf("\nplayer %v, you have %v tokens. hit any key to roll dices", turn.Name(), turn.Tokens()))
		var playerInput string
		fmt.Scanln(&playerInput)
		if playerInput == "Exit" {
			fmt.Println("You killed the game :(")
			return
		}

		// roll dice & apply chenges
		dicesResult := turn.RollDice()
		fmt.Println("You got:", dicesResult)
		for _, p := range g.Players() {
			fmt.Println(fmt.Sprintf("player %v, have %v tokens", p.Name(), p.Tokens()))
		}

		// check if any one won the game
		// exit
		winner := g.Finished()
		if winner != nil {
			fmt.Println("\nWinner: ", winner.Name())
			return
		}
	}
}

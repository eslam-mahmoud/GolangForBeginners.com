package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Welcome to LCR dice game :D")
	g := new()

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
		p := g.join(playerName)
		fmt.Println(fmt.Sprintf("player: %v joined.", p.name))
	}

	turn := g.players[0]
	for {
		// check if player have tokens 
		// if not skip him
		if turn.tokens == 0 {
			fmt.Println(fmt.Sprintf("\nplayer %v, you have 0 tokens.", turn.name))
			turn = turn.right
			continue
		}

		// if yes ask him to hit enter to roll dice
		fmt.Println(fmt.Sprintf("\nplayer %v, you have %v tokens. hit any key to roll dices", turn.name, turn.tokens))
		var playerInput string
		fmt.Scanln(&playerInput)
		if playerInput == "Exit" {
			fmt.Println("You killed the game :(")
			return
		}

		// roll dice & apply chenges
		dicesResult := turn.rollDice()
		fmt.Println("You got:", dicesResult)
		for _, p := range g.players {
			fmt.Println(fmt.Sprintf("player %v, have %v tokens", p.name, p.tokens))
		}


		// check if any one won the game
		// exit
		winner := g.finished()
		if winner != nil {
			fmt.Println("\nWinner: ", winner.name)
			return
		}

		// update turn
		turn = turn.right
	}
}

// game
type game struct {
	players []*player
}

// join add new player to the game
func (g *game) join(playerName string) *player {
	// create new player
	p := player{
		name:  playerName,
		tokens: 3,
	}

	playersCount := len(g.players)
	if playersCount > 0 {
		lastplayer := g.players[playersCount-1]
		p.left = lastplayer
		p.right = g.players[0]
		lastplayer.right = &p
		g.players[0].left = &p
	}

	g.players = append(g.players, &p)

	return &p
}
// finished check do we have only one player with Tokens
// this should be done after every turn
func (g *game) finished() (p *player) {
	playersWithTokens := 0
	for _, value := range g.players {
		if value.tokens > 0 {
			playersWithTokens++
			if playersWithTokens > 1 {
				return nil
			}
			p = value
		}
	}
	return
}
// new init the game
func new() *game {
	return &game{}
}

// player
type player struct{
	name string
	tokens int
	right *player
	left *player
}
// Roll Uses rand to return DiceFace
func (p *player) rollDice() (result []string) {
	// find out how many dices we should roll
	// if user have more than 3 tokens he can only roll 3 dices
	// or he roll exact number of tokens as dices
	dices := p.tokens
	if p.tokens > 2 {
		dices = 3
	}

	// roll the dices and update the value on the playes
	for index := 0; index < dices; index++ {
		d := dice{}
		diceResult := d.roll()
		result = append(result, diceResult)
		switch diceResult {
		case "right":
			p.tokens--
			p.right.tokens++
		case "left":
			p.tokens--
			p.left.tokens++
		case "center":
			p.tokens--
		}
	}

	return
}
// dice
type dice struct {
}

func (d *dice) roll() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(6)
	switch r {
	case 0:
		return "right"
	case 1:
		return "left"
	case 2:
		return "center"
	default:
		return "DoNothing"
	}
}
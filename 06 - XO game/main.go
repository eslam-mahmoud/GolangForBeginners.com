// XO game, practice app using only fmt, if/else, for loops, 2D arrays
package main

import (
	"fmt"
)

func main() {
	// Init the game board with empty strings
	xoBoard := [3][3]string{}
	// var to carry the current player name
	player := "X"

	// infint loop to take user input
	for {
		fmt.Println("\nPlayer:", player)

		// Read row value
		var row int
		fmt.Println("Please enter row number 0,1 or 2")
		fmt.Scanln(&row) // read the input from the user stdin
		if row < 0 || row > 2 {
			fmt.Println("Invalid entry, Please enter row number 0,1 or 2")
			continue
		}

		// Read column value
		var column int
		fmt.Println("Please enter column number 0,1 or 2")
		fmt.Scanln(&column)
		if column < 0 || column > 2 {
			fmt.Println("Invalid entry, Please enter column number 0,1 or 2")
			continue
		}

		// Set value into game board
		if xoBoard[row][column] == "" {
			xoBoard[row][column] = player
		} else {
			// index already have value
			fmt.Println("Invalid entry:", row, column, "value:", xoBoard[row][column])
			// breack and start same player turn
			continue
		}

		// display the game board after each move
		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		// Did someone win
		win := false
		for i := 0; i < 3; i++ {
			// check rows
			if xoBoard[i][0] == xoBoard[i][1] && xoBoard[i][1] == xoBoard[i][2] && xoBoard[i][2] != "" {
				win = true
				break
			}
			// check columns
			if xoBoard[0][i] == xoBoard[1][i] && xoBoard[1][i] == xoBoard[2][i] && xoBoard[2][i] != "" {
				win = true
				break
			}
		}
		if xoBoard[0][0] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][2] && xoBoard[2][2] != "" {
			win = true
		}
		if xoBoard[0][2] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][0] && xoBoard[2][0] != "" {
			win = true
		}
		if win {
			fmt.Println("Game ended, winner is player:", player)
			// end the game and exit the app
			// we can use "break" or "return"
			break
		}

		//Swap players
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}

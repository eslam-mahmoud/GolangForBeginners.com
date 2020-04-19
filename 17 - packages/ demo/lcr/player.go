package lcr

// Player struct
type Player struct {
	name   string
	tokens int
	right  *Player
	left   *Player
}

// Name is a getter for Player.name
func (p *Player) Name() string {
	return p.name
}

// Tokens is a getter for Player.tokens
func (p *Player) Tokens() int {
	return p.tokens
}

// RollDice Uses rand to return DiceFace
func (p *Player) RollDice() (result []string) {
	// find out how many dices we should roll
	// if user have more than 3 tokens he can only roll 3 dices
	// or he roll exact number of tokens as dices
	dices := p.tokens
	if p.tokens > 2 {
		dices = 3
	}

	// roll the dices and update the value on the playes
	for index := 0; index < dices; index++ {
		d := Dice{}
		diceResult := d.Roll()
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

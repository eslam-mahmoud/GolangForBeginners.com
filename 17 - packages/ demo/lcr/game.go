package lcr

// Game struct will represent the game
// should init the game using NewGame()
type Game struct {
	players []*Player
	turn    *Player
}

// Join add new player to the Game
func (g *Game) Join(playerName string) *Player {
	// create new player
	p := Player{
		name:   playerName,
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

// Finished check do we have only one player with Tokens
// this should be done after every turn
func (g *Game) Finished() (p *Player) {
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

// Players is a getter for Game.players
func (g *Game) Players() []*Player {
	return g.players
}

// Turn will return the player that should play this turn
func (g *Game) Turn() *Player {
	if g.turn == nil {
		g.turn = g.players[0]
	} else {
		g.turn = g.turn.right
	}

	return g.turn
}

// NewGame init the Game
func NewGame() *Game {
	return &Game{}
}

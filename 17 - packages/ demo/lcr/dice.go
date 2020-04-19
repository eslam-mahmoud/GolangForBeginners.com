package lcr

import (
	"math/rand"
	"time"
)

// Dice dice struct
type Dice struct {
}

// Roll will return string for dice face
func (d *Dice) Roll() string {
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

package game

import (
	"fmt"
	"os"
)

// End is the end of a game.
type End struct{}

// Render ends the game.
func (End) Render() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	os.Exit(0)
}

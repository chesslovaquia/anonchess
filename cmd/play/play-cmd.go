// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"

	"github.com/notnil/chess"
)

func main() {
	game := chess.NewGame()
	fmt.Println(game.Position().Board().Draw())
}

// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"

	"chesslovaquia/anonchess/lib/game"
)

func main() {
	fmt.Println("Position:", game.Position())
	fmt.Println("   Board:", game.Board())
	fmt.Println(game.BoardDraw())
}

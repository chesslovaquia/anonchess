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

	fmt.Println(game.Status())
	fmt.Println(game.BoardDraw())
	fmt.Println(game.Dump())

	fmt.Println(game.Status())
	game.Move("a2a4")
	fmt.Println(game.BoardDraw())
	fmt.Println(game.Dump())

	fmt.Println(game.Status())
	game.Move("h7h5")
	fmt.Println(game.BoardDraw())
	fmt.Println(game.Dump())
}

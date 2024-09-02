// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"runtime"
	"strings"
	"syscall/js"

	"chesslovaquia/anonchess/lib/game"
)

func main() {
	//
	// anonc_board
	//
	anonc_board := js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("anonc_board:", game.Board())
		return game.BoardMap()
	})
	defer anonc_board.Release()
	js.Global().Set("anonc_board", anonc_board)

	//
	// anonc_valid_move
	//
	anonc_valid_move := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("anonc_valid_move: no args")
			return false
		}
		move := strings.TrimSpace(args[0].String())
		fmt.Println("anonc_valid_move: check", move)
		for _, m := range game.ValidMoves() {
			if m == move {
				return true
			}
		}
		return false
	})
	defer anonc_valid_move.Release()
	js.Global().Set("anonc_valid_move", anonc_valid_move)

	//
	// main loop
	//
	fmt.Println("anonchess-play.wasm:", runtime.Version())
	select {}
}

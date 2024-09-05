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
		return game.Board()
	})
	defer anonc_board.Release()
	js.Global().Set("anonc_board", anonc_board)

	//
	// anonc_board_map
	//
	anonc_board_map := js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("anonc_board_map:", game.Board())
		return game.BoardMap()
	})
	defer anonc_board_map.Release()
	js.Global().Set("anonc_board_map", anonc_board_map)

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
	// anonc_move
	//
	anonc_move := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("anonc_move: no args")
			return false
		}
		m := strings.TrimSpace(args[0].String())
		fmt.Println("anonc_move:", m)
		if err := game.Move(m); err != nil {
			fmt.Println("[ERROR] anonc_move:", err)
			return false
		}
		return true
	})
	defer anonc_move.Release()
	js.Global().Set("anonc_move", anonc_move)

	//
	// anonc_move_tag
	//
	anonc_move_tag := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("anonc_move_tag: no args")
			return false
		}
		m := strings.TrimSpace(args[0].String())
		fmt.Println("anonc_move_tag:", m)
		return game.MoveTag(m)
	})
	defer anonc_move_tag.Release()
	js.Global().Set("anonc_move_tag", anonc_move_tag)

	//
	// anonc_game_dump
	//
	anonc_game_dump := js.FuncOf(func(this js.Value, args []js.Value) any {
		return game.Dump()
	})
	defer anonc_game_dump.Release()
	js.Global().Set("anonc_game_dump", anonc_game_dump)

	//
	// anonc_game_load
	//
	anonc_game_load := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("anonc_game_load: no args")
			return nil
		}
		g := strings.TrimSpace(args[0].String())
		fmt.Println("anonc_game_load:", len(g))
		if err := game.Load(g); err != nil {
			fmt.Println("[ERROR] anonc_game_load:", err)
			return nil
		}
		return game.BoardMap()
	})
	defer anonc_game_load.Release()
	js.Global().Set("anonc_game_load", anonc_game_load)

	//
	// anonc_status
	//
	anonc_status := js.FuncOf(func(this js.Value, args []js.Value) any {
		return game.Status()
	})
	defer anonc_status.Release()
	js.Global().Set("anonc_status", anonc_status)

	//
	// anonc_outcome
	//
	anonc_outcome := js.FuncOf(func(this js.Value, args []js.Value) any {
		return game.Outcome()
	})
	defer anonc_outcome.Release()
	js.Global().Set("anonc_outcome", anonc_outcome)

	//
	// anonc_turn
	//
	anonc_turn := js.FuncOf(func(this js.Value, args []js.Value) any {
		return game.Turn()
	})
	defer anonc_turn.Release()
	js.Global().Set("anonc_turn", anonc_turn)

	//
	// anonc_turn_name
	//
	anonc_turn_name := js.FuncOf(func(this js.Value, args []js.Value) any {
		return game.TurnName()
	})
	defer anonc_turn_name.Release()
	js.Global().Set("anonc_turn_name", anonc_turn_name)

	//
	// main loop
	//
	fmt.Println("anonchess-play.wasm:", runtime.Version())
	select {}
}

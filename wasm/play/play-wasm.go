// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"runtime"
	"strings"
	"syscall/js"

	"chesslovaquia/anonchess/lib/game"
	"chesslovaquia/anonchess/wasm"
)

var (
	AnoncMove          js.Func
	AnoncMoveFailed    js.Func
	AnoncIsEnPassant   js.Func
	AnoncEnPassantTake js.Func
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
	// anonc_position
	//
	anonc_position := js.FuncOf(func(this js.Value, args []js.Value) any {
		return game.Position()
	})
	defer anonc_position.Release()
	js.Global().Set("anonc_position", anonc_position)

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
	// anonc_move
	//
	AnoncMove = wasm.Move()
	defer AnoncMove.Release()
	js.Global().Set("anonc_move", AnoncMove)

	//
	// anonc_move_failed
	//
	AnoncMoveFailed = wasm.MoveFailed()
	defer AnoncMoveFailed.Release()
	js.Global().Set("anonc_move_failed", AnoncMoveFailed)

	//
	// anonc_is_enpassant
	//
	AnoncIsEnPassant = wasm.IsEnPassant()
	defer AnoncIsEnPassant.Release()
	js.Global().Set("anonc_is_enpassant", AnoncIsEnPassant)

	//
	// anonc_enpassant_take
	//
	AnoncEnPassantTake = wasm.EnPassantTake()
	defer AnoncEnPassantTake.Release()
	js.Global().Set("anonc_enpassant_take", AnoncEnPassantTake)

	//
	// main loop
	//
	fmt.Println("anonchess-play.wasm:", runtime.Version())
	select {}
}

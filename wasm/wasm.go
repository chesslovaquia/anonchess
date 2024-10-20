// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package wasm

import (
	"fmt"
	"strings"
	"syscall/js"

	"chesslovaquia/anonchess/lib/game"
	"chesslovaquia/anonchess/lib/moves"
)

func Move() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("ERROR wasm.Move: no args")
			return moves.Error
		}
		m := strings.TrimSpace(args[0].String())
		fmt.Println("wasm.Move:", m)
		return game.Move(m)
	})
}

func MoveFailed() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("ERROR wasm.MoveFailed: no args")
			return true
		}
		m := strings.TrimSpace(args[0].String())
		fmt.Println("wasm.MoveFailed:", m)
		return moves.Failed(m)
	})
}

func IsEnPassant() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("ERROR wasm IsEnPassant: no args")
			return ""
		}
		m := strings.TrimSpace(args[0].String())
		return moves.IsEnPassant(m)
	})
}

func EnPassantTake() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			fmt.Println("ERROR wasm IsEnPassant: no args")
			return ""
		}
		m := strings.TrimSpace(args[0].String())
		return moves.EnPassantTake(m)
	})
}

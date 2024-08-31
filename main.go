// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"runtime"
	"syscall/js"
)

func main() {
	// anonc_board
	anonc_board := js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("anonc_board")
		return []interface{}{
			"rnbqkbnr",
			"pppppppp",
			"        ",
			"        ",
			"        ",
			"        ",
			"PPPPPPPP",
			"RNBQKBNR",
		}
	})
	defer anonc_board.Release()
	js.Global().Set("anonc_board", anonc_board)

	// main loop
	fmt.Println("anonchess.wasm:", runtime.Version())
	select {}
}

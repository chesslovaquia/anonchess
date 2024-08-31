// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"runtime"
	"syscall/js"
)

func main() {

	// anonc_init
	anonc_init := js.FuncOf(func(this js.Value, args []js.Value) any {
		v := runtime.Version()
		fmt.Println("anonc_init:", v)
		return v
	})
	defer anonc_init.Release()
	js.Global().Set("anonc_init", anonc_init)

	fmt.Println("anonchess.wasm: main")
	select {}
}

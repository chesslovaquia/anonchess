// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// main loop
	fmt.Println("anonchess-home.wasm:", runtime.Version())
	select {}
}

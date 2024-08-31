// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// main loop
	fmt.Println("anonchess.wasm:", runtime.Version())
	select {}
}

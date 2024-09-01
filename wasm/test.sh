#!/bin/sh
set -eu
export GOOS=js
export GOARCH=wasm
exec go test -exec="`go env GOROOT`/misc/wasm/go_js_wasm_exec" "$@"

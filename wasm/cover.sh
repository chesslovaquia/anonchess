#!/bin/sh
set -eu
export GOOS=js
export GOARCH=wasm
mkdir -p htmlcov/wasm
go test -coverprofile htmlcov/wasm/coverage.out \
	-exec="`go env GOROOT`/misc/wasm/go_js_wasm_exec" "$@"
exec go tool cover -html htmlcov/wasm/coverage.out -o htmlcov/wasm/index.html

#!/bin/sh
set -eu
export GOOS=js
export GOARCH=wasm
exec go test "$@"

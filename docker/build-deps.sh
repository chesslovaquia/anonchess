#!/bin/sh
set -eu
rustup target add wasm32-unknown-unknown
cargo install wasm-pack
exit 0

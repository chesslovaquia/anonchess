#!/bin/sh
set -eu

_UID=$(id -u)
_GID=$(id -g)

install -d -m 0750 ./docker/build
install -v -t ./docker/build ./go.mod ./go.sum

exec docker build --rm \
	--build-arg "DEVEL_UID=${_UID}" \
	--build-arg "DEVEL_GID=${_GID}" \
	-t chesslovaquia/anonchess ./docker

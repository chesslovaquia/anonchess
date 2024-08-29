#!/bin/sh
set -eu

_UID=$(id -u)
_GID=$(id -g)

rm -rf ./docker/build
install -v -m 0750 -d ./docker/build
install -v -m 0644 -t ./docker/build ./package.json ./package-lock.json

exec docker build --rm \
	--build-arg "DEVEL_UID=${_UID}" \
	--build-arg "DEVEL_GID=${_GID}" \
	-t chesslovaquia/anonchess ./docker

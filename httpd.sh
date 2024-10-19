#!/bin/sh
set -eu
exec python3 -m http.server "$@"

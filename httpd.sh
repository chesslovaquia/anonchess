#!/bin/sh
set -eu
exec python3 -m http.server -b 127.0.0.1

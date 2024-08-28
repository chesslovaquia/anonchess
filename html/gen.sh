#!/bin/sh
set -eu

page=${1:?'page?'}

test -s ./html/${page} || {
	echo "./html/${page}: file not found" >&2
	exit 1
}

cat ./html/head.html >./static/${page}
cat ./html/${page}   >>./static/${page}
cat ./html/tail.html >>./static/${page}
echo "./static/${page} done!"

exit 0

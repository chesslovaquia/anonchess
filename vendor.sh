#!/bin/sh
set -eu

lila_raw=https://github.com/lichess-org/lila/raw/master

mkdir -vp ./static/lila/public/piece/merida

for x in $(echo 'b w'); do
	for y in $(echo 'B K N P Q R'); do
		n="${x}${y}"
		src="${lila_raw}/public/piece/merida/${n}.svg"
		dst="./static/lila/public/piece/merida/${n}.svg"
		echo "${src}"
		echo "  ${dst}"
		wget -q -c -O "${dst}" "${src}"
	done
done

exit 0

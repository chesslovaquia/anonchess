#!/bin/sh
set -eu

lila_raw=https://github.com/lichess-org/lila/raw/master


for t in cburnett merida; do
	mkdir -vp "./static/lila/public/piece/${t}"
	for x in b w; do
		for y in B K N P Q R; do
			n="${x}${y}"
			src="${lila_raw}/public/piece/${t}/${n}.svg"
			dst="./static/lila/public/piece/${t}/${n}.svg"
			echo "${src}"
			echo "  ${dst}"
			##wget -q -c -O "${dst}" "${src}"
		done
	done
done

exit 0

#!/bin/sh
set -eu

lila_raw=https://github.com/lichess-org/lila/raw/master

mkdir -vp ./static/lila

for f in COPYING.md LICENSE README.md; do
	src="${lila_raw}/${f}"
	dst="./static/lila/${f}"
	echo "${src}"
	echo "  ${dst}"
	wget -q -c -O "${dst}" "${src}"
done

mkdir -vp ./static/lila/public/piece
mkdir -vp ./static/lila/public/piece-css

for t in cburnett merida mpchess; do

	# svg
	mkdir -vp "./static/lila/public/piece/${t}"
	for x in b w; do
		for y in B K N P Q R; do
			n="${x}${y}"
			src="${lila_raw}/public/piece/${t}/${n}.svg"
			dst="./static/lila/public/piece/${t}/${n}.svg"
			echo "${src}"
			echo "  ${dst}"
			wget -q -c -O "${dst}" "${src}"
		done
	done

	# css
	for x in b w; do
		for y in B K N P Q R; do
			src="${lila_raw}/public/piece-css/${t}.css"
			dst="./static/lila/public/piece-css/${t}.css"
			echo "${src}"
			echo "  ${dst}"
			wget -q -c -O "${dst}" "${src}"
		done
	done

done

exit 0

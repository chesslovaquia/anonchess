#!/bin/sh
set -eu

lila_raw=https://github.com/lichess-org/lila/raw/master

mkdir -vp ./static/lila

#
# license
#
for f in COPYING.md LICENSE README.md; do
	src="${lila_raw}/${f}"
	dst="./static/lila/${f}"
	echo "${src}"
	echo "  ${dst}"
	wget -q -c -O "${dst}" "${src}"
done

#
# board
#
mkdir -vp ./static/lila/public/images/board
src="${lila_raw}/public/images/board/wood4.jpg"
dst="./static/lila/public/images/board/wood4.jpg"
echo "${src}"
echo "  ${dst}"
wget -q -c -O "${dst}" "${src}"

#
# pieces
#
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
	src="${lila_raw}/public/piece-css/${t}.css"
	dst="./static/lila/public/piece-css/${t}.css"
	echo "${src}"
	echo "  ${dst}"
	wget -q -c -O "${dst}" "${src}"

done

exit 0

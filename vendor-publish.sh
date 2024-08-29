#!/bin/sh
set -eu

install -v -d -m 0755 ./publish/lila

#~ for f in COPYING.md LICENSE README.md; do
	#~ src="./static/lila/${f}"
	#~ dst="./publish/lila/${f}"
	#~ install -v -m 0644 "${src}" "${dst}"
#~ done

install -v -d -m 0755 ./publish/lila/public
install -v -d -m 0755 ./publish/lila/public/piece
install -v -d -m 0755 ./publish/lila/public/piece-css

for t in cburnett merida mpchess; do

	install -v -d -m 0755 "./publish/lila/public/piece/${t}"

	# svg
	for x in b w; do
		for y in B K N P Q R; do
			n="${x}${y}"
			src="./static/lila/public/piece/${t}/${n}.svg"
			dst="./publish/lila/public/piece/${t}/${n}.svg"
			install -v -m 0644 "${src}" "${dst}"
		done
	done

	# css
	src="./static/lila/public/piece-css/${t}.css"
	dst="./publish/lila/public/piece-css/${t}.css"
	install -v -m 0644 "${src}" "${dst}"

done

exit 0

#!/bin/sh
set -eu

install -v -d -m 0755 ./publish/lila
install -v -d -m 0755 ./publish/lila/public

#
# license
#
#~ for f in COPYING.md LICENSE README.md; do
	#~ src="./static/lila/${f}"
	#~ dst="./publish/lila/${f}"
	#~ install -v -m 0644 "${src}" "${dst}"
#~ done

#
# board
#
install -v -m 0755 -d ./publish/lila/public/images
install -v -m 0755 -d ./publish/lila/public/images/board
install -v -m 0644 -t ./publish/lila/public/images/board ./static/lila/public/images/board/wood4.jpg

#
# pieces
#
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

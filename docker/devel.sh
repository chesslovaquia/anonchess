#!/bin/sh
set -eu
exec docker run -it --rm -u devel \
	--name anonchess \
	--hostname anonchess.local \
	-v ${PWD}:/home/devel/anonchess \
	--workdir /home/devel/anonchess \
	chesslovaquia/anonchess

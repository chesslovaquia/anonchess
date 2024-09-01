// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package game

import (
	"github.com/notnil/chess"
)

var g *chess.Game

func init() {
	g = chess.NewGame()
}

func Position() string {
	return g.Position().String()
}

func Board() string {
	return g.Position().Board().String()
}

func BoardDraw() string {
	return g.Position().Board().Draw()
}

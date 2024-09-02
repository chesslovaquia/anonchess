// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package game

import (
	"unicode"

	"github.com/notnil/chess"
)

var g *chess.Game

func init() {
	g = chess.NewGame()
}

func New(fen string) error {
	g = nil
	s, err := chess.FEN(fen)
	if err != nil {
		return err
	}
	g = chess.NewGame(s)
	return nil
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

func BoardMap() []any {
	var m []any
	var row string
	for _, square := range Board() {
		if unicode.IsDigit(square) {
			count := int(square - '0')
			for i := 0; i < count; i++ {
				row += " "
			}
			continue
		}
		if square == '/' {
			m = append(m, row)
			row = ""
			continue
		}
		row += string(square)
	}
	m = append(m, row)
	return m
}

func ValidMoves() []string {
	vm := make([]string, 0)
	for _, m := range g.ValidMoves() {
		vm = append(vm, m.String())
	}
	return vm
}

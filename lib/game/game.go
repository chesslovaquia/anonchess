// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package game

import (
	"fmt"
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

func Move(m string) error {
	move, err := chess.UCINotation{}.Decode(g.Position(), m)
	if err != nil {
		return err
	} else {
		fmt.Println("move:", move)
	}
	return g.Move(move)
}

func MoveTag(m string) string {
	move, err := chess.UCINotation{}.Decode(g.Position(), m)
	if err != nil {
		fmt.Printf("ERROR MoveTag: %s - %s\n", m, err)
		return "error"
	}
	if move.HasTag(chess.KingSideCastle) {
		return "KingSideCastle"
	}
	if move.HasTag(chess.QueenSideCastle) {
		return "QueenSideCastle"
	}
	if move.HasTag(chess.Capture) {
		return "Capture"
	}
	if move.HasTag(chess.EnPassant) {
		return "EnPassant"
	}
	if move.HasTag(chess.Check) {
		return "Check"
	}
	return ""
}

func Dump() string {
	return g.Position().String()
}

func Load(s string) error {
	return New(s)
}

func Status() string {
	s := g.Position().Status()
	if s == chess.NoMethod {
		return "Playing"
	}
	return s.String()
}

func Outcome() string {
	return g.Outcome().String()
}

func Turn() string {
	return g.Position().Turn().String()
}

func TurnName() string {
	return g.Position().Turn().Name()
}

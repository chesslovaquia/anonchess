// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package game

import (
	"testing"
)

func boardMapCheck(t *testing.T, expect []any) {
	t.Helper()
	board := BoardMap()
	for i := 0; i < 8; i++ {
		if board[i] != expect[i] {
			t.Errorf("row %d - expect '%s' - got '%s'", i+1, expect[i], board[i])
		}
	}
}

func TestBoardMap(t *testing.T) {

	New("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	boardMapCheck(t, []any{
		"rnbqkbnr",
		"pppppppp",
		"        ",
		"        ",
		"        ",
		"        ",
		"PPPPPPPP",
		"RNBQKBNR",
	})

	New("k1K5/8/8/8/8/8/8/1Q6 w - - 0 1")
	boardMapCheck(t, []any{
		"k K     ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		" Q      ",
	})

	New("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 100 23")
	boardMapCheck(t, []any{
		"  r   k ",
		" q nbppp",
		"r   p   ",
		"   pP   ",
		"pPpP    ",
		"P Q  N  ",
		"  RN PPP",
		"  R    K",
	})

	New("8/2k5/8/8/8/3K4/8/8 w - - 1 1")
	boardMapCheck(t, []any{
		"        ",
		"  k     ",
		"        ",
		"        ",
		"        ",
		"   K    ",
		"        ",
		"        ",
	})
}

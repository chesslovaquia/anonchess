// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package moves

import (
	"testing"
)

func checkEnPassant(t *testing.T, move, expect string) {
	t.Helper()
	got := CheckEnPassant(EnPassant, move)
	if (got != expect) {
		t.Errorf("CheckEnPassant %s: expect '%s' - got '%s'", move, expect, got)
	}
}

func TestEnPassantMove(t *testing.T) {
	got := CheckEnPassant(Capture, "a2b3")
	if got != Move {
		t.Errorf("CheckEnPassant a2b3: expect '%s' - got '%s'", Move, got)
	}
}

func TestEnPassantError(t *testing.T) {
	checkEnPassant(t, "none", Error)
	checkEnPassant(t, "a2a3", Error)
}

func TestEnPassantCapture(t *testing.T) {
	// white captures en passant
	// a5
	checkEnPassant(t, "a5b6", "b5")
	// b5
	checkEnPassant(t, "b5a6", "a5")
	checkEnPassant(t, "b5c6", "c5")
	// c5
	checkEnPassant(t, "c5b6", "b5")
	checkEnPassant(t, "c5d6", "d5")
	// d5
	checkEnPassant(t, "d5c6", "c5")
	checkEnPassant(t, "d5e6", "e5")
	// e5
	checkEnPassant(t, "e5d6", "d5")
	checkEnPassant(t, "e5f6", "f5")
	// f5
	checkEnPassant(t, "f5e6", "e5")
	checkEnPassant(t, "f5g6", "g5")
	// g5
	checkEnPassant(t, "g5f6", "f5")
	checkEnPassant(t, "g5h6", "h5")
	// h5
	checkEnPassant(t, "h5g6", "g5")
	// black captures en passant
	// a4
	checkEnPassant(t, "a4b3", "b4")
	// b4
	checkEnPassant(t, "b4a3", "a4")
	checkEnPassant(t, "b4c3", "c4")
	// c4
	checkEnPassant(t, "c4b3", "b4")
	checkEnPassant(t, "c4d3", "d4")
	// d4
	checkEnPassant(t, "d4c3", "c4")
	checkEnPassant(t, "d4e3", "e4")
	// e4
	checkEnPassant(t, "e4d3", "d4")
	checkEnPassant(t, "e4f3", "f4")
	// f4
	checkEnPassant(t, "f4e3", "e4")
	checkEnPassant(t, "f4g3", "g4")
	// g4
	checkEnPassant(t, "g4f3", "f4")
	checkEnPassant(t, "g4h3", "h4")
	// h4
	checkEnPassant(t, "h4g3", "g4")
}

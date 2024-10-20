// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package moves

import (
	"testing"
)

func ept(t *testing.T, move, expect string) {
	t.Helper()
	got := EnPassantTake(move)
	if (got != expect) {
		t.Errorf("EnPassantTake %s: expect '%s' - got '%s'", move, expect, got)
	}
}

func TestEnPassantTakeError(t *testing.T) {
	ept(t, "none", Error)
	ept(t, "a2a3", Error)
}

func TestEnPassantTake(t *testing.T) {
	// white captures en passant
	// a5
	ept(t, "a5b6", "b5")
	// b5
	ept(t, "b5a6", "a5")
	ept(t, "b5c6", "c5")
	// c5
	ept(t, "c5b6", "b5")
	ept(t, "c5d6", "d5")
	// d5
	ept(t, "d5c6", "c5")
	ept(t, "d5e6", "e5")
	// e5
	ept(t, "e5d6", "d5")
	ept(t, "e5f6", "f5")
	// f5
	ept(t, "f5e6", "e5")
	ept(t, "f5g6", "g5")
	// g5
	ept(t, "g5f6", "f5")
	ept(t, "g5h6", "h5")
	// h5
	ept(t, "h5g6", "g5")
	// black captures en passant
	// a4
	ept(t, "a4b3", "b4")
	// b4
	ept(t, "b4a3", "a4")
	ept(t, "b4c3", "c4")
	// c4
	ept(t, "c4b3", "b4")
	ept(t, "c4d3", "d4")
	// d4
	ept(t, "d4c3", "c4")
	ept(t, "d4e3", "e4")
	// e4
	ept(t, "e4d3", "d4")
	ept(t, "e4f3", "f4")
	// f4
	ept(t, "f4e3", "e4")
	ept(t, "f4g3", "g4")
	// g4
	ept(t, "g4f3", "f4")
	ept(t, "g4h3", "h4")
	// h4
	ept(t, "h4g3", "g4")
}

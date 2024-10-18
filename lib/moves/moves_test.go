// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package moves

import (
	"testing"
)

func TestEnPassant(t *testing.T) {
	if EnPassant("none", "none") != "" {
		t.Error("EnPassant: not none")
	}
}

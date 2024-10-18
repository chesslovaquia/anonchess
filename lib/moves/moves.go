// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package moves

const (
	Error = "error"
	KingSideCastle = "KingSideCastle"
	QueenSideCastle = "QueenSideCastle"
	EnPassant = "EnPassant"
	Check = "Check"
	Capture = "Capture"
	Move = ""
)

func CheckEnPassant(tag, move string) string {
	if tag == EnPassant {
		if move == "e5d6" {
			return "d5"
		}
		return "error"
	}
	return ""
}

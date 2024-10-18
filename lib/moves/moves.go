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
	// avoid calculations
	if tag == EnPassant {
		// white captures en passant
		if move == "b5a6" {
			return "a5"
		}
		if move == "c5b6" {
			return "b5"
		}
		if move == "d5c6" {
			return "c5"
		}
		if move == "e5d6" {
			return "d5"
		}
		if move == "f5e6" {
			return "e5"
		}
		if move == "g5f6" {
			return "f5"
		}
		if move == "h5g6" {
			return "g5"
		}
		if move == "e5d6" {
			return "d5"
		}
		// black captures en passant
		if move == "a4b3" {
			return "b4"
		}
		if move == "b4c3" {
			return "c4"
		}
		if move == "c4d3" {
			return "d4"
		}
		if move == "d4e3" {
			return "e4"
		}
		if move == "e4f3" {
			return "f4"
		}
		if move == "g4h3" {
			return "h4"
		}
		return Error
	}
	return Move
}

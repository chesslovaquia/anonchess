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
		// a5
		if move == "a5b6" {
			return "b5"
		}
		// b5
		if move == "b5a6" {
			return "a5"
		}
		if move == "b5c6" {
			return "c5"
		}
		// c5
		if move == "c5b6" {
			return "b5"
		}
		if move == "c5d6" {
			return "d5"
		}
		// d5
		if move == "d5c6" {
			return "c5"
		}
		if move == "d5e6" {
			return "e5"
		}
		// e5
		if move == "e5d6" {
			return "d5"
		}
		if move == "e5f6" {
			return "f5"
		}
		// f5
		if move == "f5e6" {
			return "e5"
		}
		if move == "f5g6" {
			return "g5"
		}
		// g5
		if move == "g5f6" {
			return "f5"
		}
		if move == "g5h6" {
			return "h5"
		}
		// h5
		if move == "h5g6" {
			return "g5"
		}
		// black captures en passant
		// a4
		if move == "a4b3" {
			return "b4"
		}
		// b4
		if move == "b4a3" {
			return "a4"
		}
		if move == "b4c3" {
			return "c4"
		}
		// c4
		if move == "c4b3" {
			return "b4"
		}
		if move == "c4d3" {
			return "d4"
		}
		// d4
		if move == "d4c3" {
			return "c4"
		}
		if move == "d4e3" {
			return "e4"
		}
		// e4
		if move == "e4d3" {
			return "d4"
		}
		if move == "e4f3" {
			return "f4"
		}
		// f4
		if move == "f4e3" {
			return "e4"
		}
		if move == "f4g3" {
			return "g4"
		}
		// g4
		if move == "g4f3" {
			return "f4"
		}
		if move == "g4h3" {
			return "h4"
		}
		// h4
		if move == "h4g3" {
			return "g4"
		}
		return Error
	}
	return Move
}

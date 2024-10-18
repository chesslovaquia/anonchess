// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package moves

func EnPassant(tag, move string) string {
	if tag == "EnPassant" {
		if move == "e5d6" {
			return "d5"
		}
		return "error"
	}
	return ""
}

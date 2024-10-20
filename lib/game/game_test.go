// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package game

import (
	"testing"

	"hash/fnv"
	"strings"

	"chesslovaquia/anonchess/lib/moves"
)

const pos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
const board = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

func initBoard(t *testing.T) {
	t.Helper()
	err := New(pos)
	if err != nil {
		t.Fatalf("init board error: %s", err)
	}
}

func initPosition(t *testing.T, p string) {
	t.Helper()
	err := New(p)
	if err != nil {
		t.Fatalf("init board error: %s", err)
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func TestNewError(t *testing.T) {
	err := New(pos)
	if err != nil {
		t.Errorf("New: %s", err)
	}
	err = New("INVALID_FEN")
	if err == nil {
		t.Error("New: expected INVALID_FEN error")
	}
}

func TestPosition(t *testing.T) {
	initBoard(t)
	p := Position()
	if p != pos {
		t.Errorf("Position: got '%s' - expect '%s'", p, pos)
	}
}

func TestBoard(t *testing.T) {
	initBoard(t)
	b := Board()
	if b != board {
		t.Errorf("Board: got '%s' - expect '%s'", b, board)
	}
}

func TestBoardDraw(t *testing.T) {
	initBoard(t)
	bd := BoardDraw()
	expect := uint32(2355033525)
	h := hash(bd)
	if h != expect {
		t.Logf("BoardDraw: %s", bd)
		t.Errorf("BoardDraw hash: got '%d' - expect '%d'", h, expect)
	}
}

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

func TestValidMoves(t *testing.T) {
	initBoard(t)
	vm := ValidMoves()
	got := strings.Join(vm, " ")
	expect := "b1a3 b1c3 g1f3 g1h3 a2a3 a2a4 b2b3 b2b4 c2c3 c2c4 d2d3 d2d4 e2e3 e2e4 f2f3 f2f4 g2g3 g2g4 h2h3 h2h4"
	if got != expect {
		t.Errorf("ValidMoves Error\ngot:    '%s'\nexpect: '%s'", got, expect)
	}
}

func TestMoveInvalidUCI(t *testing.T) {
	initBoard(t)
	if tag := Move("UCI"); tag != moves.Error {
		t.Errorf("Move expected invalid UCI but got %s", tag)
	}
}

func TestMove(t *testing.T) {
	initBoard(t)
	if tag := Move("a2a3"); tag != moves.Move {
		t.Errorf("Move a2a3 tag error: %s", tag)
	}
}

func TestMoveError(t *testing.T) {
	initBoard(t)
	if tag := Move("a2a1"); tag != moves.Error {
		t.Errorf("Move a2a2 expected error tag but got %s", tag)
	}
}

func TestMoveTag(t *testing.T) {
	initBoard(t)
	if tag := Move("a2a3"); tag != moves.Move {
		t.Errorf("MoveTag expected '%s' - got '%s'", moves.Move, tag)
	}
}

func TestMoveTagInvalidUCI(t *testing.T) {
	initBoard(t)
	if tag := Move("UCI"); tag != moves.Error {
		t.Errorf("MoveTag expected '%s' - got '%s'", moves.Error, tag)
	}
}

func TestMoveTagKingSideCastle(t *testing.T) {
	pos := "rnbqkbnr/pp3ppp/2p5/3pp3/4P3/3B1N2/PPPP1PPP/RNBQK2R w KQkq e6 0 4"
	initPosition(t, pos)
	if tag := Move("e1g1"); tag != moves.KingSideCastle {
		t.Errorf("MoveTag expected '%s' - got '%s'", moves.KingSideCastle, tag)
	}
}

func TestMoveTagQueenSideCastle(t *testing.T) {
	pos := "r3kbnr/pp1n1ppp/1qp1b3/3pp3/1P2P3/2PB1N1P/P2P1PP1/RNBQ1RK1 b kq - 0 7"
	initPosition(t, pos)
	if tag := Move("e8c8"); tag != moves.QueenSideCastle {
		t.Errorf("MoveTag expected '%s' - got '%s'", moves.QueenSideCastle, tag)
	}
}

func TestMoveTagEnPassant(t *testing.T) {
	pos := "rnbqkbnr/ppp2ppp/4p3/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3"
	initPosition(t, pos)
	if tag := Move("e5d6"); tag != moves.EnPassant {
		t.Errorf("MoveTag expected '%s' - got '%s'", moves.EnPassant, tag)
	}
}

//~ func TestMoveTagCheck(t *testing.T) {
	//~ pos := "rnbqkbnr/ppp1pppp/3p4/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2"
	//~ initPosition(t, pos)
	//~ if tag := Move("f1b5"); tag != moves.Check {
		//~ t.Errorf("MoveTag expected '%s' - got '%s'", moves.Check, tag)
	//~ }
//~ }

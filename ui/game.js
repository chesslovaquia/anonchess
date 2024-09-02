import React from 'react';

//
// saveGame
//
function saveGame() {
	sessionStorage.setItem('anonc_game', anonc_game_dump());
}

//
// loadGame
//
function loadGame() {
	const g = sessionStorage.getItem('anonc_game');
	if (g) {
		return anonc_game_load(g);
	}
	return null
}

//
// renderBoard
//
export function renderBoard() {
	console.log('renderBoard:', anonc_board());

	let board = loadGame();
	if (!board) {
		board = anonc_board_map();
	}

	const pieceImage = {
		'r': 'bR.svg', 'n': 'bN.svg', 'b': 'bB.svg',
		'q': 'bQ.svg', 'k': 'bK.svg', 'p': 'bP.svg',
		'R': 'wR.svg', 'N': 'wN.svg', 'B': 'wB.svg',
		'Q': 'wQ.svg', 'K': 'wK.svg', 'P': 'wP.svg',
	};

	const pieceAlt = {
		'r': '♜', 'n': '♞', 'b': '♝', 'q': '♛', 'k': '♚', 'p': '♟',
		'R': '♖', 'N': '♘', 'B': '♗', 'Q': '♕', 'K': '♔', 'P': '♙',
	};

	const files = 'abcdefgh';
	const ranks = '12345678';

	const sq = (file, rank) => {
		return `${files[file]}${ranks[7 - rank]}`;
	};

	const sqid = (file, rank) => {
		return `square-${files[file]}${ranks[7 - rank]}`;
	};

	return board.map((row, rowIndex) => (
		<div className="chess-row" key={rowIndex}>
			{row.split('').map((square, colIndex) => (
				<div className="chess-square" key={colIndex} id={sqid(colIndex, rowIndex)} data-square={sq(colIndex, rowIndex)} data-kind="square">
					{square !== ' ' && (
						<img
							data-kind="piece"
							data-square={sq(colIndex, rowIndex)}
							src={`../lila/public/piece/cburnett/${pieceImage[square]}`}
							alt={`${pieceAlt[square]}`}
							className="chess-piece"
						/>
					)}
				</div>
			))}
		</div>
	));
};

//
// handleMove
//
export function handleMove(piece, sq1, sq2, move) {
	console.log('handleMove:', move);
	if (anonc_valid_move(move)) {
		if (anonc_move(move)) {
			saveGame();
			piece.style.top = `${sq2.offsetTop}px`;
			piece.style.left = `${sq2.offsetLeft}px`;
			piece.dataset.square = sq2.dataset.square;
			sq1.innerHTML = '';
			sq2.innerHTML = '';
			sq2.appendChild(piece);
			console.log('handleMove:', move, 'done');
		} else {
			console.log('handleMove:', move, 'failed');
		}
	} else {
		console.log('handleMove:', move, 'invalid');
		sq2.style.border = '2px solid red';
	}
	console.log('handleMove: board', anonc_board());
}

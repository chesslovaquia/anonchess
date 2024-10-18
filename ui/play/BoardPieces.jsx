// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

function loadGame() {
	const g = sessionStorage.getItem('anonc_game');
	if (g) {
		return anonc_game_load(g);
	}
	return null
}

function renderBoard(board) {
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
}

function BoardPieces() {
	console.log('BoardPieces:', anonc_board());

	let board = loadGame();
	if (!board) {
		board = anonc_board_map();
	}

	return (
		<div id="chessboard-pieces" className="pieces-overlay">
			{renderBoard(board)}
		</div>
	);
};

export default BoardPieces;

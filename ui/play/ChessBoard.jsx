// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

import { handleMove } from '../game.js';

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
function renderBoard(renderReg) {
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
// ChessBoard
//
function ChessBoard({renderReg}) {
	const highlight = '2px solid yellow';

	let sq1 = null;
	let sq2 = null;
	let piece = null;

	const handleClick = (event) => {
		const t = event.target;
		console.log('handle click');
		console.log(t.dataset.kind, t.dataset.square, t);
		if (t.dataset.kind === 'piece') {
			if (sq1) {
				sq1.style.border = '';
			}
			if (sq2) {
				sq2.style.border = '';
			}
			sq1 = document.getElementById(`square-${t.dataset.square}`);
			sq1.style.border = highlight;
			piece = t;
		} else if (t.dataset.kind === 'square') {
			if (sq2 && sq2 !== sq1) {
				sq2.style.border = '';
			}
			if (piece) {
				const move = `${piece.dataset.square}${t.dataset.square}`;
				sq2 = t;
				sq2.style.border = highlight;
				handleMove(renderReg, piece, sq1, sq2, move);
				piece = null;
			} else {
				if (sq1 && sq1 !== sq2) {
					sq1.style.border = '';
				}
			}
		}
		if (piece) {
			console.log('piece:', piece.dataset.square);
		}
		if (sq1) {
			console.log('sq1:', sq1.dataset.square);
		}
		if (sq2) {
			console.log('sq2:', sq2.dataset.square);
		}
	};

	return (
		<div id="chessboard" className="chessboard-container" onClick={handleClick}>
			<img src="../lila/public/images/board/wood4.jpg" alt="Chessboard" className="chessboard-image" />
			<div className="pieces-overlay">
				{renderBoard(renderReg)}
			</div>
		</div>
	);
}
export default ChessBoard;

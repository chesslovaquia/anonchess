// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

function handleMove(piece, sq1, sq2, move) {
	console.log('handleMove:', move);
	if (anonc_valid_move(move)) {
		piece.style.top = `${sq2.offsetTop}px`;
		piece.style.left = `${sq2.offsetLeft}px`;
		piece.dataset.square = sq2.dataset.square;
		sq1.innerHTML = '';
		sq2.innerHTML = '';
		sq2.appendChild(piece);
		console.log('handleMove:', move, 'done');
	} else {
		console.log('handleMove:', move, 'invalid');
		sq2.style.border = '2px solid red';
	}
}

function ChessBoard() {

	const board = anonc_board();

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

	const renderBoard = () => {
		const squareName = "";
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
				handleMove(piece, sq1, sq2, move);
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
		<div className="chessboard-container" onClick={handleClick}>
			<img src="../lila/public/images/board/wood4.jpg" alt="Chessboard" className="chessboard-image" />
			<div className="pieces-overlay">
				{renderBoard()}
			</div>
		</div>
	);

}

export default ChessBoard;

// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

function ChessBoard() {

	const board = anonc_board();

	const pieceToFile = {
		'r': 'bR.svg', 'n': 'bN.svg', 'b': 'bB.svg',
		'q': 'bQ.svg', 'k': 'bK.svg', 'p': 'bP.svg',
		'R': 'wR.svg', 'N': 'wN.svg', 'B': 'wB.svg',
		'Q': 'wQ.svg', 'K': 'wK.svg', 'P': 'wP.svg',
	};

	const files = 'abcdefgh';
	const ranks = '12345678';

	const sq = (file, rank) => {
		return `${files[file]}${ranks[7 - rank]}`;
	};

	const sqid = (file, rank) => {
		return `square-${files[file]}${ranks[7 - rank]}`;
	};

	const pid = (file, rank) => {
		return `piece-${files[file]}${ranks[7 - rank]}`;
	};

	const renderBoard = () => {
		const squareName = "";
		return board.map((row, rowIndex) => (
			<div className="chess-row" key={rowIndex}>
				{row.split('').map((square, colIndex) => (
					<div className="chess-square" key={colIndex} id={sqid(colIndex, rowIndex)} data-square={sq(colIndex, rowIndex)} data-kind="square">
						{square !== ' ' && (
							<img
								id={pid(colIndex, rowIndex)}
								data-kind="piece"
								data-square={sq(colIndex, rowIndex)}
								src={`../lila/public/piece/cburnett/${pieceToFile[square]}`}
								alt=""
								className="chess-piece"
							/>
						)}
					</div>
				))}
			</div>
		));
	};

	let square = null;
	let piece = null;

	const handleClick = (event) => {
		const t = event.target;
		console.log(t.id, t);
		//~ console.log(t.attributes);
		//~ console.log(t.tagName);
		//~ console.log(t.className);
		//~ console.log(t.textContent);
		//~ console.log(t.innerHTML);
		//~ console.log(t.dataset.kind, t.dataset.square);
		if (t.dataset.kind === 'piece') {
			piece = t;
		} else if (t.dataset.kind === 'square') {
			if (piece) {
				const move = `${piece.dataset.square}${t.dataset.square}`;
				console.log('move:', move, '???');
				piece = null;
			}
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

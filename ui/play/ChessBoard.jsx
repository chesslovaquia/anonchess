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

	const renderBoard = () => {
		const squareName = "";
		return board.map((row, rowIndex) => (
			<div className="chess-row" key={rowIndex}>
				{row.split('').map((square, colIndex) => (
					<div className="chess-square" key={colIndex} data-square={`${files[colIndex]}${ranks[7 - rowIndex]}`}>
						{square !== ' ' && (
							<img
								data-square={`${files[colIndex]}${ranks[7 - rowIndex]}`}
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

	const handleClick = (event) => {
		const square = event.target;
		//~ console.log(square);
		//~ console.log(square.attributes);
		//~ console.log(square.tagName);
		//~ console.log(square.id);
		//~ console.log(square.className);
		//~ console.log(square.textContent);
		//~ console.log(square.innerHTML);
		console.log(square.dataset.square);
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

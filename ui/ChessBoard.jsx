import React from 'react';

function ChessBoard() {
	const board = [
		"rnbqkbnr",
		"pppppppp",
		"        ",
		"        ",
		"        ",
		"        ",
		"PPPPPPPP",
		"RNBQKBNR",
	];

	const renderBoard = () => {
		return board.map((row, rowIndex) => (
			<div className="w3-row" key={rowIndex}>
				{row.split('').map((square, colIndex) => (
					<div
						className={`w3-col s1 w3-center w3-border ${
							(rowIndex + colIndex) % 2 === 0 ? 'w3-light-grey' : 'w3-dark-grey'
						}`}
						key={colIndex}
					>
						<span className="w3-large">
							{square !== ' ' ? renderPiece(square) : ''}
						</span>
					</div>
				))}
			</div>
		));
	};

	const renderPiece = (piece) => {
		const pieces = {
			'r': '♜', 'n': '♞', 'b': '♝', 'q': '♛', 'k': '♚', 'p': '♟',
			'R': '♖', 'N': '♘', 'B': '♗', 'Q': '♕', 'K': '♔', 'P': '♙',
		};
		return pieces[piece] || '';
	};

	return (
		<div className="w3-margin-top">
			{renderBoard()}
		</div>
	);
}

export default ChessBoard;

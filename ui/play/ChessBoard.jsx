// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

import { renderBoard, handleMove } from '../game.js';

function ChessBoard() {
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
		<div id="chessboard" className="chessboard-container" onClick={handleClick}>
			<img src="../lila/public/images/board/wood4.jpg" alt="Chessboard" className="chessboard-image" />
			<div className="pieces-overlay">
				{renderBoard()}
			</div>
		</div>
	);
}
export default ChessBoard;

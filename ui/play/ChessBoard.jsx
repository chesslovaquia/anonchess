// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';
import BoardPieces from './BoardPieces';

import { handleMove } from '../game.js';

function ChessBoard({renderReg}) {
	const highlight = '2px solid yellow';
	const highlightTake = '2px solid yellow';

	let sq1 = null;
	let sq2 = null;
	let piece = null;

	const doMove = (target) => {
		const move = `${piece.dataset.square}${target.dataset.square}`;
		console.log('doMove:', move);
		if (target.dataset.kind === 'piece') {
			sq2 = document.getElementById(`square-${target.dataset.square}`);
			sq2.style.border = highlightTake;
		} else {
			sq2 = target;
			sq2.style.border = highlight;
		}
		handleMove(renderReg, piece, sq1, sq2, move);
		piece = null;
	};

	const handleClick = (event) => {
		const t = event.target;
		console.log('handle click');
		console.log(t.dataset.kind, t.dataset.square, t.dataset.color, t);
		if (t.dataset.kind === 'piece') {
			if (sq2) {
				sq2.style.border = '';
			}
			if (piece) {
				// take oponent piece
				console.log('handle click: take piece');
				doMove(t);
			} else {
				if (sq1) {
					sq1.style.border = '';
				}
				sq1 = document.getElementById(`square-${t.dataset.square}`);
				sq1.style.border = highlight;
				piece = t;
			}
		} else if (t.dataset.kind === 'square') {
			if (sq2 && sq2 !== sq1) {
				sq2.style.border = '';
			}
			if (piece) {
				// move piece
				console.log('handle click: move piece');
				doMove(t);
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
			<BoardPieces />
		</div>
	);
}
export default ChessBoard;

import React from 'react';

//
// saveGame
//
function saveGame() {
	sessionStorage.setItem('anonc_game', anonc_game_dump());
}

//
// handleMove
//
export function handleMove(renderReg, piece, sq1, sq2, move) {
	console.log('handleMove:', move);
	if (anonc_valid_move(move)) {
		const tag = anonc_move(move)
		if (anonc_move_failed(tag)) {
			console.log('handleMove:', move, 'failed');
		} else {
			saveGame();
			piece.style.top = `${sq2.offsetTop}px`;
			piece.style.left = `${sq2.offsetLeft}px`;
			piece.dataset.square = sq2.dataset.square;
			sq1.innerHTML = '';
			sq2.innerHTML = '';
			sq2.appendChild(piece);
			console.log('handleMove:', tag, 'done');
			renderMove(renderReg, move, tag);
		}
	} else {
		console.log('handleMove:', move, 'invalid');
		sq2.style.border = '2px solid red';
	}
	//~ console.log('handleMove: board', anonc_board());
	console.log('position:', anonc_position());
}

//
// renderMove
//
function renderMove(renderReg, move, tag) {
	console.log('renderMove:', move, tag);
	if (anonc_is_enpassant(tag)) {
		renderEnPassant(move);
	}
	renderReg.current.sideBar();
}

function renderEnPassant(move) {
	console.log('renderEnPassant:', move);
	const sq = anonc_enpassant_take(move);
	console.log('renderEnPassant take:', sq);
	const clean_sq = document.getElementById(`square-${sq}`);
	clean_sq.innerHTML = '';
}

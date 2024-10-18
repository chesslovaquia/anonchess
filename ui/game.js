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
		const tag = anonc_move_tag(move);
		console.log('move tag:', tag);
		if (anonc_move(move)) {
			saveGame();
			piece.style.top = `${sq2.offsetTop}px`;
			piece.style.left = `${sq2.offsetLeft}px`;
			piece.dataset.square = sq2.dataset.square;
			sq1.innerHTML = '';
			sq2.innerHTML = '';
			sq2.appendChild(piece);
			console.log('handleMove:', move, 'done');
			renderMove(renderReg, move, tag);
		} else {
			console.log('handleMove:', move, 'failed');
		}
	} else {
		console.log('handleMove:', move, 'invalid');
		sq2.style.border = '2px solid red';
	}
	console.log('handleMove: board', anonc_board());
}

//
// renderMove
//
function renderMove(renderReg, move, tag) {
	console.log('renderMove:', move, tag);
	const clean_sq = anonc_enpassant(tag, move);
	if (clean_sq !== '') {
		renderEnPassant(clean_sq);
	}
	renderReg.current.sideBar();
}

function renderEnPassant(sq) {
	console.log('renderEnPassant:', sq);
	if (sq !== 'error') {
		const clean_sq = document.getElementById(`square-${sq}`);
		clean_sq.innerHTML = '';
	}
}

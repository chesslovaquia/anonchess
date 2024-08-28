//import init, { start_anonymous_game, make_move } from '../pkg/anonchess.js';
import init from '../pkg/anonchess.js';

let gameId = '';

async function renderBoard() {
	console.log('render board');
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

	const chessboard = document.getElementById('chessboard');
	chessboard.innerHTML = '';  // Clear any existing content

	board.forEach((row, rowIndex) => {
		row.split('').forEach((square, colIndex) => {
			const squareElement = document.createElement('div');
			squareElement.classList.add('square');

			if ((rowIndex + colIndex) % 2 === 0) {
				squareElement.classList.add('white');
			} else {
				squareElement.classList.add('black');
			}

			const piece = {
				'r': '♜', 'n': '♞', 'b': '♝', 'q': '♛', 'k': '♚', 'p': '♟',
				'R': '♖', 'N': '♘', 'B': '♗', 'Q': '♕', 'K': '♔', 'P': '♙',
			}[square] || '';

			if (piece) {
				const pieceElement = document.createElement('span');
				pieceElement.classList.add('piece');
				pieceElement.innerHTML = piece;
				squareElement.appendChild(pieceElement);
			}

			chessboard.appendChild(squareElement);
		});
	});
}

async function startGame() {
	//~ gameId = await start_anonymous_game();
	console.log(`Game started with ID: ${gameId}`);
	renderBoard();  // Render the initial board state
}

//~ async function handleMove() {
	//~ const move = document.getElementById('move_input').value.trim();
	//~ if (!move || !gameId) {
		//~ alert('Please enter a valid move and start a game.');
		//~ return;
	//~ }

	//~ try {
		//~ await make_move(gameId, move);
		//~ console.log(`Move ${move} made successfully`);
		//~ // Update the board here if needed
		//~ // For now, we'll just print the result to the console
	//~ } catch (error) {
		//~ console.error(`Error making move: ${error}`);
	//~ }
//~ }

//~ document.getElementById('start_game').addEventListener('click', async () => {
	//~ await startGame();
//~ });

//~ document.getElementById('make_move').addEventListener('click', async () => {
	//~ await handleMove();
//~ });

init().then(() => {
	console.log('WASM module initialized: play');
	renderBoard();
});

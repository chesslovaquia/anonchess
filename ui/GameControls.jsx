import React, { useState } from 'react';

function GameControls() {
	const [move, setMove] = useState('');

	const startGame = async () => {
		// Call the Rust function to start an anonymous game
		console.log('Starting game...');
		// Logic to start game goes here
	};

	const makeMove = async () => {
		if (move.trim() === '') {
			alert('Please enter a move.');
			return;
		}
		console.log(`Making move: ${move}`);
		// Logic to make move goes here
	};

	return (
		<div className="w3-center w3-margin-top">
			<button onClick={startGame} className="w3-button w3-blue w3-large">Start Anonymous Game</button>
			<div className="w3-margin-top">
				<input
					type="text"
					value={move}
					onChange={(e) => setMove(e.target.value)}
					className="w3-input w3-border"
					placeholder="Enter move (e.g., e2e4)"
				/>
				<button onClick={makeMove} className="w3-button w3-green w3-margin-top">Make Move</button>
			</div>
		</div>
	);
}

export default GameControls;

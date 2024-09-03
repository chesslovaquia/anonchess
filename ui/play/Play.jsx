// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

import ChessBoard from './ChessBoard';
import SideBar    from './SideBar';

function Play() {
	return (
		<div>
			<div className="w3-container w3-cell w3-mobile">
				<SideBar />
			</div>
			<div className="w3-container w3-cell w3-mobile">
				<ChessBoard />
			</div>
			<div className="w3-container w3-cell w3-mobile">
				<SideBar />
			</div>
		</div>
	);
}

export default Play;

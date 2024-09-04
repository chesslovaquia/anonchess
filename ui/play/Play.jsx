// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

import ChessBoard from './ChessBoard';
import SideWhite  from './SideWhite';
import SideBlack  from './SideBlack';
import SideBar    from './SideBar';

function Play() {
	return (
		<div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<SideBlack />
			</div>
			<div className="w3-cell w3-cell-middle w3-mobile">
				<ChessBoard />
			</div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<SideWhite />
			</div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<SideBar />
			</div>
		</div>
	);
}

export default Play;

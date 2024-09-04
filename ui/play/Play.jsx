// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React, { useRef } from 'react';

import ChessBoard from './ChessBoard';
import SideWhite  from './SideWhite';
import SideBlack  from './SideBlack';
import SideBar    from './SideBar';

function Play() {
	const renderReg = useRef({
		sideBar:    () => {},
		chessBoard: () => {},
	});

	return (
		<div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<SideBar renderReg={renderReg} />
			</div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<SideBlack />
			</div>
			<div className="w3-cell w3-cell-middle w3-mobile">
				<ChessBoard renderReg={renderReg} />
			</div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<SideWhite />
			</div>
		</div>
	);
}

export default Play;

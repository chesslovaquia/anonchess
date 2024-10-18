// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React, { useRef } from 'react';

import ChessBoard from './ChessBoard';
import Player1  from './Player1';
import Player2  from './Player2';
import SideBar    from './SideBar';

function Play() {
	const renderReg = useRef({
		sideBar: () => {},
	});

	return (
		<div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<SideBar renderReg={renderReg} />
			</div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<Player2 />
			</div>
			<div className="w3-cell w3-cell-middle w3-mobile">
				<ChessBoard renderReg={renderReg} />
			</div>
			<div className="w3-cell w3-cell-top w3-mobile">
				<Player1 />
			</div>
		</div>
	);
}

export default Play;

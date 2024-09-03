// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

import ChessBoard from './ChessBoard';
import SideBar    from './SideBar';

function Play() {
	return (
		<div className="w3-container">
			<div className="w3-row">
				<div className="w3-col s1 m1 l1">
					<SideBar />
				</div>
				<div className="w3-col s11 m11 l11">
					<ChessBoard />
				</div>
			</div>
		</div>
	);
}

export default Play;

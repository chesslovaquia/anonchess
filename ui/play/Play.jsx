// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

import ChessBoard from './ChessBoard';
import SideBar    from './SideBar';

function Play() {
	return (
		<div className="w3-container w3-left-align">
			<div className="w3-row">
				<div className="w3-col w3-left s3 m2 l1 w3-mobile">
					<SideBar />
				</div>
				<div className="w3-col w3-right s9 m10 l11">
					<ChessBoard />
				</div>
			</div>
		</div>
	);
}

export default Play;

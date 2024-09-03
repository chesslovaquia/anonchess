// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

import ChessBoard from './ChessBoard';
import SideBar    from './SideBar';

function Play() {
	return (
		<div>
			<div className="w3-sidebar w3-bar-block w3-black" style={{width: '25%'}}>
				<SideBar />
			</div>
			<div style={{marginLeft: '30%'}}>
				<ChessBoard />
			</div>
		</div>
	);
}

export default Play;

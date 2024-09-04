// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

function SideWhite() {
	return (
		<ul className="w3-ul w3-gray">
			<li>{anonc_outcome()} {anonc_turn()} {anonc_turn_name()}</li>
		</ul>
	);
}
export default SideWhite;

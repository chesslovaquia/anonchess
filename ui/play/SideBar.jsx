// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React, { useState, useEffect } from 'react';

function SideBar({renderReg}) {
	const [outcome, setOutcome] = useState(anonc_outcome());
	const [turn, setTurn] = useState(anonc_turn());
	const [turnName, setTurnName] = useState(anonc_turn_name());

	useEffect(() => {
		renderReg.current.sideBar = () => {
			setOutcome(anonc_outcome());
			setTurn(anonc_turn());
			setTurnName(anonc_turn_name());
		};
	}, []);  // Empty dependency array ensures the fetch runs only once

	return (
		<ul className="w3-ul w3-gray">
			<li>{outcome}</li>
		</ul>
	);
}
export default SideBar;

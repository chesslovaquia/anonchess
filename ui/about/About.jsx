// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';

async function getPkg() {
	const pkg = './package.json';
	try {
		const resp = await fetch(pkg);
		if (!resp.ok) {
			alert(`ERROR: could not get ${pkg}`);
			return null;
		}
		console.log('parse:', pkg);
		const data = await resp.json();
		return data;
	} catch (err) {
		alert(`ERROR fetch ${pkg}: ${err.message}`);
	}
	return null
}

function About() {
	getPkg()
		.then(pkg => {
			if (pkg) {
				console.log('Pkg:', pkg);
			}
		});
	return (
		<div className="w3-container">
			<h1>AnonChess</h1>
		</div>
	);
}

export default About;

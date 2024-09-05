// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React, { useState, useEffect } from 'react';

function About() {
	const [data, setData] = useState(null);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const resp = await fetch('./package.json');
				if (!resp.ok) {
					throw new Error('Network failed!');
				}
				const jsonData = await resp.json();
				setData(jsonData);
			} catch (err) {
				setError(err.message);
			} finally {
				setLoading(false);
			}
		};

		fetchData();
	}, []);  // Empty dependency array ensures the fetch runs only once

	if (loading) {
		return (
			<div className="w3-container">
				<h1>AnonChess</h1>
				<h3>loading...</h3>
			</div>
		);
	}

	if (error) {
		return (
			<div className="w3-container">
				<h1>AnonChess</h1>
				<h3>ERROR: {error}</h3>
			</div>
		);
	}

	return (
		<div className="w3-container">
			<h1>AnonChess</h1>
			<h3>{data.version}</h3>
		</div>
	);
}

export default About;

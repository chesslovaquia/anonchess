// Copyright Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

import React from 'react';
import { createRoot } from 'react-dom/client';
import Home from './Home';

import '../wasm_exec.js';

const go = new Go();
WebAssembly.instantiateStreaming(fetch('./pkg/anonchess-home.wasm'), go.importObject).then((result) => {
	go.run(result.instance);
	const root = createRoot(document.getElementById('root'));
	root.render(<Home />);
});

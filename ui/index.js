import React from 'react';
import { createRoot } from 'react-dom/client';
import App from './App';

import './static/pkg/wasm_exec.js';

const go = new Go();
WebAssembly.instantiateStreaming(fetch('./pkg/anonchess.wasm'), go.importObject).then((result) => {

	go.run(result.instance);

	const root = createRoot(document.getElementById('root'));
	root.render(<App />);

});

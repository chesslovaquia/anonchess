const path = require('path');

module.exports = {
	//~ entry: './ui/index.js',
	//~ output: {
		//~ path: path.resolve(__dirname, 'static', 'ui'),
		//~ filename: 'bundle.js',
	//~ },
	module: {
		rules: [
			{
				test: /\.jsx?$/,
				exclude: /node_modules/,
				use: {
					loader: 'babel-loader',
					options: {
						presets: ['@babel/preset-env', '@babel/preset-react'],
					},
				},
			},
		],
	},
	resolve: {
		extensions: ['.js', '.jsx'],
	},
	devServer: {
		contentBase: path.join(__dirname, 'static'),
		compress: true,
		port: 8080,
	},
};

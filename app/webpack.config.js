var path = require('path');
var webpack = require('webpack');

module.exports = {
  // devtool: 'eval',
  entry: [
    './src/index.jsx'
  ],
  output: {
    path : path.join(__dirname, '/build'),
    filename : 'bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        loader: 'babel-loader',
        include: path.join(__dirname, 'src')
      },
      {
        test: /\.scss$/,
        use: [
          'style-loader',
          'css-loader',
          'sass-loader'
        ]
      }
    ]
  }
};

var path = require('path');

module.exports = {
  entry: './js/boids.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'static')
  },
	devtool: 'source-map',
  module: {
    rules: [
      {
        test: /\.js$/,
        use: ["source-map-loader"],
        enforce: "pre"
      },
      {
        test: /\.proto$/,
        use: ["protobuf-loader"]
      }
    ]
  }
};


var path = require('path');

module.exports = {
  entry: './js/boids.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'static')
  }
};


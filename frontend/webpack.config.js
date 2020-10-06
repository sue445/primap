const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const webpack = require('webpack');

module.exports = {
  entry: './app/app.tsx',
  plugins: [
    new CleanWebpackPlugin({
      cleanAfterEveryBuildPatterns: ['public/build']
    }),
    new HtmlWebpackPlugin({
      template: 'templates/index.html'
    }),
    new webpack.EnvironmentPlugin([
      'NODE_ENV',
      'REACT_APP_GOOGLE_BROWSER_API_KEY',
      'REACT_APP_SENTRY_DSN',
    ]),
  ],
  output: {
    path: __dirname + '/public',
    filename: 'build/[name].[contenthash].js'
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js']
  },
  module: {
    rules: [
      { test: /\.tsx?$/, loader: 'ts-loader' }
    ]
  },
}

if (process.env.NODE_ENV != "production") {
  module.exports.devtool = "inline-source-map"
}

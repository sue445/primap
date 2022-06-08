const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyPlugin = require('copy-webpack-plugin');
const devMode = process.env.NODE_ENV !== 'production';

const webpack = require('webpack');

process.traceDeprecation = true;

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
      'REACT_APP_SENTRY_RELEASE',
    ]),
    new MiniCssExtractPlugin({
      // Options similar to the same options in webpackOptions.output
      // both options are optional
      filename: 'build/' + (devMode ? '[name].css' : '[name].[fullhash].css'),
      chunkFilename: 'build/' + (devMode ? '[id].css' : '[id].[fullhash].css'),
    }),
    new CopyPlugin({
      patterns: [
        { from: 'static', to: '.' },
      ],
    }),
  ],
  output: {
    path: __dirname + '/public',
    filename: 'build/[name].[contenthash].js'
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js'],
    fallback: { "url": require.resolve("url/") }
  },
  module: {
    rules: [
      { test: /\.tsx?$/, loader: 'ts-loader' },
      {
        test: /\.css$/,
        use: [
          {
            loader: MiniCssExtractPlugin.loader,
            options: {
              publicPath: '/public/',
            },
          },
          'css-loader',
          {
            loader: 'postcss-loader',
            options: {
              postcssOptions: {
                ident: 'postcss',
                plugins: [
                  require('tailwindcss'),
                  require('autoprefixer'),
                ],
              },
            }
          },
        ],
      },
    ],
  },
  devServer: {
    port: 55301,
  },
  devtool: "source-map",
  stats: {
    errorDetails: true
  }
}

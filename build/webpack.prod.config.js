var webpack = require('webpack')
var ExtractTextPlugin = require('extract-text-webpack-plugin')
var HtmlWebpackPlugin = require('html-webpack-plugin')
var autoprefixer = require('autoprefixer')
var path = require('path')

module.exports = {
  context: path.resolve(__dirname, '../src'),
  entry: {
    app: './index.js',
    vendor: [
      'react',
      'react-dom',
      'react-redux',
      'react-router',
      'react-router-redux',
      'redux',
      'redux-actions',
      './assets/scripts/bootstrap.min.js'
    ]
  },
  output: {
    filename: '[hash]/js/[name].js',
    path: path.resolve(__dirname, '../dist'),
    publicPath: '/'
  },
  module: {
    rules: [{
      test: /\.(js|jsx)$/,
      use: 'babel-loader',
      exclude: /node_modules/
    }, {
      test: /global\.scss$/,
      loader: ExtractTextPlugin.extract({
        fallbackLoader: "style-loader",
        loader: [{
          loader: 'css-loader',
          query: {
            minimize: true
          }
        }, {
          loader: 'postcss-loader'
        }, {
          loader: 'sass-loader'
        }]
      })
    }, {
      test: /^((?!global).)*\.scss$/,
      loader: ExtractTextPlugin.extract({
        fallbackLoader: "style-loader",
        loader: [{
          loader: 'css-loader',
          query: {
            minimize: true,
            modules: true,
            importLoaders: 1,
            localIdentName: '[local]__[hash:base64:5]'
          }
        }, {
          loader: 'postcss-loader'
        }, {
          loader: 'sass-loader'
        }]
      })
    }, {
      test: /\.(png|jpg|jpeg|gif)$/,
      use: [{
        loader: "url-loader",
        options: {
          limit: 1000,
          name: 'images/[name].[ext]'
        }
      }]
    }, {
      test: /\.(eot|svg|ttf|woff|woff2)$/,
      use: [{
        loader: "file-loader",
        options: {
          name: 'fonts/[name].[ext]'
        }
      }]
    }]
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env': {
        'NODE_ENV': JSON.stringify('production')
      }
    }),
    new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery',
      'window.jQuery': 'jquery'
    }),
    new webpack.optimize.CommonsChunkPlugin({
      name: 'vendor'
    }),
    new webpack.optimize.UglifyJsPlugin({
      compress:{
        warnings: true
      }
    }),
    new ExtractTextPlugin({
      filename: '[hash]/css/app.css',
      disable: false,
      allChunks: true
    }),
    new webpack.LoaderOptionsPlugin({
      options: {
        postcss: [autoprefixer],
      },
    }),
    new HtmlWebpackPlugin({
      template: './index.html'
    })
  ]
}
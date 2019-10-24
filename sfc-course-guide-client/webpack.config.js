var HtmlWebpackPlugin = require('html-webpack-plugin');
var MiniCssExtractPlugin = require('mini-css-extract-plugin');
var VueLoaderPlugin = require('vue-loader/lib/plugin')
var Path = require('path');
var Webpack = require('webpack');

var jsFiles = ['main']
var cssFiles = ['style']
var htmlFiles = ['index']

module.exports = {
    mode: "development",
    devtool: 'cheap-module-eval-source-map',
    entry: jsFiles.map(function (name) {
        return './src/javascripts/' + name + '.js'
    }).concat(cssFiles.map(function (name) {
        return './src/stylesheets/' + name + '.css'
    })),
    output: {
        path: Path.resolve(__dirname, 'dist'),
        filename: 'bundle.js'
    },
    plugins: [
        new Webpack.HotModuleReplacementPlugin(),
        new VueLoaderPlugin(),
    ].concat(cssFiles.map(function (name) {
        return new MiniCssExtractPlugin({
            filename: name + '.css'
        })
    })).concat(htmlFiles.map(function (name) {
        return new HtmlWebpackPlugin({
            template: './src/views/' + name + '.html',
            filename: name + '.html'
        })
    })),
    module: {
        rules: [
            {
                test: /\.css$/,
                use: [
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                ],
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            }
        ]
    },
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.esm.js'
        }
    },
    devServer: {
        hot: true,
        open: true,
        openPage: '',
        port: 8080,
        watchContentBase: true
    }
};

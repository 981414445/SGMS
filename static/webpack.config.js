const webpack = require('webpack');
const path = require('path');
const fs = require('fs');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const resolveApp = relativePath =>
    path.resolve(fs.realpathSync(process.cwd()), relativePath);

const appEntry = 'js/entry/';

/**
 * 扫描函数
 */
function Scan() {
    const dirs = fs.readdirSync(resolveApp(appEntry));
    const map = {};
    dirs.forEach(file => {
        const entry = file.replace(/\.js$/, '');
        map[entry] = resolveApp(appEntry + file);
    });
    return map;
}
const dirs = Scan();

module.exports = {
    entry: dirs,
    output: {
        path: __dirname + '/build',
        filename: '[name].js'
    },
    devtool: 'source-map',
    devServer: {
        contentBase: './public',
        port: 4444, // 监听端口号
        inline: true // 实时刷新
    },
    resolve: {
        extensions: ['.js', '.less']
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['env']
                    }
                },
                exclude: /node_moudles/
            },
            {
                test: /\.(css|less)$/,
                use: [
                    {
                        loader: 'style-loader'
                    },
                    {
                        loader: 'css-loader'
                    },
                    {
                        loader: 'less-loader'
                    }
                ]
            }
        ]
    },
    plugins: [
        new CleanWebpackPlugin(['build']),
        new webpack.BannerPlugin('版权所有，翻版必究') // 在打包后的代码中添加注释
        // new webpack.optimize.UglifyJsPlugin(), //压缩js代码
    ]
};

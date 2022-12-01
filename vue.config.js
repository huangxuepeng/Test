
module.exports = {

    // 它支持webPack-dev-server的所有选项
    devServer: {
        port: 8005,
        proxy: {
            // 第一台服务器配置 
            '/u/v1': {
            target: 'http://39.105.117.187:8083',
            ws: true,
            changeOrigin: true,
            },

            '/java': {
                target: 'http://39.105.117.187:8082',
                ws: true,
                changeOrigin: true,
            }
        }
}
}
  
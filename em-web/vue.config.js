const {
  defineConfig
} = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    port: 8090,
    // 跨域问题解决 代理
    proxy: {
      '/api': {
        //后端提供的真实接口
        target: 'http://localhost:9080',
        //允许跨域
        changeOrigin: true,
        // 如果接口中是没有api的，那就直接置空，'^/api': ''，
        // 如果接口中有api，那就得写成{'^/api':'/api'}
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
})
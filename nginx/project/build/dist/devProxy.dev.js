"use strict";

// 代理
var _Proxy = {
  '/': {
    target: 'http://39.105.117.187:8085',
    secure: false,
    // false为http访问，true为https访问
    changeOrigin: true // 跨域访问设置，true代表跨域
    // pathRewrite: { '^/api': '/api' }

  } // '/': {
  //     target: 'http://192.168.91.123:8080',
  //     secure: false, // false为http访问，true为https访问
  //     changeOrigin: true // 跨域访问设置，true代表跨域
  //     // pathRewrite: { '^/api': '/api' }
  // }

};
module.exports = {
  _Proxy: _Proxy
};
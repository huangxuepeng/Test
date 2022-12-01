"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.test8081 = exports.test8080 = void 0;

var _axiosHttp = require("@/services/axios-http.js");

// let config = {} 这个里面可以传入一些基地址 然后传入下面的方法里面  例如 const API = request(config)
var _API = (0, _axiosHttp.request)(); // 配置请求头的这样类似的参数
// const _initConfig = {};
// const _authAPI = request(_initConfig);

/**
 * @params : method url params
 * @return : 接口响应的数据
 * @description :
 * @date : 2020-10-30 09:21
 * @author : huanghe
 */
// export const _exampleAPI = (params) => _API('GET', 'api/test', params);


var test8080 = function test8080(params) {
  return _API('GET', 'test', params);
};

exports.test8080 = test8080;

var test8081 = function test8081(params) {
  return _API('GET', 'api/test', params);
};

exports.test8081 = test8081;
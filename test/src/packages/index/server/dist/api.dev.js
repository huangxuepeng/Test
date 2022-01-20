"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports._exampleAPI = void 0;

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


var _exampleAPI = function _exampleAPI(params) {
  return _API('get', 'http://127.0.0.1:8080/u/v1/user/list', params);
};

exports._exampleAPI = _exampleAPI;
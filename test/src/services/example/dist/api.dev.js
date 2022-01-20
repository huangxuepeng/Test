"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.exampleAPI = void 0;

var _axiosHttp = require("@/services/axios-http.js");

// let config = {} 这个里面可以传入一些基地址 然后传入下面的方法里面  例如 const API = request(config)
var API = (0, _axiosHttp.request)(); // 配置请求头的这样类似的参数

var initConfig = {};
var authAPI = (0, _axiosHttp.request)(initConfig);
/**
 * @params : method url params
 * @return : 接口响应的数据
 * @description :
 * @date : 2020-10-30 09:21
 * @author : huanghe
 */

var exampleAPI = function exampleAPI(params) {
  return API('GET', '', params);
};

exports.exampleAPI = exampleAPI;
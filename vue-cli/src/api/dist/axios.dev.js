"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _axios = _interopRequireDefault(require("axios"));

var _index = _interopRequireDefault(require("../config/index"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(source, true).forEach(function (key) { _defineProperty(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(source).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

// 设置配置 根据开发 和 生产环境不一样
var baseUrl = process.env.NODE_ENV === 'development' ? _index["default"].baseUrl.dev : _index["default"].baseUrl.pro;

var httpRequest =
/*#__PURE__*/
function () {
  function httpRequest(baseUrl) {
    _classCallCheck(this, httpRequest);

    this.baseUrl = baseUrl;
  }

  _createClass(httpRequest, [{
    key: "getInsideConfig",
    value: function getInsideConfig() {
      var config = {
        baseUrl: this.baseUrl
      };
      return config;
    }
  }, {
    key: "interceptors",
    value: function interceptors(instance) {
      // 添加请求拦截器
      instance.interceptors.request.use(function (config) {
        // 在发送请求之前做些什么
        console.log('拦截处理请求');
        return config;
      }, function (error) {
        console.log(error); // 对请求错误做些什么

        return Promise.reject(error);
      }); // 添加响应拦截器

      instance.interceptors.response.use(function (response) {
        console.log('处理响应'); // 对响应数据做点什么

        return response;
      }, function (error) {
        console.log(error); // 对响应错误做点什么

        return Promise.reject(error);
      });
    }
  }, {
    key: "request",
    value: function request(options) {
      // 请求
      // /api/getList    /api/getHome
      var instance = _axios["default"].create();

      options = _objectSpread({}, this.getInsideConfig, {}, options);
      this.interceptors(instance);
      return instance(options);
    }
  }]);

  return httpRequest;
}();

var _default = new httpRequest(baseUrl);

exports["default"] = _default;
"use strict";

var _vue = _interopRequireDefault(require("vue"));

var _App = _interopRequireDefault(require("./App.vue"));

var _router = _interopRequireDefault(require("./router"));

var _elementUi = _interopRequireDefault(require("element-ui"));

require("element-ui/lib/theme-chalk/index.css");

var _axios = _interopRequireDefault(require("axios"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

// 配置请求的根路径
// axios.default.baseURL = '/u/v1'
_vue["default"].config.productionTip = false;

_vue["default"].use(_elementUi["default"]); // axios请求拦截


_axios["default"].interceptors.request.use(function (config) {
  config.headers.Authorization = window.sessionStorage.getItem('token'); // console.log(window.sessionStorage.getItem('token'))

  return config;
});

_vue["default"].prototype.$http = _axios["default"];
new _vue["default"]({
  router: _router["default"],
  render: function render(h) {
    return h(_App["default"]);
  }
}).$mount('#app');
"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _vue = _interopRequireDefault(require("vue"));

var _vueRouter = _interopRequireDefault(require("vue-router"));

var _Login = _interopRequireDefault(require("../components/Login.vue"));

var _Home = _interopRequireDefault(require("../components/Home.vue"));

var _Welcome = _interopRequireDefault(require("../components/Welcome.vue"));

var _Users = _interopRequireDefault(require("../components/user/Users.vue"));

var _Delete = _interopRequireDefault(require("../components/user/Delete.vue"));

var _Add = _interopRequireDefault(require("../components/user/Add.vue"));

require("../assets/css/global.css");

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

_vue["default"].use(_vueRouter["default"]);

var routes = [{
  path: '/',
  redirect: '/login'
}, {
  path: '/login',
  component: _Login["default"]
}, {
  path: '/home',
  component: _Home["default"],
  redirect: '/wecome',
  children: [{
    path: '/wecome',
    component: _Welcome["default"]
  }, {
    path: '/list',
    component: _Users["default"]
  }, {
    path: '/delete',
    component: _Delete["default"]
  }, {
    path: '/add',
    component: _Add["default"]
  }]
}];
var router = new _vueRouter["default"]({
  routes: routes
}); // 挂载路由导航守卫

router.beforeEach(function (to, from, next) {
  // to 将要访问的路径
  // from 代表从哪个路径跳转而来
  // next 是一个函数, 表示放行 next() //放行   next('/login') //跳转
  if (to.path === '/login') return next(); // 获取token

  var tokenStr = window.sessionStorage.getItem('token');
  if (!tokenStr) return next('/login');
  next();
});
var _default = router;
exports["default"] = _default;
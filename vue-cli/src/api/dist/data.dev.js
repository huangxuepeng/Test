"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.getMenu = void 0;

var _axios = _interopRequireDefault(require("./axios"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var getMenu = function getMenu() {
  return _axios["default"].request({
    url: '/u/v1/user/list',
    method: 'get'
  });
};

exports.getMenu = getMenu;
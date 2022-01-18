"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;
var _default = {
  state: {
    isCollapse: true,
    message: "ojsiuj"
  },
  mutations: {
    collapseMenu: function collapseMenu(state) {
      state.isCollapse = !state.isCollapse;
    }
  }
};
exports["default"] = _default;
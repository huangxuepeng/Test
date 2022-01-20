import Vue from 'vue';
import Vuex from 'vuex';

import indexStore from '@/store/indexPage/index.js';

Vue.use(Vuex);
const store = new Vuex.Store({
    state: {},
    mutations: {},
    actions: {},
    getters: {},
    modules: {
        indexStore
    }
});

export default store;

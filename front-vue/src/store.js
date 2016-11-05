import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

import root from './engines'

const store = new Vuex.Store(root.store)
export default store

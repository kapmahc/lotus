require('./assets/main.css')
require('./assets/reading.css')

import Vue from 'vue'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import VueI18n from 'vue-i18n'

import routes from './engines/routes'
import store from './engines/store'
import {initLocales} from './i18n'

Vue.use(Vuex)
Vue.use(VueRouter)
Vue.use(VueI18n)
new Vuex.Store(store)
initLocales()
new Vue({
  router
}).$mount('#root')

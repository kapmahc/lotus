import 'bootstrap/scss/bootstrap.scss'
import './assets/main.css'
import './assets/reading.css'

import 'tether'
import 'jquery'
import 'bootstrap'

import Vue from 'vue'
import Vuex from 'vuex'
import VueI18n from 'vue-i18n'

import router from './engines/router'
import store from './engines/store'
import {initLocales} from './i18n'

Vue.use(Vuex)
Vue.use(VueI18n)

/* eslint-disable no-new */
initLocales()
new Vue({
  store: new Vuex.Store(store),
  router
}).$mount('#root')

require('./assets/main.css')
require('./assets/reading.css')

import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

import root from './engines'
import {initLocales} from './i18n'

initLocales()
/* eslint-disable no-new */
new Vue({
  router: new VueRouter({routes: root.routes})
}).$mount('#root')

require('./assets/main.css')
require('./assets/reading.css')

import Vue from 'vue'

import root from './engines'
import i18n from './i18n'

Vue.use(i18n)
/* eslint-disable no-new */
new Vue({
  router: root.router
}).$mount('#root')

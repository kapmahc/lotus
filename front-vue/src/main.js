require('./assets/main.css')
require('./assets/reading.css')

import Vue from 'vue'

import root from './engines'

/* eslint-disable no-new */
new Vue({
  router: root.router
}).$mount('#root')

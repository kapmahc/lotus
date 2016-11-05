require('./assets/main.css')
require('./assets/reading.css')

import Vue from 'vue'
import VueRouter from 'vue-router'

import root from './engines'

Vue.use(VueRouter)
const routes = root.routes()
const router = new VueRouter({
  routes
})

/* eslint-disable no-new */
new Vue({
  router
}).$mount('#root')

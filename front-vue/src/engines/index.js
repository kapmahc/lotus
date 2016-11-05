import Vue from 'vue'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
Vue.use(Vuex)
Vue.use(VueRouter)

import auth from './auth'
import Home from './Home'

const engines = {
  auth
}

const routes = Object.keys(engines).reduce(function (obj, en) {
  return obj.concat(engines[en].routes)
}, [])
routes.push({path: '/', component: Home})

const store = {
  state: Object.keys(engines).reduce(function (obj, en) {
    return Object.assign(obj, engines[en].state)
  }, {}),
  mutations: Object.keys(engines).reduce(function (obj, en) {
    return Object.assign(obj, engines[en].mutations)
  }, {})
}

// console.log(store)
// console.log(routes)

const root = {
  store: new Vuex.Store(store),
  router: new VueRouter({routes})
}

export default root

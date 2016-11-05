require('./assets/main.css')
require('./assets/reading.css')

import Vue from 'vue'
import VueI18n from 'vue-i18n'
Vue.use(VueI18n)

import root from './engines'

// var self = window.this
// var lang = 'en-US'
// Vue.locale(lang, function () {
//   self.loading = true
//   return window.fetch(process.env.API_HOST + '/locales/' + lang, {
//     method: 'get',
//     headers: {
//       'Accept': 'application/json',
//       'Content-Type': 'application/json'
//     }
//   }).then(function (res) {
//     return res.json()
//   }).then(function (json) {
//     self.loading = false
//     if (Object.keys(json).length === 0) {
//       return Promise.reject(new Error('locale empty !!'))
//     } else {
//       return Promise.resolve(json)
//     }
//   }).catch(function (error) {
//     self.error = error.message
//     return Promise.reject()
//   })
// }, function () {
//   Vue.config.lang = lang
// })

var lang = window.localStorage.getItem('locale') || 'zh-CN'

Vue.locale(lang, function () {
  return window.fetch(`${process.env.API_HOST}/locales/${lang}`, {
    method: 'get',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  }).then(function (res) {
    return res.json()
  }).then(function (json) {
    if (Object.keys(json).length === 0) {
      return Promise.reject(new Error('locale empty !!'))
    } else {
      return Promise.resolve(json)
    }
  }).catch(function (error) {
    console.log(error.message)
    return Promise.reject()
  })
}, function () {
  Vue.config.lang = lang
})

/* eslint-disable no-new */
new Vue({
  router: root.router
}).$mount('#root')

import Vue from 'vue'

import {api} from './utils'

const key = 'locale'

function loadLocales (lang) {
  Vue.locale(lang, function () {
    return window.fetch(api(`/locales/${lang}`), {
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
}

export const swtichLang = (lang) => {
  window.localStorage.setItem(key, lang)
  loadLocales(lang)
}

export const initLocales = () => {
  loadLocales(window.localStorage.getItem(key) || 'en-US')
}

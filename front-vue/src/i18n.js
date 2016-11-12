import Vue from 'vue'
import moment from 'moment'

import {api} from './utils'

const key = 'locale'

export function currentLocale () {
  var locale = window.localStorage.getItem(key)
  if (locale == null) {
    locale = 'en-US'
    window.localStorage.setItem(key, locale)
  }
  return locale
}

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
    // Vue.config.fallbackLang = 'en-US'
  })
}

export const swtichLang = (lang) => {
  window.localStorage.setItem(key, lang)
  moment.locale(lang)
  loadLocales(lang)
}

export const initLocales = () => {
  var lang = currentLocale()
  moment.locale(lang)
  loadLocales(lang)
}

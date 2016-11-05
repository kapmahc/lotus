import i18next from 'i18next'
import XHR from 'i18next-xhr-backend'
import LanguageDetector from 'i18next-browser-languagedetector'

const key = 'locale'

i18next
  .use(XHR)
  .use(LanguageDetector)
  .init({
    fallbackLng: 'en-US',
    backend: {
      loadPath: process.env.API_HOST + '/locales/{{lng}}',
      crossDomain: true
    },
    detection: {
      order: ['querystring', 'cookie', 'localStorage', 'navigator', 'htmlTag'],

      lookupQuerystring: key,
      lookupCookie: key,
      lookupLocalStorage: key,

      // cache user language on
      caches: ['localStorage', 'cookie'],

      // optional expire and domain for set cookie
      cookieMinutes: 60 * 24 * 7,

      // optional htmlTag with lang attribute, the default is:
      htmlTag: document.documentElement
    }
  })

const I18nPlugin = {}
I18nPlugin.install = function (Vue, options) {
  Vue.prototype.$t = function (msg, options) {
    return i18next.t(msg, options)
  }
}

export default I18nPlugin

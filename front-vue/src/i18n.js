import i18n from 'i18next'
import XHR from 'i18next-xhr-backend'
import LanguageDetector from 'i18next-browser-languagedetector'

const key = 'locale'

export default function () {
  i18n
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
}

import jwtDecode from 'jwt-decode'

import {api} from '../../utils'

export default {
  state: {
    currentUser: {},
    siteInfo: {
      author: {}
    }
  },
  mutations: {
    signIn (state, token) {
      try {
        state.currentUser = jwtDecode(token)
      } catch (e) {
        console.log(e)
      }
    },
    signOut (state) {
      state.currentUser = {}
    },
    refreshLayout (state) {
      window.fetch(api(`/layout`)).then(function (res) {
        return res.json()
      }).then(function (info) {
        state.siteInfo = info
      }).catch(function (error) {
        console.log(error.message)
        return Promise.reject()
      })
    }
  }
}

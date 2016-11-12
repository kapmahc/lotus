import jwtDecode from 'jwt-decode'

import {get, _delete} from '../../utils'

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
      _delete('/users/sign-out', function () {
        window.sessionStorage.removeItem('token')
      })
    },
    refreshLayout (state) {
      get('/layout', null, function (info) {
        window.document.title = `${info.subTitle}-${info.title}`
        state.siteInfo = info
      })
      // window.fetch(api(`/layout`)).then(function (res) {
      //   return res.json()
      // }).then(function (info) {
      //   state.siteInfo = info
      // }).catch(function (error) {
      //   console.log(error.message)
      //   return Promise.reject()
      // })
    }
  }
}

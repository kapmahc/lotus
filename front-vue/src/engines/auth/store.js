import {api} from '../../utils'

export default {
  state: {
    siteInfo: {
      author: {}
    }
  },
  mutations: {
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

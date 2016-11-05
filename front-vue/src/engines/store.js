import auth from './auth/store'

const engines = {
  auth
}

const store = {
  state: Object.keys(engines).reduce(function (obj, en) {
    return Object.assign(obj, engines[en].state)
  }, {}),
  mutations: Object.keys(engines).reduce(function (obj, en) {
    return Object.assign(obj, engines[en].mutations)
  }, {})
}

export default store

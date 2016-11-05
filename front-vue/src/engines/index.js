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

export default {
  store,
  routes
}

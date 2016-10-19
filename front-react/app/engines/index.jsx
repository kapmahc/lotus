import auth from './auth'

const engines = {
  auth
}

export default {
  routes () {
    return Object.keys(engines).reduce(function (obj, en) {
      return obj.concat(engines[en].routes)
    }, [])
    // return CHAOS_ENV.engines.reduce(function(obj, en) {
    //   return obj.concat(engines[en].routes)
    // }, []);
  },
  reducers () {
    return Object.keys(engines).reduce(function (obj, en) {
      return Object.assign(obj, engines[en].reducers)
    }, {})
    // return CHAOS_ENV.engines.reduce(function(obj, en) {
    //   return Object.assign(obj, engines[en].reducers)
    // }, {});
  }
}

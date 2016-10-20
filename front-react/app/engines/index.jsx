import React from 'react'
import {Route} from 'react-router'

import auth from './auth'
import Dashboard from '../components/dashboard'

const engines = {
  auth
}

export default {
  routes () {
    var items = Object.keys(engines).reduce(function (obj, en) {
      return obj.concat(engines[en].routes)
    }, [])
    items.push(<Route key="dashboard" path="dashboard" component={Dashboard}/>)
    return items
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

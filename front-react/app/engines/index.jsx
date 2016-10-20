import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'

import {isEmpty} from './auth/utils'
import NoMatch from '../components/NoMatch'

const DashboardW = ({t, user}) => {
  return isEmpty(user)
    ? <NoMatch/>
    : (
      <div>dashboard</div>
    )
}

DashboardW.propTypes = {
  user: PropTypes.object.isRequired,
  t: PropTypes.func.isRequired
}

const DashboardM = connect(
  state => ({user: state.currentUser}),
  dispatch => ({})
)(DashboardW)

const Dashboard = translate()(DashboardM)

// -----------------------------------------------------------------------------
import {Route} from 'react-router'

import auth from './auth'

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

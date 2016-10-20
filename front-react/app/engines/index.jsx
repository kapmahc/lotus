import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'

import {isEmpty} from './auth/utils'
import NoMatch from '../components/NoMatch'

const DashboardW = ({t, user}) => {
  return isEmpty(user)
    ? <NoMatch/>
    : (
      <div className="row">
        {CONFIG.engines.map(function (en) {
          return engines[en].dashboard
        })}
      </div>
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
import {Route, IndexRoute} from 'react-router'

import auth from './auth'

const engines = {
  auth
}

export default {
  routes () {
    // var items = Object.keys(engines).reduce(function (obj, en) {
    //   return obj.concat(engines[en].routes)
    // }, [])
    var items = CONFIG.engines.reduce(function (obj, en) {
      return obj.concat(engines[en].routes)
    }, [])
    items.push(<Route key="dashboard" path="dashboard" component={Dashboard}/>)
    items.push(<IndexRoute key="home" component={engines[CONFIG.main].home}/>)
    return items
  },
  reducers () {
    // return Object.keys(engines).reduce(function (obj, en) {
    //   return Object.assign(obj, engines[en].reducers)
    // }, {})
    return CONFIG.engines.reduce(function (obj, en) {
      return Object.assign(obj, engines[en].reducers)
    }, {})
  }
}

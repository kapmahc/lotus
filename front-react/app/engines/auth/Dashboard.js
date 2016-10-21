import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'
import {ButtonGroup, Button} from 'react-bootstrap'

import {isAdmin} from './utils'
import {Logs as SelfLogs, Profile as SelfProfile, Password as SelfPassword} from './self'
import {showUserLogs, showUserProfile, toggleUserPassword} from './actions'
import {get} from '../../ajax'

const Widget = ({user, t, onShow}) => {
  var admin = <br/>
  if (isAdmin(user)) {
    admin = (
      <fieldset>
        <legend>{t('auth.dashboard.site')}</legend>
        <ButtonGroup>
          {['status', 'info', 'seo', 'users', 'leavewords', 'notices'].map(function (n, i) {
            n = `site-${n}`
            return <Button key={i} onClick={() => onShow(n)}>{t(`auth.dashboard.${n}`)}</Button>
          })}
        </ButtonGroup>
        <SelfLogs/>
        <SelfProfile/>
        <SelfPassword/>
      </fieldset>
    )
  }
  return (
    <div className="col-md-12">
      <fieldset>
        <legend>{t('auth.dashboard.self')}</legend>
        <ButtonGroup>
          {['profile', 'password', 'logs'].map(function (n, i) {
            n = `self-${n}`
            return <Button key={i} onClick={() => onShow(n)}>{t(`auth.dashboard.${n}`)}</Button>
          })}
        </ButtonGroup>
      </fieldset>
      <br/>
      {admin}
      <br/>
    </div>
  )
}

Widget.propTypes = {
  user: PropTypes.object.isRequired,
  t: PropTypes.func.isRequired,
  onShow: PropTypes.func.isRequired
}

const Model = connect(
  state => ({user: state.currentUser}),
  dispatch => ({
    onShow: function (act) {
      switch (act) {
        case 'self-logs':
          get('/self/logs', function (logs) {
            dispatch(showUserLogs(logs))
          })
          break
        case 'self-profile':
          get('/self/profile', function (info) {
            dispatch(showUserProfile(info))
          })
          break
        case 'self-password':
          dispatch(toggleUserPassword())
          break
        default:
          console.log(`click: ${act}`)
          break
      }
    }
  })
)(Widget)

export default translate()(Model)

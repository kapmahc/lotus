import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'
import {ButtonGroup, Button} from 'react-bootstrap'

import {isAdmin} from './utils'
import {Logs as UserLogs} from './users'
import {showUserLogs} from './actions'
import {get} from '../../ajax'

const Widget = ({user, t, onShow}) => {
  var admin = <br/>
  if (isAdmin(user)) {
    admin = (
      <fieldset>
        <legend>{t('auth.dashboard.site')}</legend>
        <ButtonGroup>
          <Button>{t('auth.dashboard.site-status')}</Button>
          <Button>{t('auth.dashboard.site-info')}</Button>
          <Button>{t('auth.dashboard.site-seo')}</Button>
          <Button>{t('auth.dashboard.site-users')}</Button>
          <Button>{t('auth.dashboard.site-leavewords')}</Button>
          <Button>{t('auth.dashboard.site-notices')}</Button>
        </ButtonGroup>
        <UserLogs/>
      </fieldset>
    )
  }
  return (
    <div className="col-md-12">
      <fieldset>
        <legend>{t('auth.dashboard.self')}</legend>
        <ButtonGroup>
          {['profile', 'logs'].map(function (n, i) {
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
          get('/users/logs', function (logs) {
            dispatch(showUserLogs(logs))
          })
          break
        default:
      }
    }
  })
)(Widget)

export default translate()(Model)

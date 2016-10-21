import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'
import {ButtonGroup, Button} from 'react-bootstrap'

import {isAdmin} from './utils'
import {Logs as SelfLogs, Profile as SelfProfile, Password as SelfPassword} from './self'
import {
  toggleSiteStatus,
  toggleSiteAuthor,
  toggleSiteBase,
  toggleSiteNav,

  toggleUserLogs,
  toggleUserProfile,
  toggleUserPassword
} from './actions'
import {get} from '../../ajax'
import {
  Status as SiteStatus,
  Nav as SiteNav,
  Author as SiteAuthor,
  Base as SiteBase
} from './admin'

const Widget = ({user, t, onShow}) => {
  var admin = <br/>
  if (isAdmin(user)) {
    admin = (
      <fieldset>
        <legend>{t('auth.dashboard.site')}</legend>
        <ButtonGroup>
          {['status', 'base', 'author', 'seo', 'nav', 'users', 'leavewords', 'notices'].map(function (n, i) {
            n = `site-${n}`
            return <Button key={i} onClick={() => onShow(n)}>{t(`auth.dashboard.${n}`)}</Button>
          })}
        </ButtonGroup>
        <SiteStatus/>
        <SiteBase/>
        <SiteAuthor/>
        <SiteNav/>
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
        <SelfLogs/>
        <SelfProfile/>
        <SelfPassword/>
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
            dispatch(toggleUserLogs(logs))
          })
          break
        case 'self-profile':
          get('/self/profile', function (info) {
            dispatch(toggleUserProfile(info))
          })
          break
        case 'self-password':
          dispatch(toggleUserPassword())
          break
        case 'site-status':
          get('/site/status', function (status) {
            dispatch(toggleSiteStatus(status))
          })
          break
        case 'site-base':
          get('/site/base', function (info) {
            dispatch(toggleSiteBase(info))
          })
          break
        case 'site-author':
          get('/site/author', function (info) {
            dispatch(toggleSiteAuthor(info))
          })
          break
        case 'site-nav':
          get('/site/nav', function (info) {
            dispatch(toggleSiteNav(info))
          })
          break
        default:
          console.log(`click: ${act}`)
          break
      }
    }
  })
)(Widget)

export default translate()(Model)

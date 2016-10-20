import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'
import {NavDropdown, MenuItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap'

import {signIn} from '../engines/auth/actions'
import {isEmpty} from '../engines/auth/utils'

const Widget = React.createClass({
  componentDidMount () {
    const {onLoad} = this.props
    onLoad()
  },
  render () {
    const {t, user} = this.props
    return isEmpty(user) ? (
    <NavDropdown eventKey={4} title={t('auth.sign-in-or-up')} id="personal-bar">
      <LinkContainer to={{pathname: '/users/sign-in'}}>
        <MenuItem eventKey={4.1}>{t('auth.sign-in')}</MenuItem>
      </LinkContainer>
      <LinkContainer to={{pathname: '/users/sign-up'}}>
        <MenuItem eventKey={4.2}>{t('auth.sign-up')}</MenuItem>
      </LinkContainer>
      <MenuItem divider />
      <LinkContainer to={{pathname: '/users/forgot-password'}}>
        <MenuItem eventKey={4.2}>{t('auth.forgot-password')}</MenuItem>
      </LinkContainer>
      <LinkContainer to={{pathname: '/users/confirm'}}>
        <MenuItem eventKey={4.2}>{t('auth.confirm')}</MenuItem>
      </LinkContainer>
      <LinkContainer to={{pathname: '/users/unlock'}}>
        <MenuItem eventKey={4.2}>{t('auth.unlock')}</MenuItem>
      </LinkContainer>
    </NavDropdown>
    ) : (
      <NavDropdown eventKey={4} title={t('auth.welcome', {name: user.name})} id="personal-bar">
        <LinkContainer to={{pathname: '/users/confirm'}}>
          <MenuItem eventKey={4.2}>{t('auth.confirm')}</MenuItem>
        </LinkContainer>
        <LinkContainer to={{pathname: '/users/unlock'}}>
          <MenuItem eventKey={4.2}>{t('auth.unlock')}</MenuItem>
        </LinkContainer>
      </NavDropdown>
    )
  }
})

Widget.propTypes = {
  user: PropTypes.object.isRequired,
  t: PropTypes.func.isRequired,
  onLoad: PropTypes.func.isRequired
}

const Model = connect(
  state => ({user: state.currentUser}),
  dispatch => ({
    onLoad: function () {
      var token = window.sessionStorage.getItem('token')
      if (token) {
        dispatch(signIn(token))
      }
    }
  })
)(Widget)

export default translate()(Model)

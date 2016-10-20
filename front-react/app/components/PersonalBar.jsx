import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import {NavDropdown, MenuItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap'

const Widget = ({t}) => (
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
)

Widget.propTypes = {
  t: PropTypes.func.isRequired
}

export default translate()(Widget)

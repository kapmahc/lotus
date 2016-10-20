import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'
import {ButtonGroup, Button} from 'react-bootstrap'

import {isAdmin} from './utils'

const Widget = ({user, t}) => {
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
        </ButtonGroup>
      </fieldset>
    )
  }
  return (
    <div className="col-md-12">
      <fieldset>
        <legend>{t('auth.dashboard.self')}</legend>
        <ButtonGroup>
          <Button>{t('auth.dashboard.self-profile')}</Button>
          <Button>{t('auth.dashboard.self-logs')}</Button>
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
  t: PropTypes.func.isRequired
}

const Model = connect(
  state => ({user: state.currentUser}),
  dispatch => ({})
)(Widget)

export default translate()(Model)

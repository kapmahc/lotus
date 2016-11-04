import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'

import {isEmpty} from '../engines/auth/utils'
import NoMatch from './NoMatch'

const Widget = ({t, user}) => {
  return isEmpty(user)
    ? <NoMatch/>
    : (
      <div>dashboard</div>
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

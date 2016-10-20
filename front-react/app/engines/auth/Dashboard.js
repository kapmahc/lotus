import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import { translate } from 'react-i18next'

import {isAdmin} from './utils'

const Widget = ({user}) => {
  var admin = <br/>
  if (isAdmin(user)) {
    admin = (
      <fieldset>
        <legend>admin board</legend>
      </fieldset>
    )
  }
  return (
    <div className="col-md-12">
      <fieldset>
        <legend>auth board</legend>
      </fieldset>
      {admin}
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

import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import TimeAgo from 'react-timeago'
import {Modal, ListGroup, ListGroupItem,
  HelpBlock,
  Button, FormGroup, ControlLabel, FormControl
} from 'react-bootstrap'

import {toggleSiteStatus} from './actions'

const StatusW = ({t, status, onClose}) => (
  <Modal show={status.show} onHide={onClose}>
    <Modal.Header closeButton>
      <Modal.Title>{t('auth.dashboard.site-status')}</Modal.Title>
    </Modal.Header>
    <Modal.Body>
      <ListGroup>
        <ListGroupItem>aaa</ListGroupItem>
      </ListGroup>
    </Modal.Body>
    <Modal.Footer>
      <Button onClick={onClose}>{t('buttons.close')}</Button>
    </Modal.Footer>
  </Modal>
)

StatusW.propTypes = {
  t: PropTypes.func.isRequired,
  status: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const StatusM = connect(
  state => ({status: state.siteStatus}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleSiteStatus())
    }
  })
)(StatusW)

export const Status = translate()(StatusM)

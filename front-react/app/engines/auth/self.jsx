import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import TimeAgo from 'react-timeago'
import {Modal, ListGroup, ListGroupItem,
  Button
} from 'react-bootstrap'

import {hideUserLogs} from './actions'

const LogsW = ({t, logs, onClose}) => (
  <Modal show={logs.show} onHide={onClose}>
    <Modal.Header closeButton>
      <Modal.Title>{t('auth.dashboard.self-logs')}</Modal.Title>
    </Modal.Header>
    <Modal.Body>
      <ListGroup>
        {logs.items.map(function (l, i) {
          return (
            <ListGroupItem key={i}>
              <TimeAgo date={l.created_at}/>
              : {l.message}
            </ListGroupItem>
          )
        })}
      </ListGroup>
    </Modal.Body>
    <Modal.Footer>
      <Button onClick={onClose}>{t('buttons.close')}</Button>
    </Modal.Footer>
  </Modal>
)

LogsW.propTypes = {
  t: PropTypes.func.isRequired,
  logs: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const LogsM = connect(
  state => ({logs: state.userLogs}),
  dispatch => ({
    onClose: function () {
      dispatch(hideUserLogs())
    }
  })
)(LogsW)

export const Logs = translate()(LogsM)

// -----------------------------------------------------------------------------

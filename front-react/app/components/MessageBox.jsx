import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import {Modal, Button} from 'react-bootstrap'
import {messageBox} from '../engines/auth/actions'

const Widget = ({t, onClose, info}) => (
  <Modal show={info.show} onHide={onClose}>
    <Modal.Header closeButton>
      <Modal.Title>{info.title}</Modal.Title>
    </Modal.Header>
    <Modal.Body>
      <p>{info.body}</p>
    </Modal.Body>
    <Modal.Footer>
      <Button onClick={onClose}>{t('buttons.close')}</Button>
    </Modal.Footer>
  </Modal>
)

Widget.propTypes = {
  t: PropTypes.func.isRequired,
  info: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const Model = connect(
  state => ({info: state.messageBox}),
  dispatch => ({
    onClose: function (e) {
      dispatch(messageBox({show: false, title: '', body: ''}))
    }
  })
)(Widget)

export default translate()(Model)

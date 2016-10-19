import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import {Modal, Button} from 'react-bootstrap'

const Widget = React.createClass({
  handleClose (e) {
    console.log('close')
  },
  render () {
    const {t, info} = this.props
    return (
      <Modal show={info.show} onHide={this.handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Modal heading</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <h4>{info.title}</h4>
          <p>{info.body}</p>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={this.handleClose}>{t('buttons.close')}</Button>
        </Modal.Footer>
      </Modal>
    )
  }
})

Widget.propTypes = {
  t: PropTypes.func.isRequired,
  info: PropTypes.object.isRequired
}

const Model = connect(
  state => ({info: state.messageBox}),
  dispatch => ({})
)(Widget)

export default translate()(Model)

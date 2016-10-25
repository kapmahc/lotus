import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import {Modal,
  Button, ButtonGroup,
  Table
} from 'react-bootstrap'
import TimeAgo from 'react-timeago'

import {toggleSiteNotices, toggleNoticeForm} from './actions'
import {post} from '../../ajax'

export const Index = () => (
  <div>
    notices
  </div>
)

// -----------------------------------------------------------------------------
const ListW = React.createClass({
  render () {
    const {t, notices, onClose, onNew} = this.props
    return (
     <Modal bsSize="large" show={notices.show} onHide={onClose}>
       <Modal.Header closeButton>
         <Modal.Title>{t('auth.dashboard.site-notices')}</Modal.Title>
       </Modal.Header>
       <Modal.Body>
         <Table striped bordered condensed hover>
           <thead>
             <tr>
               <th>{t('attributes.content')}</th>
               <th>{t('attributes.updatedAt')}</th>
               <th>
                 {t('buttons.manage')}
                 [<Button bsSize="xsmall" onClick={onNew} bsStyle="info">{t('buttons.new')}</Button>]
               </th>
             </tr>
           </thead>
           <tbody>
             {notices.items.map(function (n, i) {
               return (
                 <tr key={i}>
                   <td>{n.content}</td>
                   <td><TimeAgo date={n.updated_at}/></td>
                   <td>
                     <ButtonGroup>
                       <Button bsStyle="warning">{t('buttons.edit')}</Button>
                       <Button bsStyle="danger">{t('buttons.remove')}</Button>
                     </ButtonGroup>
                   </td>
                 </tr>
               )
             })}
           </tbody>
         </Table>
       </Modal.Body>
       <Modal.Footer>
         <Button onClick={onClose}>{t('buttons.close')}</Button>
       </Modal.Footer>
     </Modal>
   )
  }
})

ListW.propTypes = {
  t: PropTypes.func.isRequired,
  notices: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired,
  onNew: PropTypes.func.isRequired
}

const ListM = connect(
  state => ({notices: state.adminSiteNotices}),
  dispatch => ({
    onNew: function () {
      dispatch(toggleNoticeForm({}))
    },
    onClose: function () {
      dispatch(toggleSiteNotices())
    }
  })
)(ListW)

export const List = translate()(ListM)

// -----------------------------------------------------------------------------
const FormW = React.createClass({
  getInitialState () {
    return {
      content: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()

    var data = new window.FormData()
    const {notice} = this.props
    if (notice.id) {
      data.append('id', notice.id)
    }
    data.append('content', this.state.content)

    const {onClose} = this.props
    post('/notices', data, function (rst) {
      onClose()
    })
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t, notice, onClose} = this.props
    return (
      <form>
       <Modal show={notice.show} onHide={onClose}>
         <Modal.Header closeButton>
           <Modal.Title>{t('auth.models.notice')}</Modal.Title>
         </Modal.Header>
         <Modal.Body>
           <FormGroup controlId="content">
             <ControlLabel>{t('attributes.content')}</ControlLabel>
             <FormControl
               rows={6}
               componentClass="textarea"
               defaultValue={notice.content}
               onChange={this.handleChange} />
           </FormGroup>
         </Modal.Body>
         <Modal.Footer>
           <Button onClick={this.handleSubmit} bsStyle="primary">{t('buttons.submit')}</Button>
           <Button onClick={onClose}>{t('buttons.close')}</Button>
         </Modal.Footer>
       </Modal>
    </form>
   )
  }
})

FormW.propTypes = {
  t: PropTypes.func.isRequired,
  notice: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const FormM = connect(
  state => ({notice: state.adminNoticeForm}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleNoticeForm())
    }
  })
)(FormW)

export const Form = translate()(FormM)

import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import {Modal,
  Button, ButtonGroup,
  Table
} from 'react-bootstrap'
import TimeAgo from 'react-timeago'

import {toggleSiteNotices} from './actions'
// import {post} from '../../ajax'

export const Index = () => (
  <div>
    notices
  </div>
)

// -----------------------------------------------------------------------------
const ListW = React.createClass({
  render () {
    const {t, notices, onClose} = this.props
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
                 [<Button bsSize="xsmall" bsStyle="info">{t('buttons.new')}</Button>]
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
  onClose: PropTypes.func.isRequired
}

const ListM = connect(
  state => ({notices: state.adminSiteNotices}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleSiteNotices())
    }
  })
)(ListW)

export const List = translate()(ListM)

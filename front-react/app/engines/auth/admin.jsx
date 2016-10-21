import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import {Modal, ListGroup, ListGroupItem,
  Button, FormGroup, ControlLabel, FormControl
} from 'react-bootstrap'

import {toggleSiteStatus, toggleSiteBase, toggleSiteAuthor, toggleSiteNav} from './actions'
import {post} from '../../ajax'

const StatusW = ({t, status, onClose}) => (
  <Modal show={status.show} onHide={onClose}>
    <Modal.Header closeButton>
      <Modal.Title>{t('auth.dashboard.site-status')}</Modal.Title>
    </Modal.Header>
    <Modal.Body>
      <h4>{t('auth.dashboard.site-status-os')}</h4>
      <ListGroup>
        {status.os.map(function (l, i) {
          return <ListGroupItem key={i}>{l}</ListGroupItem>
        })}
      </ListGroup>
      <h4>{t('auth.dashboard.site-status-db')}</h4>
      <h4>{t('auth.dashboard.site-status-cache')}</h4>
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
  state => ({status: state.adminSiteStatus}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleSiteStatus())
    }
  })
)(StatusW)

export const Status = translate()(StatusM)

// -----------------------------------------------------------------------------
const BaseW = React.createClass({
  getInitialState () {
    return {
      title: '',
      subTitle: '',
      keywords: '',
      description: '',
      copyright: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()

    var data = new window.FormData()
    data.append('title', this.state.title)
    data.append('subTitle', this.state.subTitle)
    data.append('keywords', this.state.keywords)
    data.append('description', this.state.description)
    data.append('copyright', this.state.copyright)

    const {onClose} = this.props
    post('/site/base', data, function (rst) {
      onClose()
    })
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t, info, onClose} = this.props
    return (
      <form>
       <Modal show={info.show} onHide={onClose}>
         <Modal.Header closeButton>
           <Modal.Title>{t('auth.dashboard.site-base')}</Modal.Title>
         </Modal.Header>
         <Modal.Body>
           <FormGroup
             controlId="title"
           >
             <ControlLabel>{t('attributes.site.title')}</ControlLabel>
             <FormControl
               defaultValue={info.title}
               type="text"
               onChange={this.handleChange}
             />
           </FormGroup>
           <FormGroup
             controlId="subTitle"
           >
             <ControlLabel>{t('attributes.site.subTitle')}</ControlLabel>
             <FormControl
               defaultValue={info.subTitle}
               type="text"
               onChange={this.handleChange}
             />
           </FormGroup>
           <FormGroup
             controlId="keywords"
           >
             <ControlLabel>{t('attributes.site.keywords')}</ControlLabel>
             <FormControl
               type="text"
               defaultValue={info.keywords}
               onChange={this.handleChange}
             />
           </FormGroup>
           <FormGroup controlId="description">
             <ControlLabel>{t('attributes.site.description')}</ControlLabel>
             <FormControl
               rows={6}
               componentClass="textarea"
               defaultValue={info.description}
               onChange={this.handleChange} />
           </FormGroup>
           <FormGroup
             controlId="copyright"
           >
             <ControlLabel>{t('attributes.site.copyright')}</ControlLabel>
             <FormControl
               type="text"
               defaultValue={info.copyright}
               onChange={this.handleChange}
             />
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

BaseW.propTypes = {
  t: PropTypes.func.isRequired,
  info: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const BaseM = connect(
  state => ({info: state.adminSiteBase}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleSiteBase())
    }
  })
)(BaseW)

export const Base = translate()(BaseM)

// -----------------------------------------------------------------------------
const AuthorW = React.createClass({
  getInitialState () {
    return {
      email: '',
      name: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()

    var data = new window.FormData()
    data.append('email', this.state.email)
    data.append('name', this.state.name)

    const {onClose} = this.props
    post('/site/author', data, function (rst) {
      onClose()
    })
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t, info, onClose} = this.props
    return (
      <form>
       <Modal show={info.show} onHide={onClose}>
         <Modal.Header closeButton>
           <Modal.Title>{t('auth.dashboard.site-author')}</Modal.Title>
         </Modal.Header>
         <Modal.Body>
           <FormGroup
             controlId="name"
           >
             <ControlLabel>{t('attributes.site.author-name')}</ControlLabel>
             <FormControl
               defaultValue={info.name}
               type="text"
               onChange={this.handleChange}
             />
           </FormGroup>
           <FormGroup
             controlId="email"
           >
             <ControlLabel>{t('attributes.site.author-email')}</ControlLabel>
             <FormControl
               defaultValue={info.email}
               type="email"
               onChange={this.handleChange}
             />
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

AuthorW.propTypes = {
  t: PropTypes.func.isRequired,
  info: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const AuthorM = connect(
  state => ({info: state.adminSiteAuthor}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleSiteAuthor())
    }
  })
)(AuthorW)

export const Author = translate()(AuthorM)

// -----------------------------------------------------------------------------
const NavW = React.createClass({
  getInitialState () {
    return {
      top: '',
      bottom: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()

    var data = new window.FormData()
    data.append('top', this.state.top)
    data.append('bottom', this.state.bottom)

    const {onClose} = this.props
    post('/site/nav', data, function (rst) {
      onClose()
    })
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t, info, onClose} = this.props
    const links2str = function (links) {
      return links ? links.map(function (l) {
        return `${l.href} = ${l.label}`
      }).join('\n') : ''
    }
    return (
      <form>
       <Modal show={info.show} onHide={onClose}>
         <Modal.Header closeButton>
           <Modal.Title>{t('auth.dashboard.site-nav')}</Modal.Title>
         </Modal.Header>
         <Modal.Body>
           <FormGroup controlId="top">
             <ControlLabel>{t('attributes.site.top-links')}</ControlLabel>
             <FormControl
               rows={6}
               componentClass="textarea"
               defaultValue={links2str(info.top)}
               onChange={this.handleChange} />
           </FormGroup>
           <FormGroup controlId="bottom">
             <ControlLabel>{t('attributes.site.bottom-links')}</ControlLabel>
             <FormControl
               rows={6}
               componentClass="textarea"
               defaultValue={links2str(info.bottom)}
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

NavW.propTypes = {
  t: PropTypes.func.isRequired,
  info: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const NavM = connect(
  state => ({info: state.adminSiteNav}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleSiteNav())
    }
  })
)(NavW)

export const Nav = translate()(NavM)

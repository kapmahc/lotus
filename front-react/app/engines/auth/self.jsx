import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import TimeAgo from 'react-timeago'
import {Modal, ListGroup, ListGroupItem,
  HelpBlock,
  Button, FormGroup, ControlLabel, FormControl
} from 'react-bootstrap'

import {toggleUserLogs, hideUserProfile, toggleUserPassword} from './actions'
import {post} from '../../ajax'

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
      dispatch(toggleUserLogs())
    }
  })
)(LogsW)

export const Logs = translate()(LogsM)

// -----------------------------------------------------------------------------
const ProfileW = React.createClass({
  getInitialState () {
    return {
      email: '',
      name: '',
      home: '',
      logo: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()

    var data = new window.FormData()
    data.append('name', this.state.name)
    data.append('home', this.state.home)
    data.append('logo', this.state.logo)

    const {onClose} = this.props
    post('/self/profile', data, function (rst) {
      onClose()
    })
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t, profile, onClose} = this.props
    return (
      <form>
       <Modal show={profile.show} onHide={onClose}>
         <Modal.Header closeButton>
           <Modal.Title>{t('auth.dashboard.self-profile')}</Modal.Title>
         </Modal.Header>
         <Modal.Body>
           <FormGroup
             controlId="email"
           >
             <ControlLabel>{t('attributes.user.email')}</ControlLabel>
             <FormControl
               defaultValue={profile.email}
               type="email"
               readOnly
               onChange={this.handleChange}
             />
           </FormGroup>
           <FormGroup
             controlId="name"
           >
             <ControlLabel>{t('attributes.user.name')}</ControlLabel>
             <FormControl
               defaultValue={profile.name}
               type="text"
               onChange={this.handleChange}
             />
           </FormGroup>
           <FormGroup
             controlId="home"
           >
             <ControlLabel>{t('attributes.user.home')}</ControlLabel>
             <FormControl
               type="text"
               defaultValue={profile.home}
               onChange={this.handleChange}
             />
           </FormGroup>
           <FormGroup
             controlId="logo"
           >
             <ControlLabel>{t('attributes.user.logo')}</ControlLabel>
             <FormControl
               type="text"
               defaultValue={profile.logo}
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

ProfileW.propTypes = {
  t: PropTypes.func.isRequired,
  profile: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const ProfileM = connect(
  state => ({profile: state.userProfile}),
  dispatch => ({
    onClose: function () {
      dispatch(hideUserProfile())
    }
  })
)(ProfileW)

export const Profile = translate()(ProfileM)

// -----------------------------------------------------------------------------
const PasswordW = React.createClass({
  getInitialState () {
    return {
      currentPassword: '',
      password: '',
      passwordConfirmation: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()

    var data = new window.FormData()
    data.append('currentPassword', this.state.currentPassword)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)

    const {onClose} = this.props
    post('/self/password', data, function (rst) {
      this.setState({
        currentPassword: '',
        password: '',
        passwordConfirmation: ''
      })
      onClose()
    }.bind(this))
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t, passwords, onClose} = this.props
    return (
      <form>
       <Modal show={passwords.show} onHide={onClose}>
         <Modal.Header closeButton>
           <Modal.Title>{t('auth.dashboard.self-password')}</Modal.Title>
         </Modal.Header>
         <Modal.Body>
           <FormGroup
             controlId="currentPassword"
           >
             <ControlLabel>{t('attributes.user.currentPassword')}</ControlLabel>
             <FormControl
               type="password"
               value={this.state.currentPassword}
               onChange={this.handleChange}
             />
           <HelpBlock>{t('auth.we-need-your-current-password')}</HelpBlock>
           </FormGroup>
           <FormGroup
             controlId="password"
           >
             <ControlLabel>{t('attributes.user.password')}</ControlLabel>
             <FormControl
               type="password"
               value={this.state.password}
               onChange={this.handleChange}
             />
             <HelpBlock>{t('auth.password-must-in-size')}</HelpBlock>
           </FormGroup>
           <FormGroup
             controlId="passwordConfirmation"
           >
             <ControlLabel>
               {t('attributes.user.passwordConfirmation')}
             </ControlLabel>
             <FormControl
               type="password"
               value={this.state.passwordConfirmation}
               onChange={this.handleChange}
             />
             <HelpBlock>{t('auth.passwords-must-match')}</HelpBlock>
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

PasswordW.propTypes = {
  t: PropTypes.func.isRequired,
  passwords: PropTypes.object.isRequired,
  onClose: PropTypes.func.isRequired
}

const PasswordM = connect(
  state => ({passwords: state.userPassword}),
  dispatch => ({
    onClose: function () {
      dispatch(toggleUserPassword())
    }
  })
)(PasswordW)

export const Password = translate()(PasswordM)

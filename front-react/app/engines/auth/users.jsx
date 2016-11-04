import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import { Link, browserHistory } from 'react-router'
import {FormGroup, FormControl, HelpBlock,
  Button, ControlLabel} from 'react-bootstrap'

import {post} from '../../ajax'
import {messageBox, signIn} from './actions'

const SignInW = React.createClass({
  getInitialState () {
    return {
      email: '',
      password: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()

    var data = new window.FormData()
    data.append('email', this.state.email)
    data.append('password', this.state.password)

    const {onSignIn} = this.props
    post('/users/sign-in', data, function (rst) {
      this.setState({password: ''})
      onSignIn(rst.token)
      browserHistory.push('/')
    }.bind(this))
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t} = this.props
    return (
      <fieldset>
        <legend>{t('auth.pages.sign-in')}</legend>
        <form>
          <FormGroup
            controlId="email"
          >
            <ControlLabel>{t('attributes.email')}</ControlLabel>
            <FormControl
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
          </FormGroup>
          <FormGroup
            controlId="password"
          >
            <ControlLabel>{t('attributes.password')}</ControlLabel>
            <FormControl
              type="password"
              value={this.state.password}
              onChange={this.handleChange}
            />
          </FormGroup>
          <Button onClick={this.handleSubmit} type="submit">
            {t('buttons.submit')}
          </Button>
        </form>
        <br/>
        <SharedLinks/>
      </fieldset>
    )
  }
})

SignInW.propTypes = {
  user: PropTypes.object.isRequired,
  t: PropTypes.func.isRequired,
  onSignIn: PropTypes.func.isRequired
}

const SignInM = connect(
  state => ({user: state.currentUser}),
  dispatch => ({
    onSignIn: function (token) {
      dispatch(signIn(token))
    }
  })
)(SignInW)

export const SignIn = translate()(SignInM)

// -----------------------------------------------------------------------------

const SignUpW = React.createClass({
  getInitialState () {
    return {
      name: '',
      email: '',
      password: '',
      passwordConfirmation: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()
    const {t, showBox} = this.props

    var data = new window.FormData()
    data.append('email', this.state.email)
    data.append('name', this.state.name)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)

    post('/users/sign-up', data, function (user) {
      this.setState({password: '', passwordConfirmation: ''})
      showBox({show: true, title: t('success'), body: t('auth.confirm-success')})
    }.bind(this))
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t} = this.props
    return (
      <fieldset>
        <legend>{t('auth.pages.sign-up')}</legend>
        <form>
          <FormGroup
            controlId="name"
          >
            <ControlLabel>{t('attributes.username')}</ControlLabel>
            <FormControl
              type="text"
              value={this.state.name}
              onChange={this.handleChange}
            />
          </FormGroup>
          <FormGroup
            controlId="email"
          >
            <ControlLabel>{t('attributes.email')}</ControlLabel>
            <FormControl
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
          </FormGroup>
          <FormGroup
            controlId="password"
          >
            <ControlLabel>{t('attributes.password')}</ControlLabel>
            <FormControl
              type="password"
              value={this.state.password}
              onChange={this.handleChange}
            />
          <HelpBlock>{t('auth.helpers.password-must-in-size')}</HelpBlock>
          </FormGroup>
          <FormGroup
            controlId="passwordConfirmation"
          >
            <ControlLabel>
              {t('attributes.passwordConfirmation')}
            </ControlLabel>
            <FormControl
              type="password"
              value={this.state.passwordConfirmation}
              onChange={this.handleChange}
            />
          <HelpBlock>{t('helpers.passwords-must-match')}</HelpBlock>
          </FormGroup>
          <Button onClick={this.handleSubmit} type="submit">
            {t('buttons.submit')}
          </Button>
        </form>
        <br/>
        <SharedLinks/>
      </fieldset>
    )
  }
})

SignUpW.propTypes = {
  t: PropTypes.func.isRequired,
  showBox: PropTypes.func.isRequired
}

const SignUpM = connect(
  state => ({}),
  dispatch => ({
    showBox: function (info) {
      dispatch(messageBox(info))
    }
  })
)(SignUpW)

export const SignUp = translate()(SignUpM)
// -----------------------------------------------------------------------------
export const ForgotPassword = () => (
  <EmailForm action="forgot-password" />
)
// -----------------------------------------------------------------------------

const ChangePasswordW = React.createClass({
  getInitialState () {
    return {
      password: '',
      passwordConfirmation: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()
    const {t, showBox, location} = this.props

    var data = new window.FormData()
    data.append('token', location.query.token)
    data.append('password', this.state.password)
    data.append('passwordConfirmation', this.state.passwordConfirmation)

    post('/users/change-password', data, function (user) {
      this.setState({password: '', passwordConfirmation: ''})
      showBox({show: true, title: t('success'), body: t('auth.change-password-success')})
    }.bind(this))
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t} = this.props
    return (
      <fieldset>
        <legend>{t('auth.pages.change-password')}</legend>
        <form>
          <FormGroup
            controlId="password"
          >
            <ControlLabel>{t('attributes.password')}</ControlLabel>
            <FormControl
              type="password"
              value={this.state.password}
              onChange={this.handleChange}
            />
          <HelpBlock>{t('auth.helpers.password-must-in-size')}</HelpBlock>
          </FormGroup>
          <FormGroup
            controlId="passwordConfirmation"
          >
            <ControlLabel>
              {t('attributes.passwordConfirmation')}
            </ControlLabel>
            <FormControl
              type="password"
              value={this.state.passwordConfirmation}
              onChange={this.handleChange}
            />
          <HelpBlock>{t('helpers.passwords-must-match')}</HelpBlock>
          </FormGroup>
          <Button onClick={this.handleSubmit} type="submit">
            {t('buttons.submit')}
          </Button>
        </form>
        <br/>
        <SharedLinks/>
      </fieldset>
    )
  }
})

ChangePasswordW.propTypes = {
  t: PropTypes.func.isRequired,
  showBox: PropTypes.func.isRequired
}

const ChangePasswordM = connect(
  state => ({}),
  dispatch => ({
    showBox: function (info) {
      dispatch(messageBox(info))
    }
  })
)(ChangePasswordW)

export const ChangePassword = translate()(ChangePasswordM)

// -----------------------------------------------------------------------------
export const Confirm = () => (
  <EmailForm action="confirm" />
)
// -----------------------------------------------------------------------------
export const Unlock = () => (
  <EmailForm action="unlock" />
)
// -----------------------------------------------------------------------------
const SharedLinksW = ({t}) => (
  <ul>
    <li>
      <Link to="/users/sign-in">{t('auth.pages.sign-in')}</Link>
    </li>
    <li>
      <Link to="/users/sign-up">{t('auth.pages.sign-up')}</Link>
    </li>
    <li>
      <Link to="/users/forgot-password">{t('auth.pages.forgot-password')}</Link>
    </li>
    <li>
      <Link to="/users/confirm">{t('auth.pages.confirm')}</Link>
    </li>
    <li>
      <Link to="/users/unlock">{t('auth.pages.unlock')}</Link>
    </li>
  </ul>
)

SharedLinksW.propTypes = {
  t: PropTypes.func.isRequired
}

const SharedLinks = translate()(SharedLinksW)

// -----------------------------------------------------------------------------
const EmailFormW = React.createClass({
  getInitialState () {
    return {
      email: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()
    const {t, showBox, action} = this.props

    var data = new window.FormData()
    data.append('email', this.state.email)

    post(`/users/${action}`, data, function (user) {
      this.setState({password: '', passwordConfirmation: ''})
      showBox({show: true, title: t('success'), body: t(`auth.pages.${action}-success`)})
    }.bind(this))
  },
  handleChange (e) {
    var o = {}
    o[e.target.id] = e.target.value
    this.setState(o)
  },
  render () {
    const {t, action} = this.props
    return (
      <fieldset>
        <legend>{t(`auth.pages.${action}`)}</legend>
        <form>
          <FormGroup
            controlId="email"
          >
            <ControlLabel>{t('attributes.email')}</ControlLabel>
            <FormControl
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
          </FormGroup>
          <Button onClick={this.handleSubmit} type="submit">
            {t('buttons.submit')}
          </Button>
        </form>
        <br/>
        <SharedLinks/>
      </fieldset>
    )
  }
})

EmailFormW.propTypes = {
  t: PropTypes.func.isRequired,
  action: PropTypes.string.isRequired,
  showBox: PropTypes.func.isRequired
}

const EmailFormM = connect(
  state => ({}),
  dispatch => ({
    showBox: function (info) {
      dispatch(messageBox(info))
    }
  })
)(EmailFormW)

const EmailForm = translate()(EmailFormM)

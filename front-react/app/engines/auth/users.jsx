import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'
import { Link } from 'react-router'
import {FormGroup, FormControl, HelpBlock,
  Button, ControlLabel} from 'react-bootstrap'

import {post} from '../../ajax'
import {messageBox} from './actions'

const SignInW = React.createClass({
  getInitialState () {
    return {
      email: '',
      password: ''
    }
  },
  handleSubmit (e) {
    e.preventDefault()
    console.log(this.state)
    // TODO
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
        <legend>{t('auth.sign-in')}</legend>
        <form>
          <FormGroup
            controlId="email"
          >
            <ControlLabel>{t('attributes.user.email')}</ControlLabel>
            <FormControl
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
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
  t: PropTypes.func.isRequired
}

export const SignIn = translate()(SignInW)

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

    post('/users/signUp', data, function (user) {
      this.setState({password: '', passwordConfirmation: ''})
      showBox({show: true, title: t('success'), body: t('auth.sign-up-success')})
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
        <legend>{t('auth.sign-up')}</legend>
        <form>
          <FormGroup
            controlId="name"
          >
            <ControlLabel>{t('attributes.user.name')}</ControlLabel>
            <FormControl
              type="text"
              value={this.state.name}
              onChange={this.handleChange}
            />
          </FormGroup>
          <FormGroup
            controlId="email"
          >
            <ControlLabel>{t('attributes.user.email')}</ControlLabel>
            <FormControl
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
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
  state => ({info: state.siteInfo}),
  dispatch => ({
    showBox: function (info) {
      dispatch(messageBox(info))
    }
  })
)(SignUpW)

export const SignUp = translate()(SignUpM)
// -----------------------------------------------------------------------------
export const ForgotPassword = () => (
  <div> forgot password </div>
)
// -----------------------------------------------------------------------------
export const ChangePassword = () => (
  <div> change password </div>
)
// -----------------------------------------------------------------------------
export const Confirm = () => (
  <div> confirm </div>
)
// -----------------------------------------------------------------------------
export const Unlock = () => (
  <div> unlock </div>
)
// -----------------------------------------------------------------------------
const SharedLinksW = ({t}) => (
  <ul>
    <li>
      <Link to="/users/sign-in">{t('auth.sign-in')}</Link>
    </li>
    <li>
      <Link to="/users/sign-up">{t('auth.sign-up')}</Link>
    </li>
  </ul>
)

SharedLinksW.propTypes = {
  t: PropTypes.func.isRequired
}

const SharedLinks = translate()(SharedLinksW)
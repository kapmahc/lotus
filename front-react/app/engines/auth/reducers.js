import jwtDecode from 'jwt-decode'

import {SIGN_IN, SIGN_OUT,
  MESSAGE_BOX,
  SHOW_USER_LOGS, HIDE_USER_LOGS,
  SHOW_USER_PROFILE, HIDE_USER_PROFILE,
  TOGGLE_USER_PROFILE,
  REFRESH} from './actions'

const key = 'token'

const userPassword = (state = {show: false}, action) => {
  switch (action.type) {
    case TOGGLE_USER_PROFILE:
      return {show: !state.show}
    default:
      return state
  }
}

const userProfile = (state = {}, action) => {
  switch (action.type) {
    case SHOW_USER_PROFILE:
      return Object.assign({show: true}, action.info)
    case HIDE_USER_PROFILE:
      return {show: false}
    default:
      return state
  }
}

const userLogs = (state = {items: []}, action) => {
  switch (action.type) {
    case SHOW_USER_LOGS:
      return {show: true, items: action.logs}
    case HIDE_USER_LOGS:
      return {show: false, items: []}
    default:
      return state
  }
}

const currentUser = (state = {}, action) => {
  switch (action.type) {
    case SIGN_IN:
      try {
        var user = jwtDecode(action.token)
        window.sessionStorage.setItem(key, action.token)
        return user
      } catch (e) {
        console.log(e)
        return {}
      }
    case SIGN_OUT:
      window.sessionStorage.removeItem(key)
      return {}
    default:
      return state
  }
}

const siteInfo = (state = {bottomLinks: [], topLinks: [], author: {}}, action) => {
  switch (action.type) {
    case REFRESH:
      return action.info
    default:
      return state
  }
}

const messageBox = (state = {}, action) => {
  switch (action.type) {
    case MESSAGE_BOX:
      return action.info
    default:
      return state
  }
}

const reducers = {currentUser, siteInfo, messageBox, userLogs, userProfile, userPassword}
export default reducers

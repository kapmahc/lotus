import jwtDecode from 'jwt-decode'

import {SIGN_IN, SIGN_OUT,
  MESSAGE_BOX,
  SHOW_USER_LOGS, HIDE_USER_LOGS,
  REFRESH} from './actions'

const key = 'token'

const userLogs = (state = {items: []}, action) => {
  switch (action.type) {
    case SHOW_USER_LOGS:
      return {show: true, items: action.logs}
    case HIDE_USER_LOGS:
      return {show: false}
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

const reducers = {currentUser, siteInfo, messageBox, userLogs}
export default reducers

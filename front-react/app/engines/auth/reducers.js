import jwtDecode from 'jwt-decode'

import {SIGN_IN, SIGN_OUT,
  MESSAGE_BOX,
  REFRESH} from './actions'

const key = 'token'

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

const reducers = {currentUser, siteInfo, messageBox}
export default reducers

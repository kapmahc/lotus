import {SIGN_IN, SIGN_OUT, REFRESH} from './actions'

const key = 'token'

const currentUser = (state = {}, action) => {
  switch (action.type) {
    case SIGN_IN:
      console.log(action)
      // TODO parse action
      window.sessionStorage.setItem(key, action.token)
      return {
        name: null,
        uid: null,
        roles: []
      }
    case SIGN_OUT:
      window.sessionStorage.removeItem(key)
      return {}
    default:
      return state
  }
}

const siteInfo = (state = {}, action) => {
  switch (action.type) {
    case REFRESH:
      return action.info
    default:
      return state
  }
}

const reducers = {currentUser, siteInfo}
export default reducers

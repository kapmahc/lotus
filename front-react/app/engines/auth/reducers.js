import jwtDecode from 'jwt-decode'

import {SIGN_IN, SIGN_OUT,
  MESSAGE_BOX,

  TOGGLE_USER_LOGS,
  TOGGLE_USER_PROFILE,
  TOGGLE_USER_PASSWORD,

  TOGGLE_SITE_STATUS,
  TOGGLE_SITE_BASE,
  TOGGLE_SITE_NAV,
  TOGGLE_SITE_AUTHOR,
  TOGGLE_SITE_SEO,
  TOGGLE_SITE_USERS,
  TOGGLE_SITE_NOTICES,

  REFRESH} from './actions'

const key = 'token'

const adminSiteNotices = (state = {items: []}, action) => {
  switch (action.type) {
    case TOGGLE_SITE_NOTICES:
      return action.notices ? {show: true, items: action.notices} : {items: []}
    default:
      return state
  }
}

const adminSiteUsers = (state = {items: []}, action) => {
  switch (action.type) {
    case TOGGLE_SITE_USERS:
      return action.users ? {show: true, items: action.users} : {items: []}
    default:
      return state
  }
}

const adminSiteSeo = (state = {os: []}, action) => {
  switch (action.type) {
    case TOGGLE_SITE_SEO:
      return action.info ? Object.assign({show: true}, action.info) : {show: false}
    default:
      return state
  }
}

const adminSiteNav = (state = {os: []}, action) => {
  switch (action.type) {
    case TOGGLE_SITE_NAV:
      return action.info ? Object.assign({show: true}, action.info) : {show: false}
    default:
      return state
  }
}

const adminSiteAuthor = (state = {os: []}, action) => {
  switch (action.type) {
    case TOGGLE_SITE_AUTHOR:
      return action.info ? Object.assign({show: true}, action.info) : {show: false}
    default:
      return state
  }
}

const adminSiteBase = (state = {os: []}, action) => {
  switch (action.type) {
    case TOGGLE_SITE_BASE:
      return action.info ? Object.assign({show: true}, action.info) : {show: false}
    default:
      return state
  }
}

const adminSiteStatus = (state = {os: []}, action) => {
  switch (action.type) {
    case TOGGLE_SITE_STATUS:
      return action.status ? Object.assign({show: true}, action.status) : {os: []}
    default:
      return state
  }
}

const userPassword = (state = {show: false}, action) => {
  switch (action.type) {
    case TOGGLE_USER_PASSWORD:
      return {show: !state.show}
    default:
      return state
  }
}

const userProfile = (state = {}, action) => {
  switch (action.type) {
    case TOGGLE_USER_PROFILE:
      return action.info ? Object.assign({show: true}, action.info) : {show: false}
    default:
      return state
  }
}

const userLogs = (state = {items: []}, action) => {
  switch (action.type) {
    case TOGGLE_USER_LOGS:
      return action.logs ? {show: true, items: action.logs} : {show: false, items: []}
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

const reducers = {
  currentUser,
  siteInfo,
  messageBox,

  userLogs,
  userProfile,
  userPassword,

  adminSiteStatus,
  adminSiteNav,
  adminSiteBase,
  adminSiteAuthor,
  adminSiteSeo,
  adminSiteNotices,
  adminSiteUsers
}
export default reducers

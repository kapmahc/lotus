export const SIGN_IN = 'auth.sign-in'
export const SIGN_OUT = 'auth.sign-out'
export const REFRESH = 'auth.refresh'
export const MESSAGE_BOX = 'auth.message-box'

export const TOGGLE_USER_LOGS = 'auth.user.logs.toggle'
export const TOGGLE_USER_PROFILE = 'auth.user.profile.toggle'
export const TOGGLE_USER_PASSWORD = 'auth.user.password.toggle'

export const TOGGLE_SITE_STATUS = 'auth.site.status.toggle'
export const TOGGLE_SITE_BASE = 'auth.site.base.toggle'
export const TOGGLE_SITE_AUTHOR = 'auth.site.author.toggle'
export const TOGGLE_SITE_NAV = 'auth.site.nav.toggle'
export const TOGGLE_SITE_SEO = 'auth.site.seo.toggle'
export const TOGGLE_SITE_USERS = 'auth.site.users.toggle'
export const TOGGLE_SITE_NOTICES = 'auth.site.notices.toggle'

export const toggleSiteNotices = (notices) => {
  return {
    type: TOGGLE_SITE_NOTICES,
    notices
  }
}

export const toggleSiteUsers = (users) => {
  return {
    type: TOGGLE_SITE_USERS,
    users
  }
}

export const toggleSiteSeo = (info) => {
  return {
    type: TOGGLE_SITE_SEO,
    info
  }
}

export const toggleSiteAuthor = (info) => {
  return {
    type: TOGGLE_SITE_AUTHOR,
    info
  }
}

export const toggleSiteNav = (info) => {
  return {
    type: TOGGLE_SITE_NAV,
    info
  }
}

export const toggleSiteBase = (info) => {
  return {
    type: TOGGLE_SITE_BASE,
    info
  }
}

export const toggleSiteStatus = (status) => {
  return {
    type: TOGGLE_SITE_STATUS,
    status
  }
}

export const toggleUserPassword = () => {
  return {
    type: TOGGLE_USER_PASSWORD
  }
}

export const toggleUserProfile = (info) => {
  return {
    type: TOGGLE_USER_PROFILE,
    info
  }
}

export const toggleUserLogs = (logs) => {
  return {
    type: TOGGLE_USER_LOGS,
    logs
  }
}

export const refresh = (info) => {
  return {
    type: REFRESH,
    info
  }
}

export const signIn = (token) => {
  return {
    type: SIGN_IN,
    token
  }
}

export const signOut = () => {
  return {type: SIGN_OUT}
}

export const messageBox = (info) => {
  return {type: MESSAGE_BOX, info}
}

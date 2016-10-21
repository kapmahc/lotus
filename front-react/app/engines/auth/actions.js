export const SIGN_IN = 'auth.sign-in'
export const SIGN_OUT = 'auth.sign-out'
export const REFRESH = 'auth.refresh'
export const MESSAGE_BOX = 'auth.message-box'
export const SHOW_USER_LOGS = 'auth.user.logs.show'
export const HIDE_USER_LOGS = 'auth.user.logs.hide'

export const showUserLogs = (logs) => {
  return {
    type: SHOW_USER_LOGS,
    logs
  }
}

export const hideUserLogs = () => {
  return {
    type: HIDE_USER_LOGS
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

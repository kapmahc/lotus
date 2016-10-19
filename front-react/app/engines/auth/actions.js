export const SIGN_IN = 'auth.sign_in'
export const SIGN_OUT = 'auth.sign_out'
export const REFRESH = 'auth.refresh'

export const refresh = (info) => {
  return {
    type: REFRESH,
    info
  }
}

export const signIn = (token) => {
  console.log(token)
  return {
    type: SIGN_IN,
    token
  }
}

export const signOut = () => {
  return {type: SIGN_OUT}
}

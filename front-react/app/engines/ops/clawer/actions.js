export const SIGN_IN = 'ops.clawer.line.sign-in'
export const SIGN_OUT = 'ops.clawer.line.sign-in'

export const signIn = (user) => {
  return {
    type: SIGN_IN,
    user: user
  }
}

export const signOut = () => {
  return {
    type: SIGN_OUT
  }
}

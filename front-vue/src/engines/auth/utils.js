export function isSignIn (user) {
  return user && user.uid
}

export function isAdmin (user) {
  return user && user.roles && user.roles.includes('admin')
}

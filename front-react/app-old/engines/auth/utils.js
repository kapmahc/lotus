export function isEmpty (obj) {
  return Object.keys(obj).length === 0 && obj.constructor === Object
}

export function has (user, role) {
  return !isEmpty(user) && user.roles && user.roles.indexOf(role) !== -1
}

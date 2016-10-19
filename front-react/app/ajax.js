export function get (url, success, fail) {
  call('get', url, null, success, fail)
}

export function post (url, data, success, fail) {
  call('post', url, data, success, fail)
}

export function patch () {
  // TODO
}

export function _delete () {
  // TODO
}

function call (method, url, data, success, fail) {
  if (success == null) {
    success = function (rst) {
      console.log(rst)
    }
  }
  if (fail == null) {
    fail = function (err) {
      console.log('todo')
      console.log(err)
    }
  }
  window.fetch(`${CONFIG.backend}${url}`, {method: method, data: data})
  .then(function (response) {
    return response.json()
  })
  .then(success)
  .catch(fail)
}

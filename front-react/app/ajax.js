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

function call (method, url, body, success, fail) {
  if (success == null) {
    success = function (rst) {
      console.log(rst)
    }
  }
  if (fail == null) {
    fail = function (err) {
      window.alert(err)
    }
  }
  window.fetch(
    `${CONFIG.backend}${url}`,
    {
      method: method,
      body: body,
      mode: 'cors'
    }
  )
  .then(function (response) {
    if (response.ok) {
      var contentType = response.headers.get('content-type')
      if (contentType && contentType.indexOf('application/json') !== -1) {
        return response.json().then(success).catch(fail)
      }
      console.log(response)
    } else {
      response.text().then(fail)
    }
  })
  .catch(fail)
}

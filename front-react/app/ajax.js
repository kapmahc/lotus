export function get (url, success, fail) {
  call('get', url, null, success, fail)
}

export function post (url, data, success, fail) {
  call('post', url, data, success, fail)
}

export function patch (url, data, success, fail) {
  call('patch', url, data, success, fail)
}

export function _delete (url, success, fail) {
  call('delete', url, null, success, fail)
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

  var headers = new window.Headers()
  headers.append('Authorization', `Bearer ${window.sessionStorage.getItem('token')}`)

  window.fetch(
    `${CONFIG.backend}${url}`,
    {
      method: method,
      headers: headers,
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

export function get (url, data, success, fail) {
  call('get', url, data, success, fail)
}

export function post () {
}

export function patch () {
}

// export function delete () {
//
// }

function call (method, url, data, success, fail) {
  if (success == null) {
    success = function (rst) {
      console.log(rst)
    }
  }
  if (fail == null) {
    fail = function (err) {
      console.log(err)
    }
  }
  window.fetch(`${CONFIG.backend}${url}`, {method: method})
  .then(function (response) {
    return response.json()
  })
  .then(success)
  .catch(fail)
}

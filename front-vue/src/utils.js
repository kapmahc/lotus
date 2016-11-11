import $ from 'jquery'

export const api = url => `${process.env.API_HOST}${url}`

export const postForm = (url, data, success, fail) => {
  if (success == null) {
    success = function (result) {
      console.log(result)
    }
  }
  if (fail == null) {
    fail = function (err) {
      window.alert(`${err.statusText}:\n ${err.responseText}`)
    }
  }

  $.ajax({
    type: 'POST',
    url: api(url),
    data: data
  }).done(success).fail(fail)
  // window.fetch(
  //   api(url),
  //   {
  //     method: 'POST',
  //     mode: 'cors',
  //     body: body
  //   }).then(
  //     function (res) { return res.json() }
  //   ).then(
  //     function (result) {
  //       console.log(result)
  //     }).catch(function (err) { window.alert(err) })
}

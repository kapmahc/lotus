import $ from 'jquery'
import {currentLocale} from './i18n'

export const api = url => `${process.env.API_HOST}${url}`

export const get = (url, data, success, fail) => {
  ajax('GET', url, data, success, fail)
}

export const _delete = (url, success, fail) => {
  ajax('DELETE', url, null, success, fail)
}

export const postForm = (url, data, success, fail) => {
  ajax('POST', url, data, success, fail)
}

function ajax (method, url, data, success, fail) {
  if (data == null) {
    data = {}
  }
  data.locale = currentLocale()
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
    type: method,
    url: api(url),
    data: data,
    beforeSend: function (xhr) {
      xhr.setRequestHeader(
        'Authorization',
        `Bearer ${window.sessionStorage.getItem('token')}`
      )
    }
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

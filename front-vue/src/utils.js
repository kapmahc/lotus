export const api = url => `${process.env.API_HOST}${url}`

export const postForm = (url, body, success) => {
  window.fetch(
    api(url),
    {
      method: 'POST',
      body: body
    }).then(
      function (res) { return res.json() }
    ).then(
      function (result) {
        console.log(result)
      }).catch(function (err) { window.alert(err) })
}

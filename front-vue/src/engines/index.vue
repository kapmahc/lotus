<script>
import auth from './auth'

import Home from './Home'

const engines = {
  auth
}

export default {
  routes () {
    var items = Object.keys(engines).reduce(function (obj, en) {
      return obj.concat(engines[en].routes)
    }, [])
    items.push({path: '/', component: Home})
    return items
  }
}
</script>

import Home from './Home'
import Dashboard from './Dashboard'

import auth from './auth/routes'
import shop from './shop/routes'
import forum from './forum/routes'

const routes = [
  {name: 'home', path: '/', component: Home},
  {name: 'dashboard', path: '/dashboard', component: Dashboard}
].concat(auth).concat(shop).concat(forum)

export default routes

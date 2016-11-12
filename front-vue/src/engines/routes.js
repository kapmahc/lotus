import auth from './auth/routes'
import Home from './Home'
import Dashboard from './Dashboard'

const routes = [
  {name: 'home', path: '/', component: Home},
  {name: 'dashboard', path: '/dashboard', component: Dashboard}
].concat(auth)

export default routes

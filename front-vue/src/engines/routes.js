import auth from './auth/routes'
import Home from './Home'

const routes = [{name: 'home', path: '/', component: Home}].concat(auth)

export default routes

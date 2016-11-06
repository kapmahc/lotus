import auth from './auth/routes'
import Home from './Home'

const routes = [{path: '/', component: Home}].concat(auth)

export default routes

import reducers from './reducers'
import routes from './routes'
import {Index as Notices} from './notices'

const engine = {
  reducers: reducers,
  routes: routes,
  home: Notices
}

export default engine

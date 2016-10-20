import React from 'react'

import reducers from './reducers'
import routes from './routes'
import {Index as Notices} from './notices'
import Dashboard from './Dashboard'

const engine = {
  reducers: reducers,
  routes: routes,
  home: Notices,
  dashboard: <Dashboard key="auth.dashboard"/>
}

export default engine

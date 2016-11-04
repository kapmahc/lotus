import React from 'react'
import { render } from 'react-dom'
import { I18nextProvider } from 'react-i18next'
import { createStore, combineReducers } from 'redux'
import { Provider } from 'react-redux'
import { Router, Route, browserHistory } from 'react-router'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

import i18n from './i18n'
import Layout from './components/Layout'
import NoMatch from './components/NoMatch'
import reducers from './engines/reducers'

import OpsClawerRouter from './engines/ops/clawer/Router'

console.log('react version: ' + React.version)
console.log('lotus version: ' + CONFIG.version)

const store = createStore(
  combineReducers({
    ...reducers,
    routing: routerReducer
  })
)

const history = syncHistoryWithStore(browserHistory, store)

export default function (id) {
  render(
    <Provider store={store}>
      <I18nextProvider i18n={ i18n }>
        <Router history={history}>
          <Route path="/" component={Layout}>
            {OpsClawerRouter}
            <Route path="*" component={NoMatch}/>
          </Route>
        </Router>
      </I18nextProvider>
    </Provider>,
    document.getElementById(id)
  )
}

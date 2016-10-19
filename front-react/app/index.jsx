import React from 'react'
import { render } from 'react-dom'
import { I18nextProvider } from 'react-i18next'
import { Provider } from 'react-redux'
import { createStore } from 'redux'

import i18n from './i18n'
import Layout from './components/Layout'
import engines from './engines'

let store = createStore(engines.Root)

render(
  (
    <Provider store={store}>
      <I18nextProvider i18n={ i18n }>
        <Layout />
      </I18nextProvider>
    </Provider>
  ),
  document.getElementById('root')
)

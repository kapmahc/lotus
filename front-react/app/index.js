import React from 'react'
import ReactDom from 'react-dom'
import { I18nextProvider } from 'react-i18next'

import i18n from './i18n'
import Layout from './components/Layout'

ReactDom.render(
  (<I18nextProvider i18n={ i18n }><Layout /></I18nextProvider>),
  document.getElementById('root')
)

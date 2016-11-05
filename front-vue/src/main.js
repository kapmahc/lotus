import Vue from 'vue'
import Layout from './components/Layout'

require('./assets/main.css')
require('./assets/reading.css')

/* eslint-disable no-new */
new Vue({
  el: '#root',
  template: '<Layout/>',
  components: { Layout }
})

import Vue from 'vue'
import Home from './components/Home'

require('./assets/main.css')
require('./assets/reading.css')

/* eslint-disable no-new */
new Vue({
  el: '#root',
  template: '<Home/>',
  components: { Home }
})

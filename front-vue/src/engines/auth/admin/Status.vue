<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-status') }}</h3>
    <hr/>
    <pre>
      <code>
        {{JSON.stringify(status)}}
      </code>
    </pre>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {get} from '../../../utils'

export default {
  name: 'admin-status',
  data () {
    return {
      status: {}
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/admin/status', null, function (rst) {
      this.status = rst
    }.bind(this))
  },
  methods: {
  }
}
</script>

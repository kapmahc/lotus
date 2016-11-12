<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.logs') }}</h3>
    <hr/>
    <div>
      <ul class="list-group">
        <li v-for="l in logs" class="list-group-item">
          {{moment(l.created_at).fromNow()}}: {{l.message}}
        </li>
      </ul>
    </div>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {get} from '../../../utils'
import moment from 'moment'

export default {
  name: 'users-logs',
  data () {
    return {
      logs: []
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/users/logs', null, function (rst) {
      this.logs = rst
    }.bind(this))
  },
  methods: {
    moment
  }
}
</script>

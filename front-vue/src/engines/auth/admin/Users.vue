<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-users') }}</h3>
    <hr/>
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('auth.models.user')}}</th>
          <th>{{$t('auth.attributes.user-lastSignInAt')}}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in items">
          <th scope="row">{{i.name}}&lt;{{i.email}}&gt;</th>
          <td>{{i.last_sign_in_at}}</td>
        </tr>
      </tbody>
    </table>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {get} from '../../../utils'

export default {
  name: 'admin-users',
  data () {
    return {
      items: []
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/admin/users', null, function (rst) {
      this.items = rst
    }.bind(this))
  },
  methods: {
  }
}
</script>

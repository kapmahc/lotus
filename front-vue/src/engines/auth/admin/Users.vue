<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-users') }}</h3>
    <hr/>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'

export default {
  name: 'admin-users',
  data () {
    return {
    }
  },
  components: {
    AppDashboard
  },
  created () {
  },
  methods: {
  }
}
</script>

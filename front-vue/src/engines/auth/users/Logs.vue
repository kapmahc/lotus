<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.logs') }}</h3>
    <hr/>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'

export default {
  name: 'users-logs',
  data () {
    return {
      logs: []
    }
  },
  components: {
    AppDashboard
  }
}
</script>

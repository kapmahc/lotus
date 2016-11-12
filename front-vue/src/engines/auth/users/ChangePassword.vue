<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.change-password') }}</h3>
    <hr/>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'

export default {
  name: 'users-change-password',
  data () {
    return {
    }
  },
  components: {
    AppDashboard
  }
}
</script>

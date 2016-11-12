<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.info') }}</h3>
    <hr/>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'

export default {
  name: 'users-info',
  data () {
    return {
      info: {}
    }
  },
  components: {
    AppDashboard
  }
}
</script>

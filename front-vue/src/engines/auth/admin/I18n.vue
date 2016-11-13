<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-i18n') }}</h3>
    <hr/>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'

export default {
  name: 'admin-i18n',
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

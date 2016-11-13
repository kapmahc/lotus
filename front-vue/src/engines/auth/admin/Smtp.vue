<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-smtp') }}</h3>
    <hr/>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'

export default {
  name: 'admin-smtp',
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

<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-author') }}</h3>
    <hr/>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'

export default {
  name: 'admin-author',
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

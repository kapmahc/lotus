<template>
  <div id="root">
    <layout-header/>
    <div class="container">
      <slot/>
      <layout-footer/>
    </div>
  </div>
</template>

<script>
import LayoutHeader from '../components/Header'
import LayoutFooter from '../components/Footer'
import actions from './actions'

export default {
  name: 'app-layout',
  components: {
    LayoutHeader,
    LayoutFooter
  },
  created () {
    if (!this.$store.state.siteInfo.title) {
      this.$store.commit(actions.auth.refreshLayout)
    }
    if (!this.$store.state.currentUser.uid) {
      this.$store.commit(
        actions.auth.signIn,
        window.sessionStorage.getItem('token')
      )
    }
  }
}
</script>

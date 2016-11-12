<template>
  <div id="root">
    <nav class="navbar navbar-fixed-top navbar-dark bg-inverse">
      <router-link class="navbar-brand" :to="{ name: 'home' }">
        {{ info.subTitle }}
      </router-link>
      <ul class="nav navbar-nav">
        <li class="nav-item active">
          <router-link class="nav-link" :to="{ name: 'users.logs' }">
            {{ $t('auth.pages.dashboard') }}
            <span class="sr-only">(current)</span>
          </router-link>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">About111</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Contact</a>
        </li>
      </ul>

      <ul class="nav navbar-nav float-xs-right">
        <lang-bar/>
        <personal-bar/>
      </ul>
     </nav>
    <div class="container">
      <slot/>
      <layout-footer/>
    </div>
  </div>
</template>

<script>
import LayoutFooter from '../components/Footer'
import LangBar from '../components/LangBar'
import PersonalBar from '../components/PersonalBar'

export default {
  name: 'app-dashboard',
  computed: {
    info () {
      return this.$store.state.siteInfo
    }
  },
  components: {
    LayoutFooter,
    LangBar,
    PersonalBar
  }
}
</script>

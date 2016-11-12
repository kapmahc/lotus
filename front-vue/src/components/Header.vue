<template>
  <nav class="navbar navbar-fixed-top navbar-dark bg-inverse">
    <router-link class="navbar-brand" :to="{ name: 'home' }">
      {{ info.subTitle }}
    </router-link>
    <ul class="nav navbar-nav">
      <li class="nav-item active">
        <router-link class="nav-link" :to="{ name: 'home' }">
          {{ $t('auth.pages.home') }}
          <span class="sr-only">(current)</span>
        </router-link>        
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">About</a>
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
</template>

<script>
import LangBar from './LangBar'
import PersonalBar from './PersonalBar'

export default {
  name: 'layout-header',
  computed: {
    info () {
      return this.$store.state.siteInfo
    }
  },
  components: {
    LangBar,
    PersonalBar
  }
}
</script>

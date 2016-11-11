<template>
  <li v-if="user.uid" class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="switch-lang-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t('auth.pages.welcome', {name: user.sub}) }}
    </a>
    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="switch-lang-bar">
      <a v-on:click="signOut()" class="dropdown-item">{{ $t('auth.pages.sign-out') }}</a>
    </div>
  </li>
  <li v-else class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="switch-lang-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t('auth.pages.sign-in-or-up') }}
    </a>
    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="switch-lang-bar">
      <router-link v-for="l in non_sign_in" class="dropdown-item" :to="{ name: `users.${l}` }">
        {{ $t(`auth.pages.${l}`) }}
      </router-link>
    </div>
  </li>
</template>

<script>
import actions from '../engines/actions'
import router from '../engines/router'

export default {
  name: 'forgot-password',
  data () {
    return {
      non_sign_in: [
        'sign-in',
        'sign-up',
        'forgot-password',
        'confirm',
        'unlock'
      ]
    }
  },
  computed: {
    user () {
      return this.$store.state.currentUser
    }
  },
  methods: {
    signOut () {
      window.sessionStorage.removeItem('token')
      this.$store.commit(actions.auth.signOut)
      router.push({name: 'usres.sign-in'})
    }
  }
}
</script>

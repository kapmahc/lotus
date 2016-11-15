<template>
  <li v-if="user.uid" class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" :id="bar_id" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t('auth.pages.welcome', {name: user.sub}) }}
    </a>
    <div class="dropdown-menu dropdown-menu-right" :aria-labelledby="bar_id">
      <router-link class="dropdown-item" :to="{ name: 'dashboard' }">
        {{ $t('auth.pages.dashboard') }}
      </router-link>
      <a v-on:click="signOut()" class="dropdown-item">{{ $t('auth.pages.sign-out') }}</a>
    </div>
  </li>
  <li v-else class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" :id="bar_id" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t('auth.pages.sign-in-or-up') }}
    </a>
    <div class="dropdown-menu dropdown-menu-right" :aria-labelledby="bar_id">
      <router-link v-for="l in non_sign_in" class="dropdown-item" :to="{ name: `users.${l}` }">
        {{ $t(`auth.pages.${l}`) }}
      </router-link>
    </div>
  </li>
</template>

<script>
import actions from '../engines/actions'

export default {
  name: 'forgot-password',
  data () {
    return {
      bar_id: 'personal-bar',
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
      this.$store.commit(actions.auth.signOut)
      window.sessionStorage.removeItem('token')
      this.$router.push({name: 'users.sign-in'})
      // window.location.refresh
      // this.$router.push({name: 'users.sign-in'})
    }
  },
  created () {
    if (!this.$store.state.currentUser.uid) {
      this.$store.commit(
        actions.auth.signIn,
        window.sessionStorage.getItem('token')
      )
    }
  }
}
</script>

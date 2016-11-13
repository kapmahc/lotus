<template>
  <ul>
    <li v-for="l in links">
      <router-link :to="{ name: `users.${l}` }">
        {{ $t(`auth.pages.${l}`) }}
      </router-link>
    </li>
    <li>
      <router-link :to="{ name: 'leavewords.new' }">
        {{ $t('auth.pages.new-leaveword') }}
      </router-link>
    </li>
  </ul>
</template>

<script>
export default {
  name: 'users-non-sign-in-links',
  data () {
    return {
      links: [
        'sign-in',
        'sign-up',
        'forgot-password',
        'confirm',
        'unlock'
      ]
    }
  }
}
</script>

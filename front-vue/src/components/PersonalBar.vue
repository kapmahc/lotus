<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="switch-lang-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t('auth.pages.sign-in-or-up') }}
    </a>
    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="switch-lang-bar">
      <router-link class="dropdown-item" :to="{ name: 'users.sign-in' }">
        {{ $t('auth.pages.sign-in') }}
      </router-link>
      <router-link class="dropdown-item" :to="{ name: 'users.sign-up' }">
        {{ $t('auth.pages.sign-up') }}
      </router-link>
    </div>
  </li>
</template>

<script>
export default {
  name: 'forgot-password',
  data () {
    return {
    }
  }
}
</script>

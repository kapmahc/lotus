<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="switch-lang-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t('lang-bar.switch') }}
    </a>
    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="switch-lang-bar">
      <a v-on:click="swtichLang(l)" v-for="l in info.languages" class="dropdown-item">
        {{ $t(`languages.${l}`) }}
      </a>
    </div>
  </li>
</template>

<script>
import {swtichLang} from '../i18n'

export default {
  name: 'lang-bar',
  computed: {
    info () {
      return this.$store.state.siteInfo
    }
  },
  methods: {
    swtichLang
  }
}
</script>

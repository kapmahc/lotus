<template>
  <div class="row">
    <hr>
      <footer>
        <p class="float-xs-right">
          {{ $t('lang-bar.other') }}:
          <button v-on:click="swtichLang(l)" v-for="l in info.languages" class="btn btn-link btn-sm">
            {{ $t(`languages.${l}`) }}
          </button>
        </p>
        <p>
          &copy; {{ info.copyright }}
          &middot; <a href="#">Privacy</a>
          &middot; <a href="#">Terms</a>
        </p>
      </footer>
  </div>
</template>

<script>
import {swtichLang} from '../i18n'

export default {
  name: 'layout-footer',
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

<template>
  <app-layout>
    <p>A paragraph for the main content.</p>
    <p>And another one.</p>
  </app-layout>
</template>

<script>
import AppLayout from './Layout'

export default {
  name: 'home-page',
  data () {
    return {
    }
  },
  components: {
    AppLayout
  }
}
</script>

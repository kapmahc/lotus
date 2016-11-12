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

export default {
  name: 'app-layout',
  components: {
    LayoutHeader,
    LayoutFooter
  }
}
</script>

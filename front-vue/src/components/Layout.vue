<template>
  <div id="root">
    <layout-header/>
    <div class="container">
      <layout-footer/>
    </div>
  </div>
</template>

<script>
import LayoutHeader from './Header'
import LayoutFooter from './Footer'

export default {
  name: 'app-layout',
  components: {
    LayoutHeader,
    LayoutFooter
  }
}
</script>

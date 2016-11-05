<template>
  <div class="row">
    <hr>
      <footer>
        <p class="float-xs-right"><a href="#">Back to top</a></p>
        <p>
          &copy; {{ $t('attributes.body') }}
          &middot; <a href="#">Privacy</a>
          &middot; <a href="#">Terms</a>
        </p>
      </footer>
  </div>
</template>

<script>
export default {
  name: 'layout-footer',
  data () {
    return {
    }
  }
}
</script>

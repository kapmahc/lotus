<template>
  <div class="row">
    <div v-for="item in items"  class="col-md-12">
      <br/>
      <h4>{{ $t(item.title) }}</h4>
      <hr/>
      <div class="btn-group" role="group">
        <router-link v-for="lk in item.links" class="btn btn-secondary" :to="{ name: lk.href }">
          {{ $t(lk.label) }}
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['items'],
  name: 'nav-pane'
}
</script>

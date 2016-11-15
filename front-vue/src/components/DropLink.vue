<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" :id="`top-nav-${id}`" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t(item.title) }}
    </a>
    <div class="dropdown-menu" :aria-labelledby="`top-nav-${id}`">
      <router-link v-for="lk in item.links" class="dropdown-item" :to="{ name: lk.href }">
        {{ $t(lk.label) }}
      </router-link>
    </div>
  </li>
</template>

<script>
export default {
  props: ['item', 'id'],
  name: 'drop-link'
}
</script>

<template>
  <ul class="nav nav-tabs">
    <li class="nav-item">
      <router-link class="nav-link" :to="{name: 'shop.countries.index'}">
        {{$t('shop.models.country')}}
      </router-link>
    </li>
    <li class="nav-item">
      <router-link class="nav-link" :to="{name: 'shop.states.index'}">
        {{$t('shop.models.state')}}
      </router-link>
    </li>
  </ul>
</template>

<script>
export default {
  name: 'shop-shipment-zone'
}
</script>

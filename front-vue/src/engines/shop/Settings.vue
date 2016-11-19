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
    <li class="nav-item">
      <router-link class="nav-link" :to="{name: 'shop.payment-methods.index'}">
        {{$t('shop.models.payment-method')}}
      </router-link>
    </li>
    <li class="nav-item">
      <router-link class="nav-link" :to="{name: 'shop.shipping-methods.index'}">
        {{$t('shop.models.shipping-method')}}
      </router-link>
    </li>
    <li class="nav-item">
      <router-link class="nav-link" :to="{name: 'shop.currencies.index'}">
        {{$t('shop.models.currency')}}
      </router-link>
    </li>
    <li class="nav-item">
      <router-link class="nav-link" :to="{name: 'shop.postal-codes.index'}">
        {{$t('shop.models.postal-code')}}
      </router-link>
    </li>
  </ul>
</template>

<script>
export default {
  name: 'shop-admin-settings'
}
</script>

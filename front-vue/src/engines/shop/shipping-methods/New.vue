<template>
  <app-dashboard>
    <h3>{{$t('buttons.new')}} {{ $t('shop.models.shipping-method') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="name">{{ $t("attributes.name") }}</label>
        <input v-model="name" type="text" class="form-control" id="name">
      </div>
      <div class="form-group">
        <label for="type">{{ $t("shop.attributes.shipping-method-tracking") }}</label>
        <input v-model="tracking" type="text" class="form-control" id="tracking">
      </div>
      <div class="form-check">
        <label class="form-check-label">
          <input class="form-check-input" type="checkbox" v-model="active">
          {{ $t("attributes.active") }}
        </label>
      </div>
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {postForm} from '../../../utils'

export default {
  name: 'new-shipping-method',
  data () {
    return {
      name: '',
      tracking: '',
      active: true
    }
  },
  components: {
    AppDashboard
  },
  methods: {
    onSubmit () {
      postForm(
        '/shop/shipping-methods',
        {
          type: this.type,
          name: this.name,
          active: this.active,
          tracking: this.tracking
        },
        function (result) {
          this.$router.push({name: 'shop.shipping-methods.index'})
        }.bind(this)
      )
    }
  }
}
</script>

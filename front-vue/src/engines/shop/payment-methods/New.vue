<template>
  <app-dashboard>
    <h3>{{$t('buttons.new')}} {{ $t('shop.models.payment-method') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="type">{{ $t("attributes.type") }}</label>
        <input v-model="type" type="text" class="form-control" id="type">
      </div>
      <div class="form-group">
        <label for="name">{{ $t("attributes.name") }}</label>
        <input v-model="name" type="text" class="form-control" id="name">
      </div>
      <div class="form-check">
        <label class="form-check-label">
          <input class="form-check-input" type="checkbox" v-model="active">
          {{ $t("attributes.active") }}
        </label>
      </div>
      <div class="form-group">
        <label for="description">{{ $t("attributes.description") }}</label>
        <textarea v-model="description" class="form-control" id="description" rows="3"></textarea>
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
  name: 'new-payment-method',
  data () {
    return {
      name: '',
      type: '',
      description: '',
      active: true
    }
  },
  components: {
    AppDashboard
  },
  methods: {
    onSubmit () {
      postForm(
        '/shop/payment-methods',
        {
          name: this.name,
          type: this.type,
          active: this.active,
          description: this.description
        },
        function (result) {
          this.$router.push({name: 'shop.payment-methods.index'})
        }.bind(this)
      )
    }
  }
}
</script>

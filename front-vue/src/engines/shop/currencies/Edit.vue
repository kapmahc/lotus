<template>
  <app-dashboard>
    <h3>{{$t('buttons.edit')}} {{ $t('shop.models.currency') }} [{{id}}]</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="cid">{{ $t("shop.attributes.currency-cid") }}</label>
        <input readonly v-model="cid" type="text" class="form-control" id="cid">
      </div>
      <div class="form-group">
        <label for="name">{{ $t("attributes.name") }}</label>
        <input readonly v-model="name" type="text" class="form-control" id="name">
      </div>
      <div class="form-group">
        <label for="country">{{ $t("shop.models.country") }}</label>
        <input readonly v-model="country" type="text" class="form-control" id="country">
      </div>
      <div class="form-group">
        <label for="rate">{{ $t("shop.attributes.currency-rate") }}</label>
        <input v-model="rate" type="text" class="form-control" id="rate">
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
import {postForm, get} from '../../../utils'

export default {
  name: 'shop-admin-edit-curreny',
  data () {
    return {
      id: this.$route.params.id,
      name: '',
      cid: '',
      country: '',
      active: true
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get(`/shop/currencies/${this.id}`, null, function (rst) {
      this.name = rst.name
      this.cid = rst.cid
      this.country = rst.country
      this.active = rst.active
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        `/shop/currencies/${this.id}`,
        {
          rate: this.rate,
          active: this.active
        },
        function (result) {
          this.$router.push({name: 'shop.currencies.index'})
        }.bind(this)
      )
    }
  }
}
</script>

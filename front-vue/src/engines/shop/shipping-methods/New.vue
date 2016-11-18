<template>
  <app-dashboard>
    <h3>{{$t('buttons.new')}} {{ $t('shop.models.state') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="name">{{ $t("attributes.name") }}</label>
        <input v-model="name" type="text" class="form-control" id="name">
      </div>
      <div class="form-group">
        <label for="country_id">{{$t("shop.models.country")}}</label>
        <select class="form-control" id="country_id" v-model="country_id">
          <option v-for="c in countries" :value="c.id">{{c.name}}</option>
        </select>
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
  name: 'new-state',
  data () {
    return {
      name: '',
      country_id: 0,
      countries: []
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/shop/countries', null, function (rst) {
      this.countries = rst
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        '/shop/states',
        {
          name: this.name,
          country_id: this.country_id
        },
        function (result) {
          this.$router.push({name: 'shop.states.index'})
        }.bind(this)
      )
    }
  }
}
</script>

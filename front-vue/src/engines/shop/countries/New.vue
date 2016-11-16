<template>
  <app-dashboard>
    <h3>{{ $t('shop.pages.new-country') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
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
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {postForm} from '../../../utils'

export default {
  name: 'new-country',
  data () {
    return {
      name: '',
      active: true
    }
  },
  components: {
    AppDashboard
  },
  methods: {
    onSubmit () {
      postForm(
        '/shop/countries',
        {
          name: this.name,
          active: this.active
        },
        function (result) {
          this.$router.push({name: 'shop.countries.index'})
        }.bind(this)
      )
    }
  }
}
</script>

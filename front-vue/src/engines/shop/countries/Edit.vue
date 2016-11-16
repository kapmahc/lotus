<template>
  <app-dashboard>
    <h3>{{ $t('shop.pages.edit-country', {id: id}) }}</h3>
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
import {postForm, get} from '../../../utils'

export default {
  name: 'edit-country',
  data () {
    return {
      id: this.$route.params.id,
      name: '',
      active: true
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get(`/shop/countries/${this.id}`, null, function (rst) {
      this.name = rst.name
      this.active = rst.active
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        `/shop/countries/${this.id}`,
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

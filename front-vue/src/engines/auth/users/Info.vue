<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.info') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="email">{{ $t("attributes.email") }}</label>
        <input v-model="email" readonly type="email" class="form-control" id="email">
      </div>
      <div class="form-group">
        <label for="name">{{ $t("attributes.username") }}</label>
        <input v-model="name" type="text" class="form-control" id="name">
      </div>
      <div class="form-group">
        <label for="home">{{ $t("auth.attributes.user-home") }}</label>
        <input v-model="home" type="text" class="form-control" id="home">
      </div>
      <div class="form-group">
        <label for="logo">{{ $t("auth.attributes.user-logo") }}</label>
        <input v-model="logo" type="text" class="form-control" id="logo">
      </div>
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {get, postForm} from '../../../utils'

export default {
  name: 'users-info',
  data () {
    return {
      email: '',
      name: '',
      home: '',
      logo: ''
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/users/info', null, function (rst) {
      this.email = rst.email
      this.name = rst.name
      this.home = rst.home
      this.logo = rst.logo
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        '/users/info',
        {
          name: this.name,
          home: this.home,
          logo: this.logo
        },
        function (result) {
          window.alert(this.$t('messages.success'))
        }.bind(this)
      )
    }
  }
}
</script>

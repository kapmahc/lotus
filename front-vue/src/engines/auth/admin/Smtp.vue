<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-smtp') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="host">{{ $t("attributes.host") }}</label>
        <input v-model="host" type="text" class="form-control" id="host">
      </div>
      <div class="form-group">
        <label for="port">{{ $t("attributes.port") }}</label>
        <select v-model="port" class="form-control" id="port">
          <option>25</option>
          <option>465</option>
          <option>587</option>
        </select>
      </div>
      <div class="form-group">
        <label for="username">{{ $t("attributes.username") }}</label>
        <input v-model="username" type="text" class="form-control" id="username">
      </div>
      <div class="form-group">
        <label for="password">{{ $t("attributes.password") }}</label>
        <input v-model="password" type="password" aria-describedby="passwordHelp" class="form-control" id="password">
        <small id="passwordHelp" class="form-text text-muted">{{ $t('auth.helpers.password-must-in-size') }}</small>
      </div>
      <div class="form-group">
        <label for="passwordConfirmation">{{ $t("attributes.passwordConfirmation") }}</label>
        <input v-model="passwordConfirmation" type="password" aria-describedby="passwordConfirmationHelp" class="form-control" id="passwordConfirmation">
        <small id="passwordConfirmationHelp" class="form-text text-muted">{{ $t('helpers.passwords-must-match') }}</small>
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
  name: 'admin-smtp',
  data () {
    return {
      host: '',
      port: 25,
      username: '',
      password: '',
      passwordConfirmation: ''
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/admin/smtp', null, function (rst) {
      this.host = rst.host
      this.port = rst.port
      this.username = rst.username
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        '/admin/smtp',
        {
          host: this.host,
          port: this.port,
          username: this.username,
          password: this.password,
          passwordConfirmation: this.passwordConfirmation
        },
        function (result) {
          window.alert(this.$t('messages.success'))
        }.bind(this)
      )
    }
  }
}
</script>

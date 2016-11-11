<template>
  <app-layout>
    <h3>{{ $t('auth.pages.reset-password') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
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
    <br/>
    <SharedLinks />
  </app-layout>
</template>

<script>
import AppLayout from '../Layout'
import SharedLinks from './NonSignInLinks'
import {postForm} from '../../utils'

export default {
  name: 'users-reset-password',
  data () {
    return {
      password: '',
      passwordConfirmation: ''
    }
  },
  components: {
    AppLayout,
    SharedLinks
  },
  methods: {
    onSubmit () {
      postForm(
        '/users/reset-password',
        {
          token: this.$route.query.token,
          password: this.password,
          passwordConfirmation: this.passwordConfirmation
        },
        function (result) {
          window.alert(result.message)
        }
      )
    }
  }
}
</script>

<template>
  <app-layout>
    <h3>{{ $t('auth.pages.sign-in') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="email">{{ $t("attributes.email") }}</label>
        <input v-model="email" type="email" class="form-control" id="email">
      </div>
      <div class="form-group">
        <label for="password">{{ $t("attributes.password") }}</label>
        <input v-model="password" type="password" class="form-control" id="password">
      </div>
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
    <br/>
    <SharedLinks />
  </app-layout>
</template>

<script>
import AppLayout from '../../Layout'
import SharedLinks from './NonSignInLinks'
import {postForm} from '../../../utils'
import actions from '../actions'

export default {
  name: 'users-sign-in',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  components: {
    AppLayout,
    SharedLinks
  },
  methods: {
    onSubmit () {
      postForm(
        '/users/sign-in',
        {
          email: this.email,
          password: this.password
        },
        function (result) {
          var token = result.token
          this.$store.commit(actions.signIn, token)
          window.sessionStorage.setItem('token', token)
          this.$router.push({name: 'home'})
        }.bind(this)
      )
    }
  }
}
</script>

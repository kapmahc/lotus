<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.change-password') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="password">{{ $t("auth.attributes.user-currentPassword") }}</label>
        <input v-model="currentPassword" type="password" aria-describedby="currentPasswordHelp" class="form-control" id="currentPassword">
        <small id="currentPasswordHelp" class="form-text text-muted">{{ $t('auth.helpers.need-password-to-change') }}</small>
      </div>
      <div class="form-group">
        <label for="newPassword">{{ $t("auth.attributes.user-newPassword") }}</label>
        <input v-model="newPassword" type="password" aria-describedby="passwordHelp" class="form-control" id="newPassword">
        <small id="newPasswordHelp" class="form-text text-muted">{{ $t('auth.helpers.password-must-in-size') }}</small>
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
import {postForm} from '../../../utils'

export default {
  name: 'users-change-password',
  data () {
    return {
      newPassword: '',
      currentPassword: '',
      passwordConfirmation: ''
    }
  },
  components: {
    AppDashboard
  },
  methods: {
    onSubmit () {
      postForm(
        '/users/change-password',
        {
          currentPassword: this.currentPassword,
          newPassword: this.newPassword,
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

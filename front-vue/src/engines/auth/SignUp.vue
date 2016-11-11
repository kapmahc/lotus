<template>
  <app-layout>
    <h3>{{ $t('auth.pages.sign-up') }}</h3>
    <hr/>
    <form>
      <div class="form-group">
        <label for="name">{{ $t("attributes.username") }}</label>
        <input type="text" class="form-control" id="name">
      </div>
      <div class="form-group">
        <label for="email">{{ $t("attributes.email") }}</label>
        <input type="email" class="form-control" id="email">
      </div>
      <div class="form-group">
        <label for="password">{{ $t("attributes.password") }}</label>
        <input type="password" aria-describedby="passwordHelp" class="form-control" id="password">
        <small id="passwordHelp" class="form-text text-muted">{{ $t('auth.helpers.password-must-in-size') }}</small>
      </div>
      <div class="form-group">
        <label for="passwordConfirmation">{{ $t("attributes.passwordConfirmation") }}</label>
        <input type="password" aria-describedby="passwordConfirmationHelp" class="form-control" id="passwordConfirmation">
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

export default {
  name: 'users-sign-up',
  data () {
    return {
    }
  },
  components: {
    AppLayout,
    SharedLinks
  }
}
</script>

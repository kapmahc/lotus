<template>
  <app-layout>
    <h3>{{ $t('auth.pages.forgot-password') }}</h3>
    <hr/>
    <form>
      <div class="form-group">
        <label for="email">{{ $t("attributes.email") }}</label>
        <input type="email" class="form-control" id="email">
      </div>
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
    <br/>
    <SharedLinks />
    import SharedLinks from './NonSignInLinks'
  </app-layout>
</template>

<script>
import AppLayout from '../Layout'
import SharedLinks from './NonSignInLinks'

export default {
  name: 'users-forgot-password',
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

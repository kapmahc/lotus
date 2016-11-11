<template>
  <app-layout>
    <h3>{{ $t('auth.pages.unlock') }}</h3>
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
  </app-layout>
</template>

<script>
import AppLayout from '../Layout'
import SharedLinks from './NonSignInLinks'

export default {
  name: 'users-unlock',
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

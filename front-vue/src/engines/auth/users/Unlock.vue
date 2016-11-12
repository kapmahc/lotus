<template>
  <app-layout>
    <h3>{{ $t('auth.pages.unlock') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="email">{{ $t("attributes.email") }}</label>
        <input v-model="email" type="email" class="form-control" id="email">
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

export default {
  name: 'users-unlock',
  data () {
    return {
      email: ''
    }
  },
  components: {
    AppLayout,
    SharedLinks
  },
  methods: {
    onSubmit () {
      postForm(
        '/users/unlock',
        {
          email: this.email
        },
        function (result) {
          window.alert(result.message)
        }
      )
    }
  }
}
</script>

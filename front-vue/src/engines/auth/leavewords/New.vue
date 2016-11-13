<template>
  <app-layout>
    <h3>{{ $t('auth.pages.new-leaveword') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="content">{{ $t("attributes.content") }}</label>
        <textarea v-model="content" class="form-control" id="content" rows="3"></textarea>
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
import SharedLinks from '../users/NonSignInLinks'
import {postForm} from '../../../utils'

export default {
  name: 'new-leaveword',
  data () {
    return {
      content: ''
    }
  },
  components: {
    AppLayout,
    SharedLinks
  },
  methods: {
    onSubmit () {
      postForm(
        '/leavewords',
        {
          content: this.content
        },
        function (result) {
          window.alert(result.message)
        }
      )
    }
  }
}
</script>

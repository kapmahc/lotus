<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.new-notice') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="content">{{ $t("attributes.content") }}</label>
        <textarea v-model="content" class="form-control" id="content" rows="3"></textarea>
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
  name: 'new-notice',
  data () {
    return {
      content: ''
    }
  },
  components: {
    AppDashboard
  },
  methods: {
    onSubmit () {
      postForm(
        '/notices',
        {
          content: this.content
        },
        function (result) {
          this.$router.push({name: 'notices.admin'})
        }.bind(this)
      )
    }
  }
}
</script>

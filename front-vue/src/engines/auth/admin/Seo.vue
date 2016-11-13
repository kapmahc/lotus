<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-seo') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="google">{{ $t("auth.attributes.google-verify-code") }}</label>
        <input v-model="google" type="text" class="form-control" id="google">
      </div>
      <div class="form-group">
        <label for="baidu">{{ $t("auth.attributes.baidu-verify-code") }}</label>
        <input v-model="baidu" type="text" class="form-control" id="baidu">
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
  name: 'admin-seo',
  data () {
    return {
      google: '',
      baidu: ''
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/admin/seo', null, function (rst) {
      this.google = rst.google
      this.baidu = rst.baidu
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        '/admin/seo',
        {
          google: this.google,
          baidu: this.baidu
        },
        function (result) {
          window.alert(result.message)
        }
      )
    }
  }
}
</script>

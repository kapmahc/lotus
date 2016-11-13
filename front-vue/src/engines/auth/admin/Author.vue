<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-author') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="email">{{ $t("attributes.email") }}</label>
        <input v-model="email" type="email" class="form-control" id="email">
      </div>
      <div class="form-group">
        <label for="name">{{ $t("attributes.username") }}</label>
        <input v-model="name" type="text" class="form-control" id="name">
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
  name: 'admin-author',
  data () {
    return {
      email: '',
      name: ''
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/admin/author', null, function (rst) {
      this.email = rst.email
      this.name = rst.name
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        '/admin/author',
        {
          name: this.name,
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

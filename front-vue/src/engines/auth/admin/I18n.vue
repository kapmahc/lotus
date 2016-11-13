<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-i18n') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="code">{{ $t("auth.attributes.locale-code") }}</label>
        <input v-model="code" type="text" class="form-control" id="code">
      </div>
      <div class="form-group">
        <label for="message">{{ $t("auth.attributes.locale-message") }}</label>
        <textarea v-model="message" class="form-control" id="message" rows="3"></textarea>
      </div>
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
    <br/>
    <div class="list-group">
      <a v-for="item in items" class="list-group-item list-group-item-action">
        <h5 class="list-group-item-heading">{{item.code}}</h5>
        <p class="list-group-item-text">{{item.message}}</p>
      </a>
    </div>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {get, postForm} from '../../../utils'

export default {
  name: 'admin-i18n',
  data () {
    return {
      code: '',
      message: '',
      items: []
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/admin/i18n', null, function (rst) {
      this.items = rst
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        '/admin/i18n',
        {
          code: this.code,
          message: this.message
        },
        function (result) {
          window.alert(result.message)
        }
      )
    }
  }
}
</script>

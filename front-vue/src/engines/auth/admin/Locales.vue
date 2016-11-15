<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-locales') }}</h3>
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
      <a v-for="item in items" v-on:click="onEdit(item)" class="list-group-item list-group-item-action">
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
  name: 'locales-index',
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
    this.onRefresh()
  },
  methods: {
    onEdit (item) {
      this.code = item.code
      this.message = item.message
    },
    onRefresh () {
      get('/admin/locales', null, function (rst) {
        this.items = rst
      }.bind(this))
    },
    onSubmit () {
      postForm(
        '/admin/locales',
        {
          code: this.code,
          message: this.message
        },
        function (result) {
          window.alert(this.$t('messages.success'))
          this.onRefresh()
        }.bind(this)
      )
    }
  }
}
</script>

<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.admin-base') }}</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="title">{{ $t("attributes.title") }}</label>
        <input v-model="title" type="text" class="form-control" id="title">
      </div>
      <div class="form-group">
        <label for="subTitle">{{ $t("auth.attributes.subTitle") }}</label>
        <input v-model="subTitle" type="text" class="form-control" id="subTitle">
      </div>
      <div class="form-group">
        <label for="keywords">{{ $t("attributes.keywords") }}</label>
        <input v-model="keywords" type="text" class="form-control" id="keywords">
      </div>
      <div class="form-group">
        <label for="description">{{ $t("attributes.description") }}</label>
        <textarea v-model="description" class="form-control" id="description" rows="3"></textarea>
      </div>
      <div class="form-group">
        <label for="copyright">{{ $t("auth.attributes.copyright") }}</label>
        <input v-model="copyright" type="text" class="form-control" id="copyright">
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
  name: 'admin-base',
  data () {
    return {
      title: '',
      subTitle: '',
      keywords: '',
      description: '',
      copyright: ''
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/admin/base', null, function (rst) {
      this.title = rst.title
      this.subTitle = rst.subTitle
      this.keywords = rst.keywords
      this.description = rst.description
      this.copyright = rst.copyright
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        '/admin/base',
        {
          title: this.title,
          subTitle: this.subTitle,
          keywords: this.keywords,
          description: this.description,
          copyright: this.copyright
        },
        function (result) {
          window.alert(this.$t('messages.success'))
        }.bind(this)
      )
    }
  }
}
</script>

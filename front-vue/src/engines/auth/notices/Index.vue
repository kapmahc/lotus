<template>
  <app-layout>
    <h3>{{ $t('auth.pages.notices') }}</h3>
    <hr/>
    <blockquote v-for="i in items" class="blockquote">
      <p class="mb-0">{{i.content}}</p>
      <footer class="blockquote-footer">
        {{moment(i.updated_at).fromNow()}}
      </footer>
    </blockquote>
  </app-layout>
</template>

<script>
import AppLayout from '../../Layout'
import {get} from '../../../utils'
import moment from 'moment'

export default {
  name: 'users-sign-in',
  data () {
    return {
      items: []
    }
  },
  created () {
    get('/notices', null, function (rst) {
      this.items = rst
    }.bind(this))
  },
  components: {
    AppLayout
  },
  methods: {
    moment
  }
}
</script>

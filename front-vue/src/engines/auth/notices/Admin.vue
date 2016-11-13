<template>
  <app-dashboard>
    <h3>
      {{ $t('auth.pages.notices') }}
      <router-link class="btn btn-link" :to="{name: 'notices.new'}">
        {{$t('buttons.new')}}
      </router-link>
    </h3>
    <hr/>
    <div class="list-group">
      <a v-for="item in items" class="list-group-item list-group-item-action">
        <h5 class="list-group-item-heading">
          {{item.updated_at}}
          <router-link :to="{name: 'notices.edit', params: { id: item.id }}" class="btn btn-warning btn-sm">
            {{$t('buttons.edit')}}
          </router-link>
          <button v-on:click="onDelete(item.id)" class="btn btn-danger btn-sm">
            {{$t('buttons.remove')}}
          </button>
        </h5>
        <p class="list-group-item-text">{{item.content}}</p>
      </a>
    </div>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {get, _delete} from '../../../utils'

export default {
  name: 'notices',
  data () {
    return {
      items: []
    }
  },
  components: {
    AppDashboard
  },
  created () {
    this.refresh()
  },
  methods: {
    refresh () {
      get('/notices', null, function (rst) {
        this.items = rst
      }.bind(this))
    },
    onDelete (id) {
      if (window.confirm(this.$t('messages.are-you-sure'))) {
        _delete(`/notices/${id}`, function () {
          this.refresh()
        }.bind(this))
      }
    }
  }
}
</script>

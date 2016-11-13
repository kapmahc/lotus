<template>
  <app-dashboard>
    <h3>{{ $t('auth.pages.leavewords') }}</h3>
    <hr/>
    <div class="list-group">
      <a v-for="item in items" class="list-group-item list-group-item-action">
        <h5 class="list-group-item-heading">
          {{item.created_at}}
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
  name: 'leavewords',
  data () {
    return {
      items: []
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get('/leavewords', null, function (rst) {
      this.items = rst
    }.bind(this))
  },
  methods: {
    onDelete (id) {
      if (window.confirm(this.$t('messages.are-you-sure'))) {
        _delete(`/leavewords/${id}`)
      }
    }
  }
}
</script>

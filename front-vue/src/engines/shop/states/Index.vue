<template>
  <app-dashboard>
    <admin-settings/>
    <br/>
    <h3>
      {{ $t('shop.models.state') }}
    </h3>
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('attributes.name')}}</th>
          <th>{{$t('shop.models.country')}}?</th>
          <th>
            {{$t('buttons.manage')}}
            <router-link class="btn btn-link" :to="{name: 'shop.states.new'}">
              {{$t('buttons.new')}}
            </router-link>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in items">
          <th scope="row">{{i.name}}</th>
          <td>{{i.country.name}}</td>
          <td>
            <router-link class="btn btn-warning btn-sm" :to="{name: 'shop.states.edit', params:{id: i.id}}">
              {{$t('buttons.edit')}}
            </router-link>
            <button v-on:click="onDelete(i.id)" class="btn btn-danger btn-sm">
              {{$t('buttons.remove')}}
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import AdminSettings from '../Settings'
import {get, _delete} from '../../../utils'

export default {
  name: 'shop-states',
  data () {
    return {
      items: []
    }
  },
  created () {
    this.refresh()
  },
  components: {
    AppDashboard,
    AdminSettings
  },
  methods: {
    refresh () {
      get('/shop/states', null, function (rst) {
        this.items = rst
      }.bind(this))
    },
    onDelete (id) {
      if (window.confirm(this.$t('messages.are-you-sure'))) {
        _delete(`/shop/states/${id}`, function () {
          this.refresh()
        }.bind(this))
      }
    }
  }
}
</script>

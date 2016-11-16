<template>
  <app-dashboard>
    <h3>
      {{ $t('shop.pages.countries') }}
      <router-link class="btn btn-link" :to="{name: 'shop.countries.new'}">
        {{$t('buttons.new')}}
      </router-link>
    </h3>
    <hr/>
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('attributes.name')}}</th>
          <th>{{$t('attributes.active')}}</th>
          <th>{{$t('shop.attributes.country-states')}}</th>
          <th>{{$t('buttons.manage')}}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in items">
          <th scope="row">{{i.name}}</th>
          <td>{{i.active}}</td>
          <td>
            <ol>
              <li v-for="s in i.states">{{s.name}}</li>
            </ol>
          </td>
          <td>
            <router-link class="btn btn-warning btn-sm" :to="{name: 'shop.countries.edit', params:{id: i.id}}">
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
import {get, _delete} from '../../../utils'

export default {
  name: 'shop-countries',
  data () {
    return {
      items: []
    }
  },
  created () {
    this.refresh()
  },
  components: {
    AppDashboard
  },
  methods: {
    refresh () {
      get('/shop/countries', null, function (rst) {
        this.items = rst
      }.bind(this))
    },
    onDelete (id) {
      if (window.confirm(this.$t('messages.are-you-sure'))) {
        _delete(`/shop/countries/${id}`, function () {
          this.refresh()
        }.bind(this))
      }
    }
  }
}
</script>

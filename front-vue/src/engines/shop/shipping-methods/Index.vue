<template>
  <app-dashboard>
    <br/>
    <h3>
      {{ $t('shop.models.shipping-method') }}
    </h3>
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('attributes.name')}}</th>
          <th>{{$t('attributes.active')}}</th>
          <th>{{$t('shop.attributes.shipping-method-tracking')}}</th>
          <th>
            {{$t('buttons.manage')}}
            <router-link class="btn btn-link" :to="{name: 'shop.shipping-methods.new'}">
              {{$t('buttons.new')}}
            </router-link>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in items">
          <th scope="row">{{i.name}}</th>
          <td>{{i.active}}</td>
          <td>
            <a :href="i.tracking" target="_blank">{{i.tracking}}</a>
          </td>
          <td>
            <router-link class="btn btn-warning btn-sm" :to="{name: 'shop.shipping-methods.edit', params:{id: i.id}}">
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
  name: 'shop-shipping-methods',
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
      get('/shop/shipping-methods', null, function (rst) {
        this.items = rst
      }.bind(this))
    },
    onDelete (id) {
      if (window.confirm(this.$t('messages.are-you-sure'))) {
        _delete(`/shop/shipping-methods/${id}`, function () {
          this.refresh()
        }.bind(this))
      }
    }
  }
}
</script>
